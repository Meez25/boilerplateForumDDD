package services

import (
	"errors"

	"github.com/meez25/boilerplateForumDDD/internal/authentication"
	"github.com/meez25/boilerplateForumDDD/internal/user"
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

func (ss *AuthenticationService) CreateSession(user user.User) (authentication.Session, error) {
	session := authentication.NewSession(user.EmailAddress, user.ID, user.Username)

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

	_, err = user.Password.CheckPassword(password)

	if err != nil {
		return authentication.Session{}, ErrInvalidCredentials
	}

	session := authentication.NewSession(user.EmailAddress, user.ID, user.Username)

	ss.sessionRepo.Save(*session)

	return *session, nil
}

func (ss *AuthenticationService) Logout(sessionID string) error {
	return ss.sessionRepo.Delete(sessionID)
}
