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
		userID := uint(user["id"].(float64))

		var getProduct models.Product
		err := db.Debug().Where("UUID = ?", productUUID).Find(&getProduct).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}

		if uint(getProduct.AdminID) != userID {
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
		userID := uint(user["id"].(float64))

		var getVariant models.Variant
		err := db.Select("variant").Preload("product").Where("UUID = ?", variantUUID).Find(&getVariant).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}

		if uint(getVariant.Product.AdminID) != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}
