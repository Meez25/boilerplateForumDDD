package forum

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	CategoryID uuid.UUID
	TopicID    uuid.UUID
	AuthorID   uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Messages   []Message
}
