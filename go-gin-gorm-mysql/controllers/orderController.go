package controllers

import (
	"go-gin-gorm-mysql/database"
	"go-gin-gorm-mysql/models"
	"net/http"

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

	orderId := c.Param("id")

	Order := models.Order{}
	c.ShouldBindJSON(&Order)

	if err := db.Where("id = ?", orderId).Updates(&Order).Error; err != nil {
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
