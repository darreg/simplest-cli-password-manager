package entity

import "github.com/google/uuid"

type Type struct {
	ID       uuid.UUID
	Name     string
	IsBinary bool
}

func NewType(name string, binary bool) *Type {
	return &Type{
		ID:       uuid.New(),
		Name:     name,
		IsBinary: binary,
	}
}
