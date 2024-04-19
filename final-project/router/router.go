package router

import (
	controllers "final-project/controllers"
	middleware "final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/auth")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/variants/:variantUUID", controllers.GetVariantByUUID)
		productRouter.GET("/variants/", controllers.GetVariants)

		productRouter.POST("/variants/", middleware.Authentication(), controllers.CreateVariant)
		productRouter.PUT("/variants/:variantUUID", middleware.Authentication(), middleware.VariantAuthorization(), controllers.UpdateVariant)
		productRouter.DELETE("/variants/:variantUUID", middleware.Authentication(), middleware.VariantAuthorization(), controllers.DeleteVariant)

		productRouter.GET("/:productUUID", controllers.GetProductByUUID)
		productRouter.GET("/", controllers.GetProducts)

		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productUUID", middleware.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productUUID", middleware.ProductAuthorization(), controllers.DeleteProduct)
	}

	return router
}
