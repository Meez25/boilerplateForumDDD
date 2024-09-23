package services

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/authentication"
)

// MockSessionRepository est une implémentation mock de SessionRepository pour les tests
type MockSessionRepository struct {
	sessions map[string]authentication.Session
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{
		sessions: make(map[string]authentication.Session),
	}
}

func (m *MockSessionRepository) Save(session authentication.Session) error {
	m.sessions[session.ID.String()] = session
	return nil
}

func (m *MockSessionRepository) FindByID(id string) (authentication.Session, error) {
	session, ok := m.sessions[id]
	if !ok {
		return authentication.Session{}, errors.New("session not found")
	}
	return session, nil
}

func (m *MockSessionRepository) Update(session authentication.Session) error {
	m.sessions[session.ID.String()] = session
	return nil
}

func (m *MockSessionRepository) Delete(id string) error {
	delete(m.sessions, id)
	return nil
}

func TestCreateSession(t *testing.T) {
	mockRepo := NewMockSessionRepository()
	service := NewSessionService(mockRepo)

	session, err := service.CreateSession()
	if err != nil {
		t.Errorf("Erreur inattendue lors de la création de la session : %v", err)
	}

	if session.ID == uuid.Nil {
		t.Error("L'ID de la session ne devrait pas être nil")
	}

	// Vérifier que la session a bien été sauvegardée dans le repository
	_, err = mockRepo.FindByID(session.ID.String())
	if err != nil {
		t.Errorf("La session n'a pas été correctement sauvegardée : %v", err)
	}
}

func TestGetSessionByID(t *testing.T) {
	mockRepo := NewMockSessionRepository()
	service := NewSessionService(mockRepo)

	// Créer une session de test
	testSession := authentication.NewSession()
	testSession.ID = uuid.New()
	testSession.Email = "test@example.com"
	mockRepo.Save(*testSession)

	// Test de récupération d'une session existante
	retrievedSession, err := service.GetSessionByID(testSession.ID.String())
	if err != nil {
		t.Errorf("Erreur inattendue lors de la récupération de la session : %v", err)
	}

	if retrievedSession.ID != testSession.ID {
		t.Errorf("Les IDs de session ne correspondent pas. Attendu : %v, Obtenu : %v", testSession.ID, retrievedSession.ID)
	}

	if retrievedSession.Email != testSession.Email {
		t.Errorf("Les emails de session ne correspondent pas. Attendu : %v, Obtenu : %v", testSession.Email, retrievedSession.Email)
	}

	// Test de récupération d'une session inexistante
	_, err = service.GetSessionByID("non-existent-id")
	if err == nil {
		t.Error("Une erreur aurait dû être retournée pour une session inexistante")
	}
}
