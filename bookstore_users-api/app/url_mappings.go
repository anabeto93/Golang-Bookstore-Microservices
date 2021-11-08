package app

import (
	"github.com/anabeto93/bookstore/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.FindUser)
	router.PUT("/users/:id", controllers.UpdateUser);
}