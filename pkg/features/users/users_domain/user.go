package users_domain

import (
	"encoding/json"
)

type User struct {
	Id UserId
}

func NewUser(id string) (*User, error) {
	var userId, err = NewUserId(id)
	if err != nil {
		return nil, err
	}

	return &User{
		Id: userId,
	}, nil
}

func (user *User) UnmarshalJSON(data []byte) error {
	var plainUser PlainUser
	json.Unmarshal(data, &plainUser)

	newUser, err := plainUser.ToUser()
	if err != nil {
		return err
	}

	user = newUser

	return nil
}

func (user *User) MarshalJSON() ([]byte, error) {
	var plainUser PlainUser
	plainUser.FromUser(user)
	return json.Marshal(plainUser)
}
