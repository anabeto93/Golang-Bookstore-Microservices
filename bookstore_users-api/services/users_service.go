package services

import (
	"fmt"

	"github.com/anabeto93/bookstore/bookstore_users-api/domain/users"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// check if user already exists
	existingUser, err := user.Find(user.Id); if err != nil {
		if err.Status != 404 {
			return nil, err
		}
	}

	if existingUser != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("User with id %d already exists", user.Id))
	}

	existingUser, err = user.FindByEmail(user.Email); if err != nil {
		if err.Status != 404 {
			return nil, err
		}
	}

	if existingUser != nil {
		return nil, errors.NewBadRequestError(fmt.Sprintf("User with email %s already exists", user.Email))
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUser(userId int64) (*users.User, *errors.RestErr) {
	var userDTO users.User
	user, err := userDTO.Find(userId); if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d already exists", userId))
	}

	return user, nil
}

func UpdateUser(userId int64, user users.User) (*users.User, *errors.RestErr) {
	return nil, nil
}