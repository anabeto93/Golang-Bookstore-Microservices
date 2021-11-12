package users

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/anabeto93/bookstore/bookstore_users-api/datasources/mysql/users_db"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/date_utils"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

const (
	insertQuery = "INSERT INTO users (email, first_name, last_name, date_created) VALUES (?, ?, ?, ?)"
	getUserQuery = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ? LIMIT 1;"
	getUserByEmailQuery = "SELECT id, first_name, last_name, email, date_created FROM users WHERE email = ? LIMIT 1;"

)

func (user User) Find(userId int64) (*User, *errors.RestErr) {
	var result User

	stmnt, err := prepare(getUserQuery); if  err != nil {
		return nil, err
	}

	defer stmnt.Close()

	sqlUser := stmnt.QueryRow(userId)
	fmt.Println(sqlUser)
	if sqlErr := sqlUser.Scan(&result.Id, &result.Email, &result.FirstName, &result.LastName, &result.DateCreated); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Error fetching user: %s", sqlErr.Error()))
	}

	if (result == User{}) {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
	}

	return &result, nil
}

func (user User) FindByEmail(email string) (*User, *errors.RestErr) {
	var result User

	stmnt, err := prepare(getUserByEmailQuery); if  err != nil {
		return nil, err
	}

	defer stmnt.Close()

	sqlUser := stmnt.QueryRow(email)
	if sqlErr := sqlUser.Scan(&result.Id, &result.Email, &result.FirstName, &result.LastName, &result.DateCreated); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Error fetching user: %s", sqlErr.Error()))
	}

	if (result != User{}) {
		return &result, nil
	}

	return nil, errors.NewNotFoundError(fmt.Sprintf("user with email %s not found", email))
}

func (user *User) Save() *errors.RestErr {
	stmnt, err := prepare(insertQuery); if err != nil {
		return err
	}

	defer stmnt.Close()

	user.DateCreated = date_utils.GetNowString("2006-01-02 15:04:05")
	result, sqlErr := stmnt.Exec(user.Email, user.FirstName, user.LastName, user.DateCreated)

	if sqlErr != nil {
		if strings.Contains(sqlErr.Error(), "Duplicate entry") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))			
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user: %s", sqlErr.Error()))
	}

	userId, sqlErr := result.LastInsertId(); if sqlErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error saving user: %s", sqlErr.Error()))
	}
	
	user.Id = userId
	return nil
}

func prepare(query string) (*sql.Stmt, *errors.RestErr) {
	stmnt, err := users_db.Client.Prepare(query); if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Could not prepare sql statement: %s", err.Error()))
	}

	return stmnt, nil
}