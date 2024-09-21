package persistence

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryMemoryRepo struct {
	categories map[uuid.UUID]forum.Category
}

func NewCategoryMemoryRepo() *CategoryMemoryRepo {
	return &CategoryMemoryRepo{
		categories: make(map[uuid.UUID]forum.Category),
	}
}

func (r *CategoryMemoryRepo) Save(category forum.Category) error {
	r.categories[category.ID] = category
	return nil
}

func (r *CategoryMemoryRepo) FindByID(ID string) (forum.Category, error) {
	id, err := uuid.Parse(ID)
	if err != nil {
		return forum.Category{}, err
	}

	cat, ok := r.categories[id]
	if !ok {
		return forum.Category{}, ErrCategoryNotFound
	}

	return cat, nil
}

func (r *CategoryMemoryRepo) FindAll() ([]forum.Category, error) {
	categories := make([]forum.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryMemoryRepo) Update(category forum.Category) error {
	r.categories[category.ID] = category
	return nil
}

func (r *CategoryMemoryRepo) Delete(ID string) error {
	id, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	delete(r.categories, id)
	return nil
}
