package persistence

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/category"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryMemoryRepo struct {
	categories map[uuid.UUID]category.Category
}

func NewCategoryMemoryRepo() *CategoryMemoryRepo {
	return &CategoryMemoryRepo{
		categories: make(map[uuid.UUID]category.Category),
	}
}

func (r *CategoryMemoryRepo) Save(category category.Category) error {
	r.categories[category.ID] = category
	return nil
}

func (r *CategoryMemoryRepo) FindByID(ID string) (category.Category, error) {
	id, err := uuid.Parse(ID)
	if err != nil {
		return category.Category{}, err
	}

	cat, ok := r.categories[id]
	if !ok {
		return category.Category{}, ErrCategoryNotFound
	}

	return cat, nil
}

func (r *CategoryMemoryRepo) FindAll() ([]category.Category, error) {
	categories := make([]category.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryMemoryRepo) Update(category category.Category) error {
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
