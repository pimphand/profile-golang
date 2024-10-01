package controllers

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	var profile responses.Profile
	err := database.DB.Table("profiles").First(&profile).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": profile,
	})
}

func CreateProfile(ctx *gin.Context) {
	var profile responses.Profile
	err := ctx.BindJSON(&profile)
	if err != nil {
		return
	}
	err1 := database.DB.Table("profiles").Create(&profile).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": profile,
	})
}
