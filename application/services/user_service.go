package services

import (
	"errors"

	"github.com/meez25/boilerplateForumDDD/internal/user"
)

var ErrPasswordConfirmError = errors.New("Password does not match")

type UserService struct {
	userRepo      user.UserRepository
	userGroupRepo user.UserGroupRepository
}

func NewUserService(userRepo user.UserRepository, userGroupRepo user.UserGroupRepository) *UserService {
	return &UserService{
		userRepo:      userRepo,
		userGroupRepo: userGroupRepo,
	}
}

func (s *UserService) Create(username, email, password string, confirmPassword string, firstName string, lastName string, profilePicture string) (user.User, error) {

	if password != confirmPassword {
		return user.User{}, ErrPasswordConfirmError
	}

	u, err := user.NewUser(username, email, password, firstName, lastName, profilePicture)

	if err != nil {
		return user.User{}, err
	}

	if err := s.Register(u); err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (s *UserService) Register(u user.User) error {
	return s.userRepo.Save(u)
}

func (s *UserService) FindByID(ID string) (user.User, error) {
	return s.userRepo.FindByID(ID)
}

func (s *UserService) FindByEmailAddress(email string) (user.User, error) {
	return s.userRepo.FindByEmailAddress(email)
}

func (s *UserService) FindByUsername(username string) (user.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *UserService) Update(u user.User) error {
	return s.userRepo.Update(u)
}

func (s *UserService) Delete(ID string) error {
	return s.userRepo.Delete(ID)
}

func (s *UserService) RegisterGroup(ug user.UserGroup) error {
	return s.userGroupRepo.Save(ug)
}

func (s *UserService) CreateGroup(name, description string, owner user.User) (user.UserGroup, error) {
	ug, err := user.NewUserGroup(name, description, owner)

	if err != nil {
		return user.UserGroup{}, err
	}

	if err := s.RegisterGroup(ug); err != nil {
		return user.UserGroup{}, err
	}

	return ug, nil
}

func (s *UserService) FindGroupByID(ID string) (user.UserGroup, error) {
	return s.userGroupRepo.FindByID(ID)
}
