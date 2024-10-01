package controllers

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"github.com/gin-gonic/gin"
)

func GetSkill(ctx *gin.Context) {
	var skill []responses.Skill
	err := database.DB.Table("skills").Find(&skill).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": skill,
	})
}

func CreateSkill(ctx *gin.Context) {
	var skill responses.Skill
	err := ctx.BindJSON(&skill)
	if err != nil {
		return
	}
	err1 := database.DB.Table("skills").Create(&skill).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": skill,
	})
}

func UpdateSkill(ctx *gin.Context) {
	var skill responses.Skill
	err := ctx.BindJSON(&skill)
	if err != nil {
		return
	}
	err1 := database.DB.Table("skills").Save(&skill).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to update data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": skill,
	})
}

func DeleteSkill(ctx *gin.Context) {
	var skill responses.Skill
	err := database.DB.Table("skills").First(&skill).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	err1 := database.DB.Table("skills").Delete(&skill).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to delete data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": skill,
	})
}
