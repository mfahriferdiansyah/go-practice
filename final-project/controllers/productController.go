package controllers

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func CreateProduct(ctx *gin.Context) {
	db := db.GetDB()

	userData := ctx.MustGet("user").(jwt5.MapClaims)

	userUUID := userData["uuid"].(string)

	var ProductReq models.ProductRequest
	if err := ctx.ShouldBind(&ProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ProductReq.File == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "File is required.",
		})
		return
	}

	fileName := helpers.RemoveExtension(ProductReq.File.Filename)

	uploadResult, err := helpers.UploadFile(ProductReq.File, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Product := models.ProductCreation{
		Name:      ProductReq.Name,
		ImageUrl:  uploadResult,
		AdminUUID: userUUID,
	}

	err = db.Debug().Omit("Admin", "Variants").Create(&Product).Error
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

	var ProductReq models.ProductRequest
	if err := ctx.ShouldBind(&ProductReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productUUID := ctx.Param("productUUID")
	var getProduct models.Product
	if err := db.Model(&getProduct).Where("uuid = ?", productUUID).First(&getProduct).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	imageUrl := getProduct.ImageUrl
	if ProductReq.File != nil {
		fileName := helpers.RemoveExtension(ProductReq.File.Filename)

		uploadResult, err := helpers.UploadFile(ProductReq.File, fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageUrl = uploadResult
	}

	productName := getProduct.Name
	if ProductReq.Name != "" {
		productName = ProductReq.Name
	}

	if ProductReq.File == nil && ProductReq.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Please fill atleast 1 column.",
		})
	}

	updateData := models.Product{
		Name:     productName,
		ImageUrl: imageUrl,
	}

	if err := db.Where("uuid = ?", productUUID).Updates(&updateData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Error in Updating product.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product updated!",
	})
}

func GetProducts(ctx *gin.Context) {
	db := db.GetDB()

	offset, limit := helpers.GetPagination(ctx)
	name := ctx.Query("search")

	results := []models.Product{}

	query := db.Debug().Limit(limit).Offset(offset).Preload("Admin").Preload("Variants")

	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%")
	}

	err := query.Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	if len(results) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": fmt.Sprintf("No products found matching the search query '%s'", name),
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

	err := db.Debug().Preload("Admin").Preload("Variants").Where("UUID = ?", productUUID).Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	if results.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": "No product found with UUID: " + productUUID,
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
