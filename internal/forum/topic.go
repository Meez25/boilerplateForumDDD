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
	ErrMessageNotFound  = errors.New("message not found")
)

type Topic struct {
	ID          uuid.UUID
	Title       string
	RichContent string
	AuthorID    uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Messages    []Message
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
		Messages:    []Message{},
	}, nil
}

func (t *Topic) AddMessage(richContent string, authorID uuid.UUID) error {
	message, err := NewMessage(richContent, authorID)
	if err != nil {
		return err
	}

	t.Messages = append(t.Messages, message)

	return nil
}

func (t *Topic) UpdateMessage(messageID string, richContent string, authorID uuid.UUID) error {
	for i, message := range t.Messages {
		if message.ID == messageID {
			t.Messages[i].RichContent = richContent
			t.Messages[i].UpdatedAt = time.Now()
			return nil
		}
	}

	return ErrMessageNotFound
}

func (t *Topic) DeleteMessage(messageID string) error {
	for i, message := range t.Messages {
		if message.ID == messageID {
			t.Messages = append(t.Messages[:i], t.Messages[i+1:]...)
			return nil
		}
	}

	return ErrMessageNotFound
}
