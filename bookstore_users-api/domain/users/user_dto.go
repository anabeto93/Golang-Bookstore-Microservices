package users

import (
	"strings"
	"time"

	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)
type User struct {
	Id 			int64 `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email		string `json:"email"`
	Status 		string `json:"status"`
	Password	string `json:"password"`
	DateCreated string `json:"date_created"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Status = strings.TrimSpace(strings.ToLower(user.Status))
	user.Password = strings.TrimSpace(user.Password)

	if (user.Email == "") {
		return errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	// check that the datetime is valid as well
	if strings.TrimSpace(user.DateCreated) != "" {
		// meaning a datetime has been provided
		_, err := time.Parse("2006-01-02 15:04:05", user.DateCreated); if err != nil {
			return errors.NewBadRequestError("invalid datetime value")			
		}
	}

	return nil
}

func (user *User) ValidatePatch() *errors.RestErr {
	// check that the datetime is valid as well
	if strings.TrimSpace(user.DateCreated) != "" {
		// meaning a datetime has been provided
		_, err := time.Parse("2006-01-02 15:04:05", user.DateCreated); if err != nil {
			return errors.NewBadRequestError("invalid datetime value")			
		}
	}

	return nil
}