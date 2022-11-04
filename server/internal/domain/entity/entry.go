package entity

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TypeID    uuid.UUID
	Name      string
	Metadata  string
	Data      []byte
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewEntry(userID, typeID uuid.UUID, name, metadata string, data []byte, createdAt, updatedAt *time.Time) *Entry {
	return &Entry{
		ID:        uuid.New(),
		UserID:    userID,
		TypeID:    typeID,
		Name:      name,
		Metadata:  metadata,
		Data:      data,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
