package forum

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          string
	RichContent string
	AuthorID    uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMessage(richContent string, authorID uuid.UUID) (Message, error) {
	id := uuid.New()

	if richContent == "" {
		return Message{}, ErrEmptyRichContent
	}

	if authorID == uuid.Nil {
		return Message{}, ErrEmptyAuthorID
	}

	return Message{
		ID:          id.String(),
		RichContent: richContent,
		AuthorID:    authorID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
