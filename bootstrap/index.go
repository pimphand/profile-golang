package bootstrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"github.com/gin-gonic/gin"
)

func App() {
	database.ConnectDatabase()

	app := gin.Default()

	routes.InitRoute(app)

	err := app.Run(config.Port)
	if err != nil {
		return
	}
}
