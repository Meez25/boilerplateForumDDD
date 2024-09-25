package persistence

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

type UserMemoryRepository struct {
	users map[uuid.UUID]user.User
	sync.Mutex
}

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyExists = errors.New("Email already exists")

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[uuid.UUID]user.User),
	}
}

func (r *UserMemoryRepository) Save(u user.User) error {
	// Find by email address
	user, _ := r.FindByEmailAddress(u.EmailAddress)

	if user.EmailAddress != "" {
		return ErrEmailAlreadyExists
	}

	r.Lock()
	r.users[u.ID] = u
	r.Unlock()
	return nil
}

func (r *UserMemoryRepository) FindByID(ID string) (user.User, error) {
	for _, u := range r.users {
		if u.ID.String() == ID {
			return u, nil
		}
	}
	return user.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) FindByEmailAddress(email string) (user.User, error) {
	for _, u := range r.users {
		if u.EmailAddress == email {
			return u, nil
		}
	}
	return user.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) FindByUsername(username string) (user.User, error) {
	for _, u := range r.users {
		if u.Username == username {
			return u, nil
		}
	}
	return user.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) Update(u user.User) error {
	r.Lock()
	r.users[u.ID] = u
	r.Unlock()
	return nil
}

func (r *UserMemoryRepository) Delete(ID string) error {
	r.Lock()
	delete(r.users, uuid.MustParse(ID))
	r.Unlock()
	return nil
}
