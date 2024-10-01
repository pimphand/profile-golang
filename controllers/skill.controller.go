package controllers

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"github.com/gin-gonic/gin"
)

func GetEducation(ctx *gin.Context) {
	var education []responses.Education
	err := database.DB.Table("educations").Find(&education).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": education,
	})
}

func CreateEducation(ctx *gin.Context) {
	var education responses.Education
	err := ctx.BindJSON(&education)
	if err != nil {
		return
	}
	err1 := database.DB.Table("educations").Create(&education).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": education,
	})
}

func UpdateEducation(ctx *gin.Context) {
	var education responses.Education
	err := ctx.BindJSON(&education)
	if err != nil {
		return
	}
	err1 := database.DB.Table("educations").Save(&education).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to update data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": education,
	})
}

func DeleteEducation(ctx *gin.Context) {
	var education responses.Education
	err := database.DB.Table("educations").First(&education).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	err1 := database.DB.Table("educations").Delete(&education).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to delete data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": education,
	})
}
