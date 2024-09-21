package user

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(username, email, password string, firstName string, lastName string) (User, error) {
	u, err := NewUser(username, email, password, firstName, lastName)

	if err != nil {
		return User{}, err
	}

	if err := s.Register(u); err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *UserService) Register(u User) error {
	return s.repo.Save(u)
}

func (s *UserService) FindByID(ID string) (User, error) {
	return s.repo.FindByID(ID)
}

func (s *UserService) FindByEmailAddress(email string) (User, error) {
	return s.repo.FindByEmailAddress(email)
}

func (s *UserService) FindByUsername(username string) (User, error) {
	return s.repo.FindByUsername(username)
}

func (s *UserService) Update(u User) error {
	return s.repo.Update(u)
}

func (s *UserService) Delete(ID string) error {
	return s.repo.Delete(ID)
}
