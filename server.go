package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"runtime"

	"golang-base/configs"
	"golang-base/db"
	shopHttp "golang-base/modules/shop/delivery/http"

	cors "github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	// just an example

	"google.golang.org/grpc/reflection"
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
	// TODO: init grpc connection here

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

	// client := db.GetDB()
	gRPCPort := configs.GRPC.Port
	lis, err := net.Listen("tcp", "0.0.0.0:"+gRPCPort)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
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

	client := db.GetDB()

	rootGroup := r.Group("/")
	shopHttp.InitAPIShop(rootGroup, client)

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
