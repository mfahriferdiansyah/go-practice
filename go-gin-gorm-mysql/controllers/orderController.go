package controllers

import (
	"go-gin-gorm-mysql/database"
	"go-gin-gorm-mysql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()

	Order := models.Order{}
	c.ShouldBindJSON(&Order)

	err := db.Debug().Create(&Order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order Created!", "data": Order})
}

func GetOrders(c *gin.Context) {
	db := database.GetDB()

	Orders := []models.Order{}
	if err := db.Debug().Preload("Items").Find(&Orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request Success!", "data": Orders})
}

func GetOrderById(c *gin.Context) {
	db := database.GetDB()

	orderId := c.Param("id")

	Order := models.Order{}
	if err := db.Debug().Preload("Items").First(&Order, orderId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request Success!", "data": Order})
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()

	orderIdStr := c.Param("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid order ID"})
		return
	}

	Order := models.Order{}
	c.ShouldBindJSON(&Order)

	Order.ID = uint(orderId)

	for i := range Order.Items {
		Order.Items[i].OrderID = uint(orderId)
	}

	if err := db.Debug().Where("order_id = ?", orderId).Delete(models.Item{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := db.Debug().Updates(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Update Success!"})
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()

	orderId := c.Param("id")

	Order := models.Order{}
	if err := db.Debug().Where("id = ?", orderId).Delete(&Order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Delete Success!"})
}
