package main

import (
	controller "users/Infrastructure/Controller"
	inMemoryUserRepository "users/Infrastructure/Repository"
	"users/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.UserRepo = inMemoryUserRepository.NewUserRepository()

	router := gin.Default()
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUserByID)
	router.POST("/users", controller.PostUsers)
	router.DELETE("/users/:id", controller.DeleteUser)

	router.Run("0.0.0.0:8080")
}
