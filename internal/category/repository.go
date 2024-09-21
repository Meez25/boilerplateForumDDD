package category

type CategoryRepository interface {
	Save(category Category) error
	FindByID(ID string) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) error
	Delete(ID string) error
}
