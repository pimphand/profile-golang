package controllers

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/responses"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetExperience(ctx *gin.Context) {
	var experience []responses.Experience
	err := database.DB.Table("experiences").Find(&experience).Error
	if err != nil {
		ctx.JSON(404, gin.H{
			"data": "Data not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": experience,
	})
}

func CreateExperience(ctx *gin.Context) {
	var experience responses.Experience
	err := ctx.BindJSON(&experience)
	if err != nil {
		return
	}
	err1 := database.DB.Table("experiences").Create(&experience).Error
	if err1 != nil {
		ctx.JSON(500, gin.H{
			"data": "Failed to create data",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": experience,
	})
}

func UpdateExperience(ctx *gin.Context) {
	var experience responses.Experience

	// Bind the incoming JSON to the experience struct
	if err := ctx.BindJSON(&experience); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	// Get the ID from the URL parameters
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	// Update the experience in the database
	if err := database.DB.Table("experiences").Where("id = ?", idInt).Save(&experience).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   "Failed to update data",
			"details": err.Error(), // Include error details for debugging purposes
		})
		return
	}

	// Return the updated experience data
	ctx.JSON(200, gin.H{
		"data": experience,
	})
}

func DeleteExperience(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Delete the experience from the database
	result := database.DB.Table("experiences").Where("id = ?", idInt).Delete(&responses.Experience{})
	if err := result.Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   "Failed to delete data",
			"details": err.Error(), // Include error details for debugging purposes
		})
		return
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{
			"error": "Experience not found",
		})
		return
	}

	// Return a success message
	ctx.JSON(200, gin.H{
		"message": "Experience deleted successfully",
	})
}
