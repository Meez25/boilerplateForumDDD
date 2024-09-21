package infrastructure

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

type UserMemoryRepository struct {
	users map[uuid.UUID]user.User
}

var ErrUserNotFound = errors.New("user not found")

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[uuid.UUID]user.User),
	}
}

func (r *UserMemoryRepository) Save(u user.User) error {
	r.users[u.ID] = u
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
	r.users[u.ID] = u
	return nil
}

func (r *UserMemoryRepository) Delete(ID string) error {
	delete(r.users, uuid.MustParse(ID))
	return nil
}
