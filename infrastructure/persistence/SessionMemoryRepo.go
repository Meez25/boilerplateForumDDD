package persistence

import (
	"errors"
	"sync"

	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

var ErrSessionNotFound = errors.New("Session not found")

type SessionMemoryRepo struct {
	sessions map[string]authentication.Session
	sync.Mutex
}

func NewSessionMemoryRepo() *SessionMemoryRepo {
	return &SessionMemoryRepo{
		sessions: make(map[string]authentication.Session),
	}
}

func (sm *SessionMemoryRepo) Save(session authentication.Session) {
	sm.Lock()
	sm.sessions[session.ID.String()] = session
	sm.Unlock()
}

func (sm *SessionMemoryRepo) FindByID(id string) (authentication.Session, error) {
	s, ok := sm.sessions[id]
	if !ok {
		return authentication.Session{}, ErrSessionNotFound
	}

	return s, nil
}

func (sm *SessionMemoryRepo) Update(session authentication.Session) error {

	sm.Lock()
	sm.sessions[session.ID.String()] = session
	sm.Unlock()
	return nil
}

func (sm *SessionMemoryRepo) Delete(id string) error {
	sm.Lock()
	delete(sm.sessions, id)
	sm.Unlock()
	return nil
}
