package http

import (
	"app-api/ent"
	"app-api/middlewares"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitAPIUser(r *gin.RouterGroup, dbEnt *ent.Client, accountConnection *grpc.ClientConn) {
	userController := NewUserHandler(dbEnt, accountConnection)

	// authUserRoutes := r.Group("/auth")
	// {
	// 	authUserRoutes.POST("/register", userController.UserRegister)
	// 	authUserRoutes.POST("/login", userController.UserLogin)
	// }

	apiUserRoutes := r.Group("/api")
	apiUserRoutes.Use(middlewares.AuthenRequest())
	{
		apiUserRoutes.GET("/admin/me", userController.GetCurrentUser)
	}
	// r.GET("test-send-job", userController.TestSendJob)
	// webhookReceiveRoutes := r.Group("/user")

	// webhookReceiveRoutes.Use(middleware.VerifyWebhookRequest())
	// webhookReceiveRoutes.POST("/webhook", userController.ShowWebhookInfo)

}
