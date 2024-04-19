package controllers

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateVariant(ctx *gin.Context) {
	db := db.GetDB()

	Variant := models.Variant{}

	if err := ctx.ShouldBind(&Variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUUID := uuid.New()
	Variant.UUID = newUUID.String()

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
	if err := db.Model(&getVariant).Where("uuid = ?", variantUUID).First(&getVariant).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	Variant.ID = uint(getVariant.ID)

	updateData := models.Variant{
		VariantName: Variant.VariantName,
		Quantity:    Variant.Quantity,
		ProductID:   Variant.ProductID,
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

	err := db.Debug().Preload("Product").Find(&results).Error
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

func GetVariantByUUID(ctx *gin.Context) {
	db := db.GetDB()

	results := models.Variant{}
	variantUUID := ctx.Param("variantUUID")

	err := db.Debug().Preload("Product").Where("UUID = ?", variantUUID).Find(&results).Error
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
