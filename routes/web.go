package routes

import (
	"gin-gonic-gorm/controllers/product_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	//login
	route.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})

	//product
	route.GET("/products", product_controller.GetAllProducts)
}
