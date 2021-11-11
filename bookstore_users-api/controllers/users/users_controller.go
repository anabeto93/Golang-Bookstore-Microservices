package users

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetUsers(c *gin.Context) {
}

func FindUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64); if err != nil {
		res := errors.NewBadRequestError("invalid user id")
		c.JSON(int(res.Status), res)
		return
	}
	user, findErr := services.FindUser(userId); if findErr != nil {
		c.JSON(int(findErr.Status), findErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64); if err != nil {
		res := errors.NewBadRequestError("invalid user id")
		c.JSON(int(res.Status), err)
		return
	}
	var payload users.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println(err.Error())
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user, updateErr := services.UpdateUser(userId, payload); if updateErr != nil {
		c.JSON(int(updateErr.Status), updateErr)
		return
	}
	c.JSON(http.StatusOK, user)
}