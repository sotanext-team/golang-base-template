package http

import (
	"app-api/pkg/oauth2client"

	"github.com/gin-gonic/gin"
)

// A valid oauth2 client (check the store) that additionally requests an OpenID Connect id token

func InitAPIAuth(r *gin.RouterGroup) {

	authController := NewAuthHandler()

	r.GET("/login", authController.Index)
	r.GET("/callback", oauth2client.CallbackHandler())
	// authUserRoutes := r.Group("/auth")
	// rbacRouter := authUserRoutes.Group("/permission-manager")
	{
		// rbacRouter.POST("/policy-add", authController.PolicyAdd)
		// rbacRouter.POST("/policy-adds", authController.PolicyAdds)
		// rbacRouter.DELETE("/policy-remove", authController.PolicyRemove)

		// rbacRouter.POST("/role-adds", authController.RoleAdds)
		// rbacRouter.POST("/role-add", authController.RoleAdd)

		// rbacRouter.GET("/get-all-role", authController.GetAllRoles)
		// rbacRouter.GET("/get-role-for-user", authController.GetRolesForUser)

		// rbacRouter.GET("/get-implicit-roles-for-user", authController.GetImplicitRolesForUser)
		// rbacRouter.GET("/get-permissions-for-user", authController.GetPermissionsForUser)
		// rbacRouter.GET("/get-implicit-permissions-for-user", authController.GetImplicitPermissionsForUser)
		// rbacRouter.GET("/get-policy", authController.GetPolicy)
		// rbacRouter.GET("/get-grouping-policy", authController.GetGroupingPolicy)
		// rbacRouter.DELETE("/delete-role-for-user", authController.DeleteRoleForUser)
	}

	// r.GET("test-send-job", userController.TestSendJob)
	// webhookReceiveRoutes := r.Group("/user")

	// webhookReceiveRoutes.Use(middleware.VerifyWebhookRequest())
	// webhookReceiveRoutes.POST("/webhook", userController.ShowWebhookInfo)

}
