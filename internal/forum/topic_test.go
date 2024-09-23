package forum

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewTopic(t *testing.T) {
	validAuthorID := uuid.New()

	tests := []struct {
		name        string
		title       string
		richContent string
		authorID    uuid.UUID
		wantErr     error
	}{
		{"Valid Topic", "Test Title", "Test Content", validAuthorID, nil},
		{"Empty Title", "", "Test Content", validAuthorID, ErrEmptyTitle},
		{"Empty Rich Content", "Test Title", "", validAuthorID, ErrEmptyRichContent},
		{"Empty Author ID", "Test Title", "Test Content", uuid.Nil, ErrEmptyAuthorID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTopic(tt.title, tt.richContent, tt.authorID)
			if err != tt.wantErr {
				t.Errorf("NewTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.ID == uuid.Nil {
					t.Error("NewTopic() got nil ID")
				}
				if got.Title != tt.title {
					t.Errorf("NewTopic() got title = %v, want %v", got.Title, tt.title)
				}
				if got.RichContent != tt.richContent {
					t.Errorf("NewTopic() got richContent = %v, want %v", got.RichContent, tt.richContent)
				}
				if got.AuthorID != tt.authorID {
					t.Errorf("NewTopic() got authorID = %v, want %v", got.AuthorID, tt.authorID)
				}
				if time.Since(got.CreatedAt) > time.Second {
					t.Errorf("NewTopic() CreatedAt not recent enough: %v", got.CreatedAt)
				}
				if time.Since(got.UpdatedAt) > time.Second {
					t.Errorf("NewTopic() UpdatedAt not recent enough: %v", got.UpdatedAt)
				}
				if len(got.Messages) != 0 {
					t.Errorf("NewTopic() got %d messages, want 0", len(got.Messages))
				}
			}
		})
	}
}

func TestTopic_AddMessage(t *testing.T) {
	authorID := uuid.New()
	topic, _ := NewTopic("Test Topic", "Test Content", authorID)

	tests := []struct {
		name        string
		richContent string
		authorID    uuid.UUID
		wantErr     bool
	}{
		{"Valid Message", "Test Message", authorID, false},
		{"Empty Rich Content", "", authorID, true},
		{"Empty Author ID", "Test Message", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := topic.AddMessage(tt.richContent, tt.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Topic.AddMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(topic.Messages) != 1 {
					t.Errorf("Topic.AddMessage() got %d messages, want 1", len(topic.Messages))
				}
				lastMessage := topic.Messages[len(topic.Messages)-1]
				if lastMessage.RichContent != tt.richContent {
					t.Errorf("Topic.AddMessage() got richContent = %v, want %v", lastMessage.RichContent, tt.richContent)
				}
				if lastMessage.AuthorID != tt.authorID {
					t.Errorf("Topic.AddMessage() got authorID = %v, want %v", lastMessage.AuthorID, tt.authorID)
				}
			}
		})
	}
}

func TestTopic_UpdateMessage(t *testing.T) {
	authorID := uuid.New()
	topic, _ := NewTopic("Test Topic", "Test Content", authorID)
	_ = topic.AddMessage("Original Message", authorID)
	messageID := topic.Messages[0].ID

	tests := []struct {
		name        string
		messageID   string
		richContent string
		authorID    uuid.UUID
		wantErr     error
	}{
		{"Valid Update", messageID, "Updated Message", authorID, nil},
		{"Message Not Found", "non-existent-id", "Updated Message", authorID, ErrMessageNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := topic.UpdateMessage(tt.messageID, tt.richContent, tt.authorID)
			if err != tt.wantErr {
				t.Errorf("Topic.UpdateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				updatedMessage := topic.Messages[0]
				if updatedMessage.RichContent != tt.richContent {
					t.Errorf("Topic.UpdateMessage() got richContent = %v, want %v", updatedMessage.RichContent, tt.richContent)
				}
				if time.Since(updatedMessage.UpdatedAt) > time.Second {
					t.Errorf("Topic.UpdateMessage() UpdatedAt not recent enough: %v", updatedMessage.UpdatedAt)
				}
			}
		})
	}
}

func TestTopic_DeleteMessage(t *testing.T) {
	authorID := uuid.New()
	topic, _ := NewTopic("Test Topic", "Test Content", authorID)
	_ = topic.AddMessage("Message to Delete", authorID)
	messageID := topic.Messages[0].ID

	tests := []struct {
		name      string
		messageID string
		wantErr   error
	}{
		{"Valid Delete", messageID, nil},
		{"Message Not Found", "non-existent-id", ErrMessageNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initialMessageCount := len(topic.Messages)
			err := topic.DeleteMessage(tt.messageID)
			if err != tt.wantErr {
				t.Errorf("Topic.DeleteMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if len(topic.Messages) != initialMessageCount-1 {
					t.Errorf("Topic.DeleteMessage() got %d messages, want %d", len(topic.Messages), initialMessageCount-1)
				}
			}
		})
	}
}
