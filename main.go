package main

import (
	"stud-distributor/controllers"
	"stud-distributor/database"
	"stud-distributor/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("host=localhost user=postgres password=postgres dbname=go_jwt_gorm port=5433 sslmode=disable TimeZone=Europe/Moscow")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/token/refresh", controllers.RefreshToken)
		api.POST("/upload/csv-stud", controllers.ProcessCSVEndpoint) // из файла регаем студентов
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
