package routes

import (
	"gin-gonic-gorm/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://vuejs.test"}, // Correct origin or add more if necessary
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"}, // Add more headers if you need to expose them
		AllowCredentials: true,                                        // For credentials, specific origins must be used, not "*"
		MaxAge:           12 * 3600,                                   // Cache preflight response for 12 hours
	}))
	//login
	route.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})

	route.GET("/profile", controllers.GetProfile)
	route.POST("/profile/create", controllers.CreateProfile)

	//education
	route.GET("/educations", controllers.GetEducation)
	route.POST("/educations", controllers.CreateEducation)

	//experience
	route.GET("/experiences", controllers.GetExperience)
	route.POST("/experiences", controllers.CreateExperience)
	route.PUT("/experiences", controllers.UpdateExperience)
	route.DELETE("/experiences", controllers.DeleteExperience)

	//project
	route.GET("/projects", controllers.GetProject)
	route.POST("/projects", controllers.CreateProject)

	//skill
	route.GET("/skills", controllers.GetSkill)
	route.POST("/skills", controllers.CreateSkill)
}
