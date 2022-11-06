package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	LoginTime    *time.Time
	LastSeenTime *time.Time
}

func NewSession(userID uuid.UUID, loginTime, lastSeenTime *time.Time) *Session {
	return &Session{
		ID:           uuid.New(),
		UserID:       userID,
		LoginTime:    loginTime,
		LastSeenTime: lastSeenTime,
	}
}
