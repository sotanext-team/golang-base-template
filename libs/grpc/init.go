package grpc

import (
	"context"
	"time"

	"app-api/configs"

	pbDS "github.com/es-hs/erpc/deploy"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GRPCConn *gRPCConn
)

type gRPCConn struct {
	DeployService struct {
		Conn   *grpc.ClientConn
		Client pbDS.ShopClient
		Cancel context.CancelFunc
		Ctx    context.Context
	}
}

func init() {
	GRPCConn = &gRPCConn{}
}

func ConnectDeployService() error {
	var err error
	// Contact the server and print out its response.
	GRPCConn.DeployService.Ctx, GRPCConn.DeployService.Cancel = context.WithTimeout(context.Background(), 1*time.Second)
	// Set up a connection to the server.
	target := configs.GRPC.Server.Deploy
	GRPCConn.DeployService.Conn, err = grpc.DialContext(GRPCConn.DeployService.Ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		// TODO: Sentry
		logrus.Error(err)
		return err
	}
	GRPCConn.DeployService.Client = pbDS.NewShopClient(GRPCConn.DeployService.Conn)

	return nil
}
