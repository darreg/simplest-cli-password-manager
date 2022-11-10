package entity

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	Name         string
	Login        string
	PasswordHash string
}

func NewUser(name, login, passwordHash string) *User {
	return &User{
		ID:           uuid.New(),
		Name:         name,
		Login:        login,
		PasswordHash: passwordHash,
	}
}
