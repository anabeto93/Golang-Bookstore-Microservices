package users

import (
	"fmt"
	"net/http"

	"github.com/anabeto93/bookstore/bookstore_users-api/domain/users"
	"github.com/anabeto93/bookstore/bookstore_users-api/services"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	createdUser, err := services.CreateUser(user); if err != nil {
		fmt.Println(err)
		c.JSON(int(err.Status), err)
		return
	}
	fmt.Println("Created user", createdUser)
	c.JSON(http.StatusCreated, createdUser)
}

func GetUsers(c *gin.Context) {}

func FindUser(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}