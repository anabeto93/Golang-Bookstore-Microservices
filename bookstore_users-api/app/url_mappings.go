package app

import (
	"github.com/anabeto93/bookstore/bookstore_users-api/controllers/ping"
	"github.com/anabeto93/bookstore/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.GET("/users", users.GetAll)
	router.GET("/users/:id", users.Find)
	router.PUT("/users/:id", users.Update)
	router.PATCH("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)
	// internal services
	router.GET("/internal/users/search", users.Search)
}