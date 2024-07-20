package user_domain

import "encoding/json"

type User struct {
	Id string `json:"id"`
}

func (user *User) ToJson() ([]byte, error) {
	return json.Marshal(user)
}
