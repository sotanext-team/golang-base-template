package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAPIShop(r *gin.RouterGroup, db *gorm.DB) {
	shopController := NewShopHandler(db)

	// r.Use(middlewares.AuthenRequest())
	// r.Use(middlewares.AuthenShopRequest())
	shopRoutes := r.Group("/shops")
	{
		shopRoutes.POST("/", shopController.ShopRegister)
		shopRoutes.GET("/", shopController.List)
		shopRoutes.GET("/:shop_id", shopController.Get)
	}
}
