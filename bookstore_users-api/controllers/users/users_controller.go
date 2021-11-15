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

func isPublic(c *gin.Context) bool {
	return c.GetHeader("X-Public") == "true"
}

func getUserId(id string) (int64, *errors.RestErr) {
	uId, err := strconv.ParseInt(id, 10, 64); if err != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return uId, nil
}

func Create(c *gin.Context) {
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
	c.JSON(http.StatusCreated, createdUser.Marshall(isPublic(c)))
}

func GetAll(c *gin.Context) {
	users, err := services.GetAllUsers(); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(isPublic(c)))
}

func Find(c *gin.Context) {
	userId, err := getUserId(c.Param("id")); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	user, err := services.FindUser(userId); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(isPublic(c)))
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("id")); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	var payload users.User
	if bindErr := c.ShouldBindJSON(&payload); bindErr != nil {
		fmt.Println(bindErr.Error())
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	user, updateErr := services.UpdateUser(userId, payload, isPartial); if updateErr != nil {
		c.JSON(int(updateErr.Status), updateErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(isPublic(c)))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("id")); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status); if err != nil {
		c.JSON(int(err.Status), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(isPublic(c)))
}