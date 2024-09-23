package user

type UserRepository interface {
	Save(user User) error
	FindByID(ID string) (User, error)
	FindByEmailAddress(email string) (User, error)
	FindByUsername(username string) (User, error)
	Update(user User) error
	Delete(ID string) error
}

type UserGroupRepository interface {
	Save(userGroup UserGroup) error
	FindByID(ID string) (UserGroup, error)
	Update(userGroup UserGroup) error
	Delete(ID string) error
}
