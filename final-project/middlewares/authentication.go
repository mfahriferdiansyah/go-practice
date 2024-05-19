package middleware

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathenticated",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("user", verifyToken)

		userData := ctx.MustGet("user").(jwt5.MapClaims)
		userUUID := userData["uuid"].(string)
		userEmail := userData["email"].(string)

		getUser := models.Admin{}
		db := db.GetDB()
		err = db.Where("email = ? && uuid = ?", userEmail, userUUID).First(&getUser).Error
		if err != nil || getUser.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "Please login first.",
			})
			return
		}

		ctx.Next()
	}
}
