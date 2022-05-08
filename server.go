package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"time"

	"app-api/configs"
	"app-api/db"
	"app-api/ent"
	"app-api/middlewares"
	"app-api/modules/custom_component"

	authHttp "app-api/modules/auth/http"
	shopHttp "app-api/modules/shop/delivery/http"
	userHttp "app-api/modules/user/http"

	appApiGRPC "app-api/modules/shop/delivery/rpc"

	// cors "github.com/rs/cors/wrapper/gin"
	"entgo.io/contrib/entgql"
	cors "github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql"
	gqlgenHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	"app-api/graph/generated"
	resolvers "app-api/graph/resolvers"

	authz "github.com/es-hs/authzclient"
	"github.com/es-hs/erpc"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	_shopModuleUseCase "app-api/modules/shop/usecase"
	_templateSectionModuleUseCase "app-api/modules/template_section/usecase"
	_themeModuleUseCase "app-api/modules/theme/usecase"
	_themeTemplateModuleUseCase "app-api/modules/theme_template/usecase"
	_todoModuleUseCase "app-api/modules/todo/usecase" // just an example
	_userModuleUseCase "app-api/modules/user/usecase"

	esContext "github.com/es-hs/es-helper/context"
	"google.golang.org/grpc/reflection"

	grpcLibs "app-api/libs/grpc"
)

var (
	task = ""
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Info("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()\n", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	if len(os.Args) > 1 {
		readFlags()
	}
}

func main() {
	// todo move to config package
	authzServerAddress := configs.GRPC.Server.Authz
	logrus.Info("Init connection to authz service")
	err := authz.InitAuthClient(authzServerAddress, 1*time.Second, grpc.WithInsecure(), grpc.WithBlock())
	defer authz.Conn.Close()

	// init connection to app service
	erpc.GetConnection(configs.GRPC.Server.Account, 2*time.Second)

	if err != nil {
		// TODO: Sentry here
		// panic(err)
		logrus.Error(err)
	}

	// Connect gRPC DeployService
	if err = grpcLibs.ConnectDeployService(); err == nil {
		defer grpcLibs.GRPCConn.DeployService.Conn.Close()
		defer grpcLibs.GRPCConn.DeployService.Cancel()
	}

	if len(os.Args) > 1 {
		executeCommand()
	} else {
		runServer()
	}
}

func executeCommand() {
	flag.Parse()
	switch task {
	case "service":
		runService()
	case "server":
		runServer()
	default:
		logrus.Warn("Unknow command:", task)
	}
}

func readFlags() {
	flag.StringVar(&task, "task", "", "Task to run: service/server")
}

func runService() {
	db.InitDB()

	client := db.GetEntDB()
	accountConnection, _ := erpc.GetConnection(configs.GRPC.Server.Account, 2*time.Second)
	gRPCPort := configs.GRPC.Port
	lis, err := net.Listen("tcp", "0.0.0.0:"+gRPCPort)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	appApiGRPC.RegisterServer(s, client, accountConnection)
	reflection.Register(s)
	logrus.Infof("server listening at %v\n", lis.Addr())
	logrus.Info("log to know new version updated")
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func runServer() {
	db.InitDB()

	r := gin.Default()
	// Apply the middleware to the router (works with groups too)
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// r.Use(cors.New(config))
	// todo don't do this when push to production
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Accept", "Accept-Language", "Content-Type", "Authorization", "Origin", "ES-SHOP-ID"}
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	corsOptions := cors.New(config)

	// load template
	r.LoadHTMLGlob("./templates/*")
	r.Use(corsOptions)
	// access gin context
	r.Use(GinContextToContextMiddleware())

	client := db.GetClient()

	accountConnection, _ := erpc.GetConnection(configs.GRPC.Server.Account, 2*time.Second)

	defer client.Close()

	rootGroup := r.Group("/")
	userHttp.InitAPIUser(rootGroup, client, accountConnection)
	shopHttp.InitAPIShop(rootGroup, client, accountConnection)
	authHttp.InitAPIAuth(rootGroup)

	g := r.Group("/graphql")

	g.Use(middlewares.AuthenRequest())
	// g.Use(authorizations.WithAuthorization(authorizations.AuthEnforcer))

	srv := graphqlHandler(client, accountConnection)
	{
		g.GET("", playgroundHandler())
		g.POST("/query", gin.WrapH(srv))
		g.GET("/query", gin.WrapH(srv))
	}

	r.Static("/test", "public/graphql")
	r.GET("healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	// token, _ := utils.CreateLoginToken(models.UserCheck{})
	logrus.Info("Server is running...")
	r.Run() // listen and serve on 0.0.0.0:3003 (for windows "localhost:3003")
}

// Defining the Graphql handler
func graphqlHandler(client *ent.Client, accountConnection *grpc.ClientConn) *gqlgenHandler.Server {
	schema := graphqlSchema(client, accountConnection)

	server := gqlgenHandler.NewDefaultServer(schema)
	// Limiting Query Complexity
	server.Use(extension.FixedComplexityLimit(200))
	// Support transactional mutation
	server.Use(entgql.Transactioner{TxOpener: client})

	return server
}

func graphqlSchema(client *ent.Client, accountConnection *grpc.ClientConn) graphql.ExecutableSchema {
	todoUseCase := _todoModuleUseCase.NewTodoUseCase() // just an example
	userUseCase := _userModuleUseCase.NewUserUseCase(client, accountConnection)
	shopUseCase := _shopModuleUseCase.NewShopUseCase(client, accountConnection)
	templateSectionUseCase := _templateSectionModuleUseCase.NewTemplateSectionUseCase()
	templateSectionVersionUseCase := _templateSectionModuleUseCase.NewTemplateSectionVersionUseCase()
	themeUseCase := _themeModuleUseCase.NewThemeUseCase()
	themeTemplateUseCase := _themeTemplateModuleUseCase.NewThemeTemplateUseCase()

	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			Client:                        client,
			TodoUseCase:                   todoUseCase,
			CustomComponentUseCase:        custom_component.NewUseCase(),
			UserUseCase:                   userUseCase,
			ShopUseCase:                   shopUseCase,
			TemplateSectionUseCase:        templateSectionUseCase,
			TemplateSectionVersionUseCase: templateSectionVersionUseCase,
			ThemeUseCase:                  themeUseCase,
			ThemeTemplateUseCase:          themeTemplateUseCase,
		},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	})
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	ginContextKey := esContext.GinContext(esContext.GinContextKey)

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
