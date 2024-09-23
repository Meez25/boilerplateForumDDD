package persistence

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

var ErrUserGroupNotFound = errors.New("user group not found")

type UserGroupMemoryRepo struct {
	UserGroups map[uuid.UUID]user.UserGroup
}

func NewUserGroupMemoryRepo() *UserGroupMemoryRepo {
	return &UserGroupMemoryRepo{
		UserGroups: make(map[uuid.UUID]user.UserGroup),
	}
}

func (r *UserGroupMemoryRepo) Save(ug user.UserGroup) error {
	r.UserGroups[ug.ID] = ug
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
	r.UserGroups[ug.ID] = ug
	return nil
}

func (r *UserGroupMemoryRepo) Delete(ID string) error {
	delete(r.UserGroups, uuid.MustParse(ID))
	return nil
}
