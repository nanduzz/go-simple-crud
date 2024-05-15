package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nanduzz/go-simple-crud/controller"
	"github.com/nanduzz/go-simple-crud/repository"
	"github.com/nanduzz/go-simple-crud/util"
)

func init() {
	util.LoadEnv()

	controller.Initialize(
		repository.FindAll,
		repository.FindUserById,
		repository.Create,
	)
}

func main() {
	router := gin.Default()

	router.GET("/actuator/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "live, but not an actual health check",
		})
	})

	router.GET("/users", controller.FindAll)
	router.GET("/users/:id", controller.FindUserHandlerById)
	router.POST("/users", controller.CreateUserHandler)

	httpPort := os.Getenv("HTTP_PORT")
	router.Run(fmt.Sprintf(":%s", httpPort))
}
