package users

import (
	"fmt"

	"github.com/anabeto93/bookstore/bookstore_users-api/datasources/mysql/users_db"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/date_utils"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)

const (
	insertQuery = "INSERT INTO users (email, first_name, last_name, date_created) VALUES (?, ?, ?, ?)"
)

func (user User) Find(userId int64) (*User, *errors.RestErr) {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[userId]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
	}
	return result, nil
}

func (user User) FindByEmail(email string) (*User, *errors.RestErr) {
	var result *User

	for _, v := range usersDB {
		if v.Email == email {
			result = v
		} else {
			result = nil
		}
	}

	if result != nil {
		return result, nil
	}

	return nil, errors.NewNotFoundError(fmt.Sprintf("user with email %s not found", email))
}

func (user *User) Save() *errors.RestErr {
	stmnt, err := users_db.Client.Prepare(insertQuery); if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Could not prepare sql insert statement: %s", err.Error()))
	}
	defer stmnt.Close()

	user.DateCreated = date_utils.GetNowString("2006-01-02 15:04:05")
	result, err := stmnt.Exec(user.Email, user.FirstName, user.LastName, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user: %s", err.Error()))
	}

	userId, err := result.LastInsertId(); if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user: %s", err.Error()))
	}
	
	user.Id = userId
	return nil
}