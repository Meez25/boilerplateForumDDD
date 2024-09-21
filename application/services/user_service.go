package services

import (
	"github.com/meez25/boilerplateForumDDD/internal/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(username, email, password string, firstName string, lastName string) (user.User, error) {
	u, err := user.NewUser(username, email, password, firstName, lastName)

	if err != nil {
		return user.User{}, err
	}

	if err := s.Register(u); err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (s *UserService) Register(u user.User) error {
	return s.repo.Save(u)
}

func (s *UserService) FindByID(ID string) (user.User, error) {
	return s.repo.FindByID(ID)
}

func (s *UserService) FindByEmailAddress(email string) (user.User, error) {
	return s.repo.FindByEmailAddress(email)
}

func (s *UserService) FindByUsername(username string) (user.User, error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) Update(u user.User) error {
	return s.repo.Update(u)
}

func (s *UserService) Delete(ID string) error {
	return s.repo.Delete(ID)
}
