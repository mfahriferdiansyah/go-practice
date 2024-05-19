package controllers

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateVariant(ctx *gin.Context) {
	db := db.GetDB()

	Variant := models.VariantCreation{}
	if err := ctx.ShouldBind(&Variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}

func UpdateVariant(ctx *gin.Context) {
	db := db.GetDB()

	contentType := helpers.HttpContent(ctx)
	Variant := models.Variant{}

	variantUUID := ctx.Param("variantUUID")

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Variant)
	} else {
		ctx.ShouldBind(&Variant)
	}

	var getVariant models.Variant
	if err := db.Where("uuid = ?", variantUUID).First(&getVariant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	// Variant.ID = uint(getVariant.ID)

	updateData := models.Variant{
		VariantName: Variant.VariantName,
		Quantity:    Variant.Quantity,
		ProductUUID: Variant.ProductUUID,
	}

	if err := db.Model(&Variant).Where("uuid = ?", variantUUID).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var updatedVariant models.Variant
	if err := db.Where("uuid = ?", variantUUID).Find(&updatedVariant).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not retrieve updated book data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedVariant,
	})
}

func GetVariants(ctx *gin.Context) {
	db := db.GetDB()
	results := []models.Variant{}

	offset, limit := helpers.GetPagination(ctx)
	name := ctx.Query("search")

	query := db.Debug().Limit(limit).Offset(offset)

	if name != "" {
		query = query.Where("LOWER(variant_name) LIKE ?", "%"+strings.ToLower(name)+"%")
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

func GetVariantByUUID(ctx *gin.Context) {
	db := db.GetDB()

	results := models.Variant{}
	variantUUID := ctx.Param("variantUUID")

	err := db.Debug().Where("UUID = ?", variantUUID).Find(&results).Error
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
			"message": "No variant found with UUID: " + variantUUID,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func DeleteVariant(ctx *gin.Context) {
	db := db.GetDB()

	variantUUID := ctx.Param("variantUUID")

	variant := models.Variant{}
	if err := db.Debug().Where("uuid = ?", variantUUID).Delete(&variant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete success",
	})
}
