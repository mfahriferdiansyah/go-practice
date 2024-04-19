package controllers

import (
	db "final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appJSON = "application/json"
)

func Register(ctx *gin.Context) {
	db := db.GetDB()
	contentType := helpers.HttpContent(ctx)
	Admin := models.Admin{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	newUUID := uuid.New()
	Admin.UUID = newUUID.String()

	err := db.Debug().Create(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	data := models.AdminResponse{
		ID:    Admin.ID,
		Name:  Admin.Name,
		Email: Admin.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})

}

func Login(ctx *gin.Context) {
	db := db.GetDB()
	contentType := helpers.HttpContent(ctx)
	Admin := models.Admin{}
	var password string

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	password = Admin.Password

	err := db.Debug().Where("email = ?", Admin.Email).Take(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(Admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(Admin.ID, Admin.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
