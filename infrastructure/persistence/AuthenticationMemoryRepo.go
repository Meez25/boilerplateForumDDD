package persistence

import (
	"errors"
	"sync"

	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

var ErrSessionNotFound = errors.New("Session not found")

type AuthenticationMemoryRepo struct {
	sessions map[string]authentication.Session
	sync.Mutex
}

func NewSessionMemoryRepo() *AuthenticationMemoryRepo {
	return &AuthenticationMemoryRepo{
		sessions: make(map[string]authentication.Session),
	}
}

func (sm *AuthenticationMemoryRepo) Save(session authentication.Session) error {
	sm.Lock()
	sm.sessions[session.ID.String()] = session
	sm.Unlock()
	return nil
}

func (sm *AuthenticationMemoryRepo) FindByID(id string) (authentication.Session, error) {
	s, ok := sm.sessions[id]
	if !ok {
		return authentication.Session{}, ErrSessionNotFound
	}

	return s, nil
}

func (sm *AuthenticationMemoryRepo) Update(session authentication.Session) error {

	sm.Lock()
	sm.sessions[session.ID.String()] = session
	sm.Unlock()
	return nil
}

func (sm *AuthenticationMemoryRepo) Delete(id string) error {
	sm.Lock()
	delete(sm.sessions, id)
	sm.Unlock()
	return nil
}
