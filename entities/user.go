package boilerplateforumddd

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	EmailAddress string
	Password     string
	FirstName    string
	LastName     string
	CreatedAt    time.Time
}
