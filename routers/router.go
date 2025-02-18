package routers

import (
	"log"
	"os"
	"test_case_putri/config"
	taskhandler "test_case_putri/handlers/task_handler"
	userhandler "test_case_putri/handlers/user_handler"
	"test_case_putri/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	config.LoadEnv()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v3noauth := r.Group("/api")

	auth := v3noauth.Group("/auth")
	{
		auth.POST("login", userhandler.LoginHandler)
		auth.POST("logout", userhandler.LogoutHandler)
	}

	user := v3noauth.Group("/users").Use(middlewares.Authentication())
	{
		user.POST("login", userhandler.LoginHandler)
		user.POST("logout", userhandler.LogoutHandler)
		user.GET("", userhandler.GetUsersHandler)
		user.GET("/:id", userhandler.GetUserByIdHandler)
		user.POST("", userhandler.InsertUserHandler)
		user.PUT("/:id", userhandler.UpdateUserHandler)
		user.DELETE("/:id", userhandler.DeleteUserHandler)
	}

	task := v3noauth.Group("/tasks").Use(middlewares.Authentication())
	{
		task.GET("", taskhandler.GetTasksHandler)
		task.GET("/:id", taskhandler.GetTaskByIdHandler)
		task.GET("by-user/:user_id", taskhandler.GetTaskByUserIdHandler)
		task.POST("", taskhandler.InsertTaskHandler)
		task.PUT("/:id", taskhandler.UpdateTaskHandler)
		task.DELETE("/:id", taskhandler.DeleteTaskHandler)
	}

	return r
}

func InitialRouter() {
	port := ":" + os.Getenv("ACTIVE_PORT")
	if err := Routes().Run(port); err != nil {
		log.Fatal(err)
	}
}
