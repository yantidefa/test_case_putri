package routers

import (
	"log"
	"os"
	"test_case_putri/config"
	userhandler "test_case_putri/handlers/user_handler"

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

	user := v3noauth.Group("/users")
	{
		user.GET("", userhandler.GetAllUserHandler)
		user.GET("/:id", userhandler.GetUserByIdHandler)
		user.POST("", userhandler.InsertUserHandler)
		user.PUT("/:id", userhandler.UpdateUserHandler)
		user.DELETE("/:id", userhandler.DeleteUserHandler)
	}
	

	return r
}

func InitialRouter() {
	port := ":" + os.Getenv("ACTIVE_PORT")
	if err := Routes().Run(port); err != nil {
		log.Fatal(err)
	}
}