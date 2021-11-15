package users

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/anabeto93/bookstore/bookstore_users-api/datasources/mysql/users_db"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/date_utils"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
	"github.com/anabeto93/bookstore/bookstore_users-api/utils/mysql_utils"
)

const (
	insertQuery = "INSERT INTO users (email, first_name, last_name, password, status, date_created) VALUES (?, ?, ?, ?, ?, ?)"
	getUserQuery = "SELECT id, first_name, last_name, email, password, status, date_created FROM users WHERE id = ? LIMIT 1;"
	getUserByEmailQuery = "SELECT id, first_name, last_name, email, password, status, date_created FROM users WHERE email = ? LIMIT 1;"
	allUsersQuery = "SELECT id, first_name, last_name, email, password, status, date_created FROM users LIMIT 1000;"
	updateUserQuery = "UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ?, status = ?, date_created = ? WHERE id = ?"
	deleteQuery = "DELETE FROM users WHERE id = ?;"
	getByStatus = "SELECT id, first_name, last_name, email, password, status, date_created FROM users WHERE status = ? LIMIT 1000;"
)

func (user User) GetAll() ([]User, *errors.RestErr) {
	var users = []User{}

	stmnt, err := prepare(allUsersQuery); if err != nil {
		return nil, err
	}

	defer stmnt.Close()

	sqlUsers, sqlErr := stmnt.Query(); if sqlErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Could not fetch users: %s", sqlErr.Error()))
	}
	defer sqlUsers.Close()
	
	for sqlUsers.Next() {
		var usr User
		if sqlErr := sqlUsers.Scan(&usr.Id, &usr.Email, &usr.FirstName, &usr.LastName, &usr.Password, &usr.Status, &usr.DateCreated); sqlErr != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("Could not fetch users: %s", sqlErr.Error()))
		}
		users = append(users, usr)
	}
	return users, nil
}

func (user User) Find(userId int64) (*User, *errors.RestErr) {
	var result User

	stmnt, err := prepare(getUserQuery); if  err != nil {
		return nil, err
	}

	defer stmnt.Close()

	sqlUser := stmnt.QueryRow(userId)

	if sqlErr := sqlUser.Scan(&result.Id, &result.FirstName, &result.LastName, &result.Email, &result.Password, &result.Status, &result.DateCreated); err != nil {
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
	if sqlErr := sqlUser.Scan(&result.Id, &result.Email, &result.FirstName, &result.LastName, &result.Password, &result.Status, &result.DateCreated); err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Error fetching user: %s", sqlErr.Error()))
	}

	if (result != User{}) {
		return &result, nil
	}

	return nil, errors.NewNotFoundError(fmt.Sprintf("user with email %s not found", email))
}

func (user User) FindByStatus(status string) ([]User, *errors.RestErr) {
	var users = []User{}

	stmnt, err := prepare(getByStatus); if err != nil {
		return nil, err
	}

	defer stmnt.Close()

	sqlUsers, sqlErr := stmnt.Query(status); if sqlErr != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Could not fetch users: %s", sqlErr.Error()))
	}
	defer sqlUsers.Close()
	
	for sqlUsers.Next() {
		var usr User
		if sqlErr := sqlUsers.Scan(&usr.Id, &usr.Email, &usr.FirstName, &usr.LastName, &usr.Password, &usr.Status, &usr.DateCreated); sqlErr != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("Could not fetch users: %s", sqlErr.Error()))
		}
		users = append(users, usr)
	}
	return users, nil
}

func (user *User) Save() *errors.RestErr {
	stmnt, err := prepare(insertQuery); if err != nil {
		return err
	}

	defer stmnt.Close()

	user.DateCreated = date_utils.GetNowString("2006-01-02 15:04:05")
	result, sqlErr := stmnt.Exec(user.Email, user.FirstName, user.LastName, user.Password, user.Status, user.DateCreated)

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

func (user *User) Update() (int64, *errors.RestErr) {
	stmnt, err := prepare(updateUserQuery); if err != nil {
		return 0, err
	}

	defer stmnt.Close()

	fname := user.FirstName
	lname := user.LastName
	email := user.Email
	date_created := user.DateCreated
	pwd := user.Password
	status := user.Status

	result, sqlErr := stmnt.Exec(fname, lname, email, pwd, status, date_created, user.Id);
	if sqlErr != nil {
		return 0, mysql_utils.ParseError("email", email, sqlErr, "Error updating user")
	}

	rows, sqlErr := result.RowsAffected(); if sqlErr != nil {
		return 0, mysql_utils.ParseError("email", email, sqlErr, "Error updating user")
	}

	return rows, nil
}

func (user *User) Destroy() *errors.RestErr {
	stmnt, err := prepare(deleteQuery); if err != nil {
		return err
	}

	defer stmnt.Close()

	_, sqlErr := stmnt.Exec(user.Id); if sqlErr != nil {
		fmt.Println(fmt.Sprintf("Error deleting user: %s", sqlErr.Error()))
		return mysql_utils.ParseError("id", strconv.Itoa(int(user.Id)), sqlErr, "Error deleting user")
	}
	return nil
}

func prepare(query string) (*sql.Stmt, *errors.RestErr) {
	stmnt, err := users_db.Client.Prepare(query); if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Could not prepare sql statement: %s", err.Error()))
	}

	return stmnt, nil
}