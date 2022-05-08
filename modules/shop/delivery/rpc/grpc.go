package rpc

import (
	"context"

	"app-api/ent"
	"app-api/models"
	shopUsecase "app-api/modules/shop/usecase"

	esUtils "github.com/es-hs/es-helper/utils"

	pb "github.com/es-hs/erpc/app-api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShopRpcServer struct {
	pb.UnimplementedShopServer
	shopUsecase shopUsecase.ShopUseCase
}

func RegisterServer(server *grpc.Server, dbEnt *ent.Client, accountConnection *grpc.ClientConn) {
	shopUsecase := shopUsecase.NewShopUseCase(dbEnt, accountConnection)
	pb.RegisterShopServer(server, &ShopRpcServer{
		shopUsecase: shopUsecase,
	})
}

func (instance *ShopRpcServer) CreateShop(ctx context.Context, request *pb.CreateShopRequest) (*pb.CreateShopReply, error) {
	shopCreate := models.Shop{}
	shopCreate.ShopName = request.GetShopName()
	shopCreate.DefaultDomain = esUtils.GenESDomain(request.GetShopName())
	logrus.Info(shopCreate.DefaultDomain)
	logrus.Info(request.GetShopName())
	logrus.Info(esUtils.GenESDomain(request.GetShopName()))
	shopCreate, err := instance.shopUsecase.CreateShop(ctx, shopCreate)
	if err != nil {
		// TODO: Sentry
		logrus.Error(err)
		return &pb.CreateShopReply{}, status.New(codes.Aborted, "FAILURE").Err()
	}

	return &pb.CreateShopReply{}, nil
}

func (instance *ShopRpcServer) CreateShopRevert(ctx context.Context, request *pb.CreateShopRequest) (*pb.CreateShopRevertReply, error) {
	shopCreate := models.Shop{}
	shopCreate.ShopName = request.GetShopName()
	shopCreate.DefaultDomain = esUtils.GenESDomain(request.GetShopName())
	err := instance.shopUsecase.CreateShopRevert(ctx, shopCreate)
	if err != nil {
		return &pb.CreateShopRevertReply{}, status.New(codes.Aborted, "FAILURE").Err()
	}
	return &pb.CreateShopRevertReply{}, nil
}

func (instance *ShopRpcServer) GetShop(ctx context.Context, request *pb.GetShopRequest) (*pb.GetShopReply, error) {
	shopId := request.GetId()
	shopInfo, err := instance.shopUsecase.GetShop(ctx, uint64(shopId))
	if err != nil && !ent.IsNotFound(err) {
		return &pb.GetShopReply{}, status.New(codes.Aborted, "FAILURE").Err()
	}
	if err != nil {
		return &pb.GetShopReply{
			Id:            uint64(0),
			ShopName:      "",
			DefaultDomain: "",
		}, nil
	}
	return &pb.GetShopReply{
		Id:            uint64(shopInfo.ID),
		ShopName:      shopInfo.ShopName,
		DefaultDomain: shopInfo.DefaultDomain,
	}, nil
}

func (instance *ShopRpcServer) GetShopByDomain(ctx context.Context, request *pb.GetShopByDomainRequest) (*pb.GetShopByDomainReply, error) {
	domain := request.GetDomain()
	domain = esUtils.GenESDomain(request.GetDomain())

	shopInfo, err := instance.shopUsecase.GetShopByDomain(ctx, domain)
	logrus.Info(ent.IsNotFound(err))
	if err != nil && !ent.IsNotFound(err) {
		return &pb.GetShopByDomainReply{}, status.New(codes.Aborted, "FAILURE").Err()
	}
	if err != nil {
		return &pb.GetShopByDomainReply{
			Shop: &pb.GetShopReply{
				Id:            0,
				ShopName:      "",
				DefaultDomain: "",
			},
		}, nil
	}
	return &pb.GetShopByDomainReply{
		Shop: &pb.GetShopReply{
			Id:            uint64(shopInfo.ID),
			ShopName:      shopInfo.ShopName,
			DefaultDomain: shopInfo.DefaultDomain,
		},
	}, nil
}
