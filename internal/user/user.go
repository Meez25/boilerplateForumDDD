package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	EmailAddress string
	Password     string
	FirstName    string
	LastName     string
	ProfilePic   string
	CreatedAt    time.Time
}

var (
	ErrEmptyUsername = errors.New("username can't be empty")
	ErrEmptyEmail    = errors.New("email can't be empty")
	ErrEmptyPassword = errors.New("password can't be empty")
)

func NewUser(username, email, password, firstName, lastName string, profilePicture string) (User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return User{}, err
	}

	if username == "" {
		return User{}, ErrEmptyUsername
	}

	if email == "" {
		return User{}, ErrEmptyEmail
	}

	if password == "" {
		return User{}, ErrEmptyPassword
	}

	return User{
		ID:           id,
		Username:     username,
		EmailAddress: email,
		Password:     password,
		FirstName:    firstName,
		ProfilePic:   profilePicture,
		LastName:     lastName,
		CreatedAt:    time.Now(),
	}, nil
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
