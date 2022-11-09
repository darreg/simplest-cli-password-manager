package model

type Type struct {
	ID       string
	Name     string
	IsBinary bool
}

func NewType(id, name string, binary bool) *Type {
	return &Type{
		ID:       id,
		Name:     name,
		IsBinary: binary,
	}
}
