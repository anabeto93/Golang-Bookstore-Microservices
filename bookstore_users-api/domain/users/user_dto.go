package users

import (
	"strings"
	"time"

	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email	string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if (user.Email == "") {
		return errors.NewBadRequestError("invalid email address")
	}

	// check that the datetime is valid as well
	if strings.TrimSpace(user.DateCreated) != "" {
		// meaning a datetime has been provided
		_, err := time.Parse("", user.DateCreated); if err != nil {
			return errors.NewBadRequestError("invalid datetime value")			
		}
	}

	return nil
}