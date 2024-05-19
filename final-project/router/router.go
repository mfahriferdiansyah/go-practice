package router

import (
	controllers "final-project/controllers"
	middleware "final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	// userRouter := router.Group("/auth")
	// {
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	// }

	// productRouter := router.Group("/products")
	// {
	router.GET("/products/variants/:variantUUID", controllers.GetVariantByUUID)
	router.GET("/products/variants/", controllers.GetVariants)

	router.POST("/products/variants/", middleware.Authentication(), middleware.CreateVariantAuthorization(), controllers.CreateVariant)
	router.PUT("/products/variants/:variantUUID", middleware.Authentication(), middleware.VariantAuthorization(), controllers.UpdateVariant)
	router.DELETE("/products/variants/:variantUUID", middleware.Authentication(), middleware.VariantAuthorization(), controllers.DeleteVariant)

	router.GET("/products/:productUUID", controllers.GetProductByUUID)
	router.GET("/products/", controllers.GetProducts)

	// router.Use(middleware.Authentication())
	router.POST("/products/", middleware.Authentication(), controllers.CreateProduct)
	router.PUT("/products/:productUUID", middleware.Authentication(), middleware.ProductAuthorization(), controllers.UpdateProduct)
	router.DELETE("/products/:productUUID", middleware.Authentication(), middleware.ProductAuthorization(), controllers.DeleteProduct)
	// }

	return router
}
