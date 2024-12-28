package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"stud-distributor/controllers"
	"stud-distributor/database"
	"stud-distributor/middlewares"
)

func main() {
	// Initialize Database
	database.Connect("host=localhost user=postgres password=postgres dbname=stud_distributor port=5433 sslmode=disable TimeZone=Europe/Moscow")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8081")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/token", controllers.GenerateToken)
			auth.POST("/token/refresh", controllers.RefreshToken)

		}
		api.POST("/register/user", controllers.RegisterUser)
		api.POST("/register/csv-stud", controllers.RegisterUserByCSV) // из файла регаем студентов без выборов специальностей - studs_jinfo
		//api.POST("/update/student", controllers.)
		api.POST("/distribute", controllers.DistributeUser) //отправляем бротиша в группу по его хотению
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/user/:id", controllers.GetUser) //by id
			secured.POST("/user", controllers.GetUserByEmail)
			secured.GET("/ping", controllers.Ping)
		}
		create := api.Group("/create") //.Use(middlewares.Auth())
		{
			create.POST("/groups", controllers.CreateGroups)
		}
		get := api.Group("/get") //.Use(middlewares.Auth())
		{
			get.GET("/groups", controllers.GetGroups)
			get.GET("/users", controllers.GetUsers)
			get.GET("/group/:id", controllers.GetGroupById)
			get.GET("/specialities", controllers.GetSpecialityNames)
		}
		delete := api.Group("/delete")
		{
			delete.DELETE("/groups/:id", controllers.DeleteGroup) //.Use(middlewares.Auth())
		}
	}
	return router
}
