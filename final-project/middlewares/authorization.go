package middleware

import (
	db "final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		productUUID := ctx.Param("productUUID")

		user := ctx.MustGet("user").(jwt5.MapClaims)
		userUUID := user["uuid"]

		var getProduct models.Product
		err := db.Debug().Preload("Admin").Preload("Variants").Where("UUID = ?", productUUID).Find(&getProduct).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		if getProduct.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data Not Found",
			})
			return
		}

		if getProduct.AdminUUID != userUUID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		variantUUID := ctx.Param("variantUUID")

		user := ctx.MustGet("user").(jwt5.MapClaims)
		userUUID := user["uuid"]

		var getVariant models.VariantAuth
		err := db.Where("UUID = ?", variantUUID).Preload("Product").Find(&getVariant).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}

		if getVariant.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Variant Not Found",
			})
			return
		}

		if getVariant.Product.AdminUUID != userUUID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func CreateVariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.GetDB()
		Variant := models.VariantCreation{}
		if err := ctx.ShouldBind(&Variant); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var getProduct models.Product
		err := db.Debug().Preload("Admin").Preload("Variants").Where("UUID = ?", Variant.ProductUUID).Find(&getProduct).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		if getProduct.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Product Not Found",
			})
			return
		}

		userData := ctx.MustGet("user").(jwt5.MapClaims)
		userUUID := userData["uuid"].(string)

		if userUUID != getProduct.Admin.UUID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
