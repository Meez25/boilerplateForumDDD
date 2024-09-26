package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

type CategoryService struct {
	repo forum.CategoryRepository
}

type CategoryAndChildren struct {
	forum.Category
	Children []forum.Category
}

func NewCategoryService(repo forum.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(title, description string, parentID *string) (*forum.Category, error) {
	var parentIDUUID *uuid.UUID

	if parentID != nil {
		id, err := uuid.Parse(*parentID)
		if err != nil {
			return nil, err
		}

		parentIDUUID = &id
	}

	newCategory, err := forum.NewCategory(title, description, parentIDUUID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(*newCategory)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (s *CategoryService) FindByID(ID string) (*forum.Category, error) {
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

func (s *CategoryService) GetAllCategoryAndChildren() ([]CategoryAndChildren, error) {
	var categoriesAndChildren []CategoryAndChildren

	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		if category.ParentID == nil {
			categoriesAndChildren = append(categoriesAndChildren, CategoryAndChildren{
				Category: category,
			})
		} else {
			for i, cat := range categoriesAndChildren {
				if cat.ID == *category.ParentID {
					categoriesAndChildren[i].Children = append(categoriesAndChildren[i].Children, category)
				}
			}
		}
	}

	fmt.Println("retrieve categories", categoriesAndChildren)

	return categoriesAndChildren, nil
}
