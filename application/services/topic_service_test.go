package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

// MockTopicRepository est un mock de TopicRepository
type MockTopicRepository struct {
	topics map[string]forum.Topic
}

func NewMockTopicRepository() *MockTopicRepository {
	return &MockTopicRepository{
		topics: make(map[string]forum.Topic),
	}
}

func (m *MockTopicRepository) Save(topic forum.Topic) error {
	m.topics[topic.ID.String()] = topic
	return nil
}

func (m *MockTopicRepository) Update(topic forum.Topic) error {
	m.topics[topic.ID.String()] = topic
	return nil
}

func (m *MockTopicRepository) FindByID(ID string) (forum.Topic, error) {
	topic, ok := m.topics[ID]
	if !ok {
		return forum.Topic{}, errors.New("topic not found")
	}
	return topic, nil
}

func (m *MockTopicRepository) FindAll() ([]forum.Topic, error) {
	var topics []forum.Topic
	for _, topic := range m.topics {
		topics = append(topics, topic)
	}
	return topics, nil
}

func (m *MockTopicRepository) Delete(ID string) error {
	delete(m.topics, ID)
	return nil
}

func TestTopicService_CreateTopic(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Création réussie d'un topic", func(t *testing.T) {
		title := "Test Topic"
		richContent := "Test Content"
		authorID := uuid.New()

		topic, err := service.CreateTopic(title, richContent, authorID)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if topic.Title != title {
			t.Errorf("Titre attendu %s, obtenu %s", title, topic.Title)
		}
		if topic.RichContent != richContent {
			t.Errorf("Contenu attendu %s, obtenu %s", richContent, topic.RichContent)
		}
		if topic.AuthorID != authorID {
			t.Errorf("AuthorID attendu %s, obtenu %s", authorID, topic.AuthorID)
		}
	})
}

func TestTopicService_AddMessage(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Ajout réussi d'un message", func(t *testing.T) {
		topic, _ := forum.NewTopic("Test Topic", "Initial Content", uuid.New())
		mockRepo.Save(topic)

		newContent := "New Message"
		authorID := uuid.New()

		err := service.AddMessage(topic.ID.String(), newContent, authorID)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		updatedTopic, _ := mockRepo.FindByID(topic.ID.String())
		if len(updatedTopic.Messages) != 1 {
			t.Errorf("Nombre de messages attendu 1, obtenu %d", len(updatedTopic.Messages))
		}
		if updatedTopic.Messages[0].RichContent != newContent {
			t.Errorf("Contenu attendu %s, obtenu %s", newContent, updatedTopic.Messages[0].RichContent)
		}
	})
}

func TestTopicService_GetTopicByID(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Récupération réussie d'un topic", func(t *testing.T) {
		topic, _ := forum.NewTopic("Test Topic", "Content", uuid.New())
		mockRepo.Save(topic)

		retrievedTopic, err := service.GetTopicByID(topic.ID.String())

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if !reflect.DeepEqual(topic, retrievedTopic) {
			t.Errorf("Topic récupéré ne correspond pas au topic original")
		}
	})

	t.Run("Topic non trouvé", func(t *testing.T) {
		_, err := service.GetTopicByID(uuid.New().String())

		if err == nil {
			t.Error("Une erreur était attendue")
		}
	})
}

func TestTopicService_GetAllTopics(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Récupération de tous les topics", func(t *testing.T) {
		topic1, _ := forum.NewTopic("Topic 1", "Content 1", uuid.New())
		topic2, _ := forum.NewTopic("Topic 2", "Content 2", uuid.New())
		mockRepo.Save(topic1)
		mockRepo.Save(topic2)

		topics, err := service.GetAllTopics()

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}
		if len(topics) != 2 {
			t.Errorf("Nombre de topics attendu 2, obtenu %d", len(topics))
		}
	})
}

func TestTopicService_UpdateMessage(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Mise à jour réussie d'un message", func(t *testing.T) {
		topic, _ := forum.NewTopic("Test Topic", "Initial Content", uuid.New())
		mockRepo.Save(topic)

		// Add message to topic
		topic.AddMessage("Initial Message", uuid.New())
		mockRepo.Update(topic)

		newContent := "Updated Content"
		err := service.UpdateMessage(topic.ID.String(), topic.Messages[0].ID, newContent, topic.Messages[0].AuthorID)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		updatedTopic, _ := mockRepo.FindByID(topic.ID.String())
		if updatedTopic.Messages[0].RichContent != newContent {
			t.Errorf("Contenu attendu %s, obtenu %s", newContent, updatedTopic.Messages[0].RichContent)
		}
	})
}

func TestTopicService_DeleteTopic(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Suppression réussie d'un topic", func(t *testing.T) {
		topic, _ := forum.NewTopic("Test Topic", "Content", uuid.New())
		mockRepo.Save(topic)

		err := service.DeleteTopic(topic.ID.String())

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		_, err = mockRepo.FindByID(topic.ID.String())
		if err == nil {
			t.Error("Le topic devrait être supprimé")
		}
	})
}

func TestTopicService_DeleteMessage(t *testing.T) {
	mockRepo := NewMockTopicRepository()
	service := NewTopicService(mockRepo)

	t.Run("Suppression réussie d'un message", func(t *testing.T) {
		topic, _ := forum.NewTopic("Test Topic", "Initial Content", uuid.New())
		topic.AddMessage("Second Message", uuid.New())
		mockRepo.Save(topic)

		err := service.DeleteMessage(topic.ID.String(), topic.Messages[0].ID)

		if err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		updatedTopic, _ := mockRepo.FindByID(topic.ID.String())
		if len(updatedTopic.Messages) != 0 {
			t.Errorf("Nombre de messages attendu 0, obtenu %d", len(updatedTopic.Messages))
		}
	})
}
