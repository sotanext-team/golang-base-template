package http

import (
	"app-api/ent"
	"app-api/middlewares"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitAPIShop(r *gin.RouterGroup, dbEnt *ent.Client, accountConnection *grpc.ClientConn) {
	shopController := NewShopHandler(dbEnt, accountConnection)

	// r.Use(middlewares.AuthenRequest())
	// r.Use(middlewares.AuthenShopRequest())
	r.GET("/login_shop", shopController.ShopLogin)
	authShopRoutes := r.Group("/auth")
	authShopRoutes.Use(middlewares.AuthenRequest())
	{
		authShopRoutes.POST("/shop/register", shopController.ShopRegister)
		authShopRoutes.POST("/shop/login/:shop-domain", shopController.ShopLogin)
	}
	testPermission := r.Group("/auth/")
	testPermission.Use(middlewares.AuthenShopRequest())
}
