package forum

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewCategory(t *testing.T) {
	parentID := uuid.New()

	tests := []struct {
		name        string
		title       string
		description string
		parentID    *uuid.UUID
		wantErr     error
	}{
		{"Valid Category", "Test Title", "Test Description", &parentID, nil},
		{"Empty Title", "", "Test Description", &parentID, ErrEmptyCategoryTitle},
		{"Empty Description", "Test Title", "", &parentID, ErrorEmptyCategoryDescription},
		{"No Parent", "Test Title", "Test Description", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCategory(tt.title, tt.description, tt.parentID)
			if err != tt.wantErr {
				t.Errorf("NewCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.ID == uuid.Nil {
					t.Error("NewCategory() got nil ID")
				}
				if got.Title != tt.title {
					t.Errorf("NewCategory() got title = %v, want %v", got.Title, tt.title)
				}
				if got.Description != tt.description {
					t.Errorf("NewCategory() got description = %v, want %v", got.Description, tt.description)
				}
				if tt.parentID == nil && got.ParentID != nil {
					t.Errorf("NewCategory() got parentID = %v, want nil", *got.ParentID)
				} else if tt.parentID != nil && (got.ParentID == nil || *got.ParentID != *tt.parentID) {
					t.Errorf("NewCategory() got parentID = %v, want %v", *got.ParentID, *tt.parentID)
				}
				if time.Since(got.CreatedAt) > time.Second {
					t.Errorf("NewCategory() CreatedAt not recent enough: %v", got.CreatedAt)
				}
			}
		})
	}
}

func TestCategory_Update(t *testing.T) {
	category, _ := NewCategory("Initial Title", "Initial Description", nil)

	tests := []struct {
		name        string
		title       string
		description string
		wantErr     error
	}{
		{"Valid Update", "New Title", "New Description", nil},
		{"Empty Title", "", "New Description", ErrEmptyTitle},
		{"Empty Description", "New Title", "", ErrorEmptyCategoryDescription},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := category.Update(tt.title, tt.description)
			if err != tt.wantErr {
				t.Errorf("Category.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if category.Title != tt.title {
					t.Errorf("Category.Update() got title = %v, want %v", category.Title, tt.title)
				}
				if category.Description != tt.description {
					t.Errorf("Category.Update() got description = %v, want %v", category.Description, tt.description)
				}
			}
		})
	}
}

func TestCategory_SetParentID(t *testing.T) {
	category, _ := NewCategory("Test Title", "Test Description", nil)
	parentID := uuid.New()

	tests := []struct {
		name     string
		parentID *uuid.UUID
		want     *uuid.UUID
	}{
		{"Set Parent", &parentID, &parentID},
		{"Clear Parent", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category.SetParentID(tt.parentID)
			got := category.GetParentID()
			if tt.want == nil && got != nil {
				t.Errorf("Category.SetParentID() = %v, want nil", *got)
			} else if tt.want != nil && (got == nil || *got != *tt.want) {
				t.Errorf("Category.SetParentID() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestCategory_Getters(t *testing.T) {
	parentID := uuid.New()
	category, _ := NewCategory("Test Title", "Test Description", &parentID)

	t.Run("GetID", func(t *testing.T) {
		if got := category.GetID(); got == uuid.Nil {
			t.Errorf("Category.GetID() = %v, want non-nil UUID", got)
		}
	})

	t.Run("GetTitle", func(t *testing.T) {
		if got := category.GetTitle(); got != "Test Title" {
			t.Errorf("Category.GetTitle() = %v, want %v", got, "Test Title")
		}
	})

	t.Run("GetDescription", func(t *testing.T) {
		if got := category.GetDescription(); got != "Test Description" {
			t.Errorf("Category.GetDescription() = %v, want %v", got, "Test Description")
		}
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		if got := category.GetCreatedAt(); time.Since(got) > time.Second {
			t.Errorf("Category.GetCreatedAt() = %v, want recent time", got)
		}
	})
}

func TestCategory_IsRootAndIsChild(t *testing.T) {
	parentID := uuid.New()

	tests := []struct {
		name      string
		parentID  *uuid.UUID
		wantRoot  bool
		wantChild bool
	}{
		{"Root Category", nil, true, false},
		{"Child Category", &parentID, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, _ := NewCategory("Test Title", "Test Description", tt.parentID)
			if got := category.IsRoot(); got != tt.wantRoot {
				t.Errorf("Category.IsRoot() = %v, want %v", got, tt.wantRoot)
			}
			if got := category.IsChild(); got != tt.wantChild {
				t.Errorf("Category.IsChild() = %v, want %v", got, tt.wantChild)
			}
		})
	}
}
