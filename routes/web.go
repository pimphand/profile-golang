package routes

import (
	"gin-gonic-gorm/controllers/product_controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Update this to the origin of your Vue.js app
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Optional: cache preflight response for 12 hours
	}))
	//login
	route.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})

	//product
	route.GET("/products", product_controller.GetAllProducts)
}
