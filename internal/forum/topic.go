package forum

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle       = errors.New("title can't be empty")
	ErrEmptyRichContent = errors.New("rich content can't be empty")
	ErrEmptyAuthorID    = errors.New("author ID can't be empty")
)

type Topic struct {
	ID          uuid.UUID
	Title       string
	RichContent string
	AuthorID    uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTopic(title, richContent string, authorID uuid.UUID) (Topic, error) {
	id := uuid.New()

	if title == "" {
		return Topic{}, ErrEmptyTitle
	}

	if richContent == "" {
		return Topic{}, ErrEmptyRichContent
	}

	if authorID == uuid.Nil {
		return Topic{}, ErrEmptyAuthorID
	}

	return Topic{
		ID:          id,
		Title:       title,
		RichContent: richContent,
		AuthorID:    authorID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
