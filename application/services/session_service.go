package services

import (
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

type SessionService struct {
	sessionRepo authentication.SessionRepository
}

func NewSessionService(sessionRepo authentication.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: sessionRepo,
	}
}

func (ss *SessionService) CreateSession() (authentication.Session, error) {
	session := authentication.NewSession()

	ss.sessionRepo.Save(*session)

	return *session, nil
}

func (ss *SessionService) GetSessionByID(sessionID string) (authentication.Session, error) {
	session, err := ss.sessionRepo.FindByID(sessionID)

	if err != nil {
		return authentication.Session{}, err
	}

	return session, nil
}
