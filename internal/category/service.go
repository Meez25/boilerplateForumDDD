package category

import "github.com/google/uuid"

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(title, description string, parentID *string) (*Category, error) {
	var parentIDUUID *uuid.UUID

	if parentID != nil {
		id, err := uuid.Parse(*parentID)
		if err != nil {
			return nil, err
		}

		parentIDUUID = &id
	}

	newCategory, err := NewCategory(title, description, parentIDUUID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(*newCategory)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (s *CategoryService) FindByID(ID string) (*Category, error) {
	category, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
