package entity

import "github.com/google/uuid"

type Type struct {
	ID   uuid.UUID
	Name string
}

func NewType(name string) *Type {
	return &Type{
		ID:   uuid.New(),
		Name: name,
	}
}
