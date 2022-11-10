package model

type User struct {
	ID    string
	Name  string
	Login string
}

func NewUser(id, name, login string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Login: login,
	}
}
