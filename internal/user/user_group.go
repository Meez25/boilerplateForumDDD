package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyName = errors.New("name can't be empty")
)

type UserGroup struct {
	ID          uuid.UUID
	Name        string
	Description string
	Members     []User
	Owner       User
}

func NewUserGroup(name, description string, owner User) (UserGroup, error) {
	id := uuid.New()

	if name == "" {
		return UserGroup{}, ErrEmptyName
	}

	return UserGroup{
		ID:          id,
		Name:        name,
		Description: description,
		Owner:       owner,
	}, nil
}

func (ug *UserGroup) AddMember(user User) {
	ug.Members = append(ug.Members, user)
}

func (ug *UserGroup) RemoveMember(user User) {
	for i, member := range ug.Members {
		if member.ID == user.ID {
			ug.Members = append(ug.Members[:i], ug.Members[i+1:]...)
			return
		}
	}
}

func (ug *UserGroup) IsMember(user User) bool {
	for _, member := range ug.Members {
		if member.ID == user.ID {
			return true
		}
	}
	return false
}

func (ug *UserGroup) IsOwner(user User) bool {
	return ug.Owner.ID == user.ID
}

func (ug *UserGroup) ChangeOwner(user User) {
	ug.Owner = user
}

func (ug *UserGroup) ChangeName(name string) {
	ug.Name = name
}

func (ug *UserGroup) ChangeDescription(description string) {
	ug.Description = description
}
