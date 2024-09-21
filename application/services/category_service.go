package services

import (
	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/category"
)

type CategoryService struct {
	repo category.CategoryRepository
}

func NewCategoryService(repo category.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(title, description string, parentID *string) (*category.Category, error) {
	var parentIDUUID *uuid.UUID

	if parentID != nil {
		id, err := uuid.Parse(*parentID)
		if err != nil {
			return nil, err
		}

		parentIDUUID = &id
	}

	newCategory, err := category.NewCategory(title, description, parentIDUUID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(*newCategory)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (s *CategoryService) FindByID(ID string) (*category.Category, error) {
	category, err := s.repo.FindByID(ID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s *CategoryService) AddSubCategory(parentID, subCategoryID string) error {
	parentCategory, err := s.repo.FindByID(parentID)
	if err != nil {
		return err
	}

	subCategory, err := s.repo.FindByID(subCategoryID)
	if err != nil {
		return err
	}

	subCategory.SetParentID(&parentCategory.ID)

	err = s.repo.Save(subCategory)
	if err != nil {
		return err
	}

	return nil
}
