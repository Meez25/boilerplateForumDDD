package services_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

// MockCategoryRepository est un mock de CategoryRepository
type MockCategoryRepository struct {
	categories map[string]forum.Category
}

func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{
		categories: make(map[string]forum.Category),
	}
}

func (m *MockCategoryRepository) Save(category forum.Category) error {
	m.categories[category.ID.String()] = category
	return nil
}

func (m *MockCategoryRepository) FindByID(ID string) (forum.Category, error) {
	category, ok := m.categories[ID]
	if !ok {
		return forum.Category{}, errors.New("category not found")
	}
	return category, nil
}

func (m *MockCategoryRepository) Update(category forum.Category) error {
	m.categories[category.ID.String()] = category
	return nil
}

func (m *MockCategoryRepository) Delete(ID string) error {
	delete(m.categories, ID)
	return nil
}

func (m *MockCategoryRepository) FindAll() ([]forum.Category, error) {
	var categories []forum.Category
	for _, category := range m.categories {
		categories = append(categories, category)
	}
	return categories, nil
}

func TestCategoryService_Create(t *testing.T) {
	mockRepo := NewMockCategoryRepository()
	service := services.NewCategoryService(mockRepo)

	t.Run("Création réussie sans parent", func(t *testing.T) {
		title := "Test Category"
		description := "Test Description"

		category, err := service.Create(title, description, nil)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if category == nil {
			t.Error("La catégorie ne devrait pas être nil")
		}
		if category.Title != title {
			t.Errorf("Titre attendu %s, obtenu %s", title, category.Title)
		}
		if category.Description != description {
			t.Errorf("Description attendue %s, obtenue %s", description, category.Description)
		}
		if category.ParentID != nil {
			t.Error("ParentID devrait être nil")
		}
	})

	t.Run("Création réussie avec parent", func(t *testing.T) {
		title := "Sub Category"
		description := "Sub Description"
		parentID := uuid.New().String()

		category, err := service.Create(title, description, &parentID)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if category == nil {
			t.Error("La catégorie ne devrait pas être nil")
		}
		if category.Title != title {
			t.Errorf("Titre attendu %s, obtenu %s", title, category.Title)
		}
		if category.Description != description {
			t.Errorf("Description attendue %s, obtenue %s", description, category.Description)
		}
		if category.ParentID == nil {
			t.Error("ParentID ne devrait pas être nil")
		}
	})
}

func TestCategoryService_FindByID(t *testing.T) {
	mockRepo := NewMockCategoryRepository()
	service := services.NewCategoryService(mockRepo)

	t.Run("Catégorie trouvée", func(t *testing.T) {
		id := uuid.New()
		expectedCategory := forum.Category{ID: id, Title: "Test"}
		mockRepo.Save(expectedCategory)

		category, err := service.FindByID(id.String())

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if category == nil {
			t.Error("La catégorie ne devrait pas être nil")
		}
		if !reflect.DeepEqual(*category, expectedCategory) {
			t.Errorf("Catégorie attendue %+v, obtenue %+v", expectedCategory, *category)
		}
	})

	t.Run("Catégorie non trouvée", func(t *testing.T) {
		id := uuid.New().String()

		category, err := service.FindByID(id)

		if err == nil {
			t.Error("Une erreur était attendue")
		}
		if category != nil {
			t.Error("La catégorie devrait être nil")
		}
	})
}

func TestCategoryService_AddSubCategory(t *testing.T) {
	mockRepo := NewMockCategoryRepository()
	service := services.NewCategoryService(mockRepo)

	t.Run("Ajout réussi d'une sous-catégorie", func(t *testing.T) {
		parentID := uuid.New()
		subCategoryID := uuid.New()
		parentCategory := forum.Category{ID: parentID}
		subCategory := forum.Category{ID: subCategoryID}

		mockRepo.Save(parentCategory)
		mockRepo.Save(subCategory)

		err := service.AddSubCategory(parentID.String(), subCategoryID.String())

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		updatedSubCategory, _ := mockRepo.FindByID(subCategoryID.String())
		if updatedSubCategory.ParentID == nil || *updatedSubCategory.ParentID != parentID {
			t.Error("Le ParentID de la sous-catégorie n'a pas été mis à jour correctement")
		}
	})

	t.Run("Erreur lors de la recherche de la catégorie parente", func(t *testing.T) {
		parentID := uuid.New().String()
		subCategoryID := uuid.New().String()

		err := service.AddSubCategory(parentID, subCategoryID)

		if err == nil {
			t.Error("Une erreur était attendue")
		}
	})

	t.Run("Erreur lors de la recherche de la sous-catégorie", func(t *testing.T) {
		parentID := uuid.New()
		subCategoryID := uuid.New().String()
		parentCategory := forum.Category{ID: parentID}

		mockRepo.Save(parentCategory)

		err := service.AddSubCategory(parentID.String(), subCategoryID)

		if err == nil {
			t.Error("Une erreur était attendue")
		}
	})
}

