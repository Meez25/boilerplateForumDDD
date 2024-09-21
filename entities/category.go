package boilerplateforumddd

import (
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
