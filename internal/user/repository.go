package user

type UserRepository interface {
	Save(user User) error
	FindByID(ID string) (User, error)
	FindByEmailAddress(email string) (User, error)
	FindByUsername(username string) (User, error)
	Update(user User) error
	Delete(ID string) error
}
