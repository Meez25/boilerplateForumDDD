package forum

type CategoryRepository interface {
	Save(category Category) error
	FindByID(ID string) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) error
	Delete(ID string) error
}

type TopicRepository interface {
	Save(topic Topic) error
	FindByID(ID string) (Topic, error)
	FindAll() ([]Topic, error)
	Update(topic Topic) error
	Delete(ID string) error
}
