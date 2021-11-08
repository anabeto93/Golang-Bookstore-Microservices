package app

import (
	"github.com/anabeto93/bookstore/bookstore_users-api/controllers/ping"
	"github.com/anabeto93/bookstore/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users", users.GetUsers)
	router.GET("/users/:id", users.FindUser)
	router.PUT("/users/:id", users.UpdateUser);
}