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
	Password     Password
	FirstName    string
	LastName     string
	ProfilePic   string
	CreatedAt    time.Time
	SuperAdmin   bool
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

	user := User{
		ID:           id,
		Username:     username,
		EmailAddress: email,
		FirstName:    firstName,
		ProfilePic:   profilePicture,
		LastName:     lastName,
		CreatedAt:    time.Now(),
		SuperAdmin:   false,
	}

	user.setPassword(password)

	return user, nil
}

func (u *User) setPassword(password string) {
	encryptedPassword, _ := generateFromPassword(password)

	u.Password = Password{Password: encryptedPassword}
}

func (u *User) GiveSuperAdmin() {
	u.SuperAdmin = true
}

func (u *User) RemoveSuperAdmin() {
	u.SuperAdmin = false
}
