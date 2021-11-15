package users

import "encoding/json"

type PublicUser  struct {
	Id 			int64 `json:"user_id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Status 		string `json:"status"`
}

type PrivateUser struct {
	Id 			int64 `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email		string `json:"email"`
	Status 		string `json:"status"`
	DateCreated string `json:"date_created"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id: user.Id,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Status: user.Status,
		}
	}

	var result PrivateUser
	userJson, _ := json.Marshal(user)
	json.Unmarshal(userJson, &result)
	return result
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))

	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}

	return result
}