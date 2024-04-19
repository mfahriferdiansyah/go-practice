package controllers

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := db.GetDB()

	userData := ctx.MustGet("user").(jwt5.MapClaims)

	userID := uint(userData["id"].(float64))

	var ProductReq models.ProductRequest
	if err := ctx.ShouldBind(&ProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := helpers.RemoveExtension(ProductReq.File.Filename)

	uploadResult, err := helpers.UploadFile(ProductReq.File, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUUID := uuid.New()
	Product := models.Product{
		UUID:     newUUID.String(),
		Name:     ProductReq.Name,
		ImageUrl: uploadResult,
		AdminID:  userID,
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}

func UpdateProduct(ctx *gin.Context) {
	db := db.GetDB()

	userData := ctx.MustGet("user").(jwt5.MapClaims)
	contentType := helpers.HttpContent(ctx)
	Product := models.Product{}

	productUUID := ctx.Param("productUUID")
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	var getProduct models.Product
	if err := db.Model(&getProduct).Where("uuid = ?", productUUID).First(&getProduct).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	Product.ID = uint(getProduct.ID)
	Product.AdminID = userID

	updateData := models.Product{
		Name:     Product.Name,
		ImageUrl: Product.ImageUrl,
	}

	if err := db.Model(&Product).Where("uuid = ?", productUUID).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var updatedProduct models.Product
	if err := db.Where("uuid = ?", productUUID).First(&updatedProduct).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not retrieve updated book data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func GetProducts(ctx *gin.Context) {
	db := db.GetDB()

	results := []models.Product{}

	err := db.Debug().Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func GetProductByUUID(ctx *gin.Context) {
	db := db.GetDB()

	results := models.Product{}
	productUUID := ctx.Param("productUUID")

	err := db.Debug().Where("UUID = ?", productUUID).Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := db.GetDB()

	productUUID := ctx.Param("productUUID")

	product := models.Product{}
	if err := db.Debug().Where("UUID = ?", productUUID).Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
}
