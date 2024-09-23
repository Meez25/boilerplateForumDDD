package persistence

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

var ErrUserGroupNotFound = errors.New("user group not found")

type UserGroupMemoryRepo struct {
	UserGroups map[uuid.UUID]user.UserGroup
	sync.Mutex
}

func NewUserGroupMemoryRepo() *UserGroupMemoryRepo {
	return &UserGroupMemoryRepo{
		UserGroups: make(map[uuid.UUID]user.UserGroup),
	}
}

func (r *UserGroupMemoryRepo) Save(ug user.UserGroup) error {
	r.Lock()
	r.UserGroups[ug.ID] = ug
	r.Unlock()
	return nil
}

func (r *UserGroupMemoryRepo) FindByID(ID string) (user.UserGroup, error) {
	for _, ug := range r.UserGroups {
		if ug.ID.String() == ID {
			return ug, nil
		}
	}
	return user.UserGroup{}, ErrUserGroupNotFound
}

func (r *UserGroupMemoryRepo) Update(ug user.UserGroup) error {
	r.Lock()
	r.UserGroups[ug.ID] = ug
	r.Unlock()
	return nil
}

func (r *UserGroupMemoryRepo) Delete(ID string) error {
	r.Lock()
	delete(r.UserGroups, uuid.MustParse(ID))
	r.Unlock()
	return nil
}
