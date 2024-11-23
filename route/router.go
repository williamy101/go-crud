package route

import (
	"go-asg4/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	productRouter := router.Group("/product")
	{
		productRouter.POST("/", services.CreateProduct)
		productRouter.GET("/", services.ReadProducts)
		productRouter.GET("/search", services.ReadProductByIdOrCategory)
		productRouter.PUT("/:id", services.UpdateProduct)
		productRouter.DELETE("/:id", services.DeleteProduct)
		productRouter.POST("/:id/image", services.UploadProductImage)
		productRouter.GET("/:id/image", services.DownloadProductImage)
	}
	inventoryRouter := router.Group("/inventory")
	{
		inventoryRouter.GET("/", services.ReadProductStock)
		inventoryRouter.PUT("/:product_id", services.UpdateInventory)
	}

	orderRouter := router.Group("/orders")
	{
		orderRouter.POST("/", services.CreateOrder)
		orderRouter.GET("/:id", services.ReadOrderByID)
	}
}
