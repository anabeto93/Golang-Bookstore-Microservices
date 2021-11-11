package users

import (
	"fmt"

	"github.com/anabeto93/bookstore/bookstore_users-api/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)
func (user User) Find(userId int64) (*User, *errors.RestErr) {
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
	usersDB[user.Id] = user;
	return nil
}