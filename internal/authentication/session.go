package authentication

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID         uuid.UUID
	Email      string
	UserID     string
	Username   string
	keys       map[string]string
	validUntil time.Time
}

func NewSession() *Session {
	return &Session{
		ID:         uuid.New(),
		keys:       make(map[string]string),
		validUntil: time.Now().Add(24 * time.Hour),
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
	return time.Now().Before(s.validUntil)
}
