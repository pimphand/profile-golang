package controllers

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"github.com/gin-gonic/gin"
)

func GetProject(ctx *gin.Context) {
	var project []responses.Project
	err := database.DB.Table("projects").Find(&project).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": project,
	})

}

func CreateProject(ctx *gin.Context) {
	var project responses.Project
	err := ctx.BindJSON(&project)
	if err != nil {
		return
	}
	err1 := database.DB.Table("projects").Create(&project).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": project,
	})
}
