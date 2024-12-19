package main

import (
	"stud-distributor/controllers"
	"stud-distributor/database"
	"stud-distributor/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("host=localhost user=postgres password=postgres dbname=stud_distributor port=5433 sslmode=disable TimeZone=Europe/Moscow")
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
		api.POST("/register/csv-stud", controllers.RegisterUserByCSV) // из файла регаем студентов без выборов специальностей - studs_jinfo
		//api.POST("/update/student", controllers.)
		api.POST("/distribute", controllers.DistributeUser) //отправляем бротиша в группу по его хотению
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
		create := api.Group("/create") //.Use(middlewares.Auth())
		{
			create.POST("/groups", controllers.CreateGroups)
		}
		get := api.Group("/get")
		{
			get.GET("/groups", controllers.GetGroups)
			get.GET("/users", controllers.GetUsers)
		}
		delete := api.Group("/delete")
		{
			delete.DELETE("/groups/:id", controllers.DeleteGroup)
		}

	}
	return router
}
