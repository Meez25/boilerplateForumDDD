package category

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID
	Title       string
	Description string
	ParentID    *uuid.UUID
	CreatedAt   time.Time
}

var (
	ErrEmptyTitle         = errors.New("title can't be empty")
	ErrorEmptyDescription = errors.New("description can't be empty")
)

func NewCategory(title, description string, parentID *uuid.UUID) (*Category, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	if description == "" {
		return nil, ErrorEmptyDescription
	}

	return &Category{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		ParentID:    parentID,
		CreatedAt:   time.Now(),
	}, nil
}

func (c *Category) Update(title, description string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	if description == "" {
		return ErrorEmptyDescription
	}

	c.Title = title
	c.Description = description

	return nil
}

func (c *Category) SetParentID(parentID *uuid.UUID) {
	if parentID == nil {
		c.ParentID = nil
		return
	}

	c.ParentID = parentID
}

func (c *Category) GetParentID() *uuid.UUID {
	return c.ParentID
}

func (c *Category) GetID() uuid.UUID {
	return c.ID
}

func (c *Category) GetTitle() string {
	return c.Title
}

func (c *Category) GetDescription() string {
	return c.Description
}

func (c *Category) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Category) IsRoot() bool {
	return c.ParentID == nil
}

func (c *Category) IsChild() bool {
	return c.ParentID != nil
}
