package services

import (
	"errors"

	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthenticationService struct {
	sessionRepo authentication.SessionRepository
	userService UserService
}

func NewAuthenticationService(sessionRepo authentication.SessionRepository, userService *UserService) *AuthenticationService {
	return &AuthenticationService{
		sessionRepo: sessionRepo,
		userService: *userService,
	}
}

func (ss *AuthenticationService) CreateSession(email string) (authentication.Session, error) {
	session := authentication.NewSession(email)

	ss.sessionRepo.Save(*session)

	return *session, nil
}

func (ss *AuthenticationService) GetSessionByID(sessionID string) (authentication.Session, error) {
	session, err := ss.sessionRepo.FindByID(sessionID)

	if err != nil {
		return authentication.Session{}, err
	}

	return session, nil
}

func (ss *AuthenticationService) Authenticate(email, password string) (authentication.Session, error) {
	user, err := ss.userService.FindByEmailAddress(email)

	if err != nil {
		return authentication.Session{}, err
	}

	if !user.CheckPassword(password) {
		return authentication.Session{}, ErrInvalidCredentials
	}

	session := authentication.NewSession(email)

	ss.sessionRepo.Save(*session)

	return *session, nil
}

func (ss *AuthenticationService) Logout(sessionID string) error {
	return ss.sessionRepo.Delete(sessionID)
}
