package services

import (
	"fmt"
	"strings"

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

func UpdateUser(userId int64, user users.User, isPartialUpdate bool) (*users.User, *errors.RestErr) {
	if isPartialUpdate {
		if err := user.ValidatePatch(); err != nil {
			return nil, err
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
	}
	var userDTO users.User
	existingUser, err := userDTO.Find(userId); if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d not found.", userId))
	}

	if isPartialUpdate {
		if strings.TrimSpace(user.Email) != "" {
			existingUser.Email = user.Email
		}
		if strings.TrimSpace(user.FirstName) != "" {
			existingUser.FirstName = user.FirstName
		}
		if strings.TrimSpace(user.LastName) != "" {
			existingUser.LastName = user.LastName
		}
		if strings.TrimSpace(user.DateCreated) != "" {
			existingUser.DateCreated = user.DateCreated
		}
	} else {
		existingUser.Email = user.Email
		existingUser.DateCreated = user.DateCreated
		existingUser.FirstName = user.FirstName
		existingUser.LastName = user.LastName
	}

	_, err = existingUser.Update(); if err != nil {
		return nil, err;
	}

	return existingUser, nil
}

func GetAllUsers() ([]users.User, *errors.RestErr) {
	var userDTO users.User
	users, err := userDTO.GetAll(); if err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	var userDTO users.User
	existingUser, err := userDTO.Find(userId); if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found.", userId))
	}

	if err = existingUser.Destroy(); err != nil {
		return err
	}

	return nil
}