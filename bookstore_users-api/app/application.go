package app

import (
	"github.com/anabeto93/bookstore/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("About starting application...")
	router.Run(":6080")
}