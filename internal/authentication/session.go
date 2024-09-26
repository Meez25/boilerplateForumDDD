package authentication

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID
	Email      string
	UserID     uuid.UUID
	Username   string
	keys       map[string]string
	ValidUntil time.Time
}

func NewSession(email string, userID uuid.UUID, username string) *Session {
	return &Session{
		ID:         uuid.New(),
		Email:      email,
		UserID:     userID,
		Username:   username,
		keys:       make(map[string]string),
		ValidUntil: time.Now().Add(24 * time.Hour),
	}
}

func (s *Session) AddValue(key string, value string) {
	if s.keys == nil {
		s.keys = make(map[string]string)
	}
	s.keys[key] = value
}

func (s *Session) DeleteValue(key string) {
	delete(s.keys, key)
}

func (s *Session) GetValue(key string) string {
	return s.keys[key]
}

func (s *Session) IsValid() bool {
	return time.Now().Before(s.ValidUntil)
}

func (s *Session) GetValidUntil() time.Time {
	return s.ValidUntil
}
