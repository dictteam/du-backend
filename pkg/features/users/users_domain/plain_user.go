package users_domain

// PlainUser is a DTO used to marshal and unmarshal user to json
type PlainUser struct {
	Id string `json:"id"`
}

// FromUser converts the user values into a plain user since User fields are
// always valid we can fill PlainUser fields directly with user fields
func (p *PlainUser) FromUser(user *User) {
	p.Id = user.Id.Value()
}

func (p *PlainUser) ToUser() (*User, error) {
	newUser, err := NewUser(p.Id)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
