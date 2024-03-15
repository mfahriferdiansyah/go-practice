package router

import (
	"go-gin-gorm-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	r := gin.Default()

	orderRouter := r.Group("/orders")
	{
		orderRouter.POST("/", controllers.CreateOrder)
		orderRouter.GET("/", controllers.GetOrders)
		orderRouter.GET("/:id", controllers.GetOrderById)
		orderRouter.PUT("/:id", controllers.UpdateOrder)
		orderRouter.DELETE("/:id", controllers.DeleteOrder)
	}

	return r
}
