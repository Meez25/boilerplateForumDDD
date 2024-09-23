package authentication

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewSession(t *testing.T) {
	session := NewSession("coucou@example.com")

	if session.ID == uuid.Nil {
		t.Error("NewSession devrait générer un ID non nul")
	}

	if session.keys == nil {
		t.Error("NewSession devrait initialiser la map keys")
	}

	if !session.validUntil.After(time.Now()) {
		t.Error("NewSession devrait définir validUntil dans le futur")
	}
}

func TestAddValue(t *testing.T) {
	session := NewSession("coucou@example.com")
	session.AddValue("test_key", "test_value")

	if value, exists := session.keys["test_key"]; !exists || value != "test_value" {
		t.Error("AddValue n'a pas correctement ajouté la paire clé-valeur")
	}
}

func TestDeleteValue(t *testing.T) {
	session := NewSession("coucou@example.com")
	session.AddValue("test_key", "test_value")
	session.DeleteValue("test_key")

	if _, exists := session.keys["test_key"]; exists {
		t.Error("DeleteValue n'a pas correctement supprimé la clé")
	}
}

func TestGetValue(t *testing.T) {
	session := NewSession("coucou@example.com")
	session.AddValue("test_key", "test_value")

	if value := session.GetValue("test_key"); value != "test_value" {
		t.Errorf("GetValue n'a pas retourné la bonne valeur. Attendu : 'test_value', Obtenu : '%s'", value)
	}

	if value := session.GetValue("non_existent_key"); value != "" {
		t.Errorf("GetValue devrait retourner une chaîne vide pour une clé inexistante. Obtenu : '%s'", value)
	}
}

func TestIsValid(t *testing.T) {
	session := NewSession("coucou@example.com")

	if !session.IsValid() {
		t.Error("Une nouvelle session devrait être valide")
	}

	// Test avec une session expirée
	expiredSession := &Session{
		ID:         uuid.New(),
		validUntil: time.Now().Add(-1 * time.Hour),
	}

	if expiredSession.IsValid() {
		t.Error("Une session expirée ne devrait pas être valide")
	}
}

func TestSessionWithNilKeys(t *testing.T) {
	session := &Session{ID: uuid.New()}

	// Test AddValue avec keys nil
	session.AddValue("test_key", "test_value")
	if session.keys == nil {
		t.Error("AddValue devrait initialiser keys si nil")
	}
	if value, exists := session.keys["test_key"]; !exists || value != "test_value" {
		t.Error("AddValue n'a pas correctement ajouté la paire clé-valeur après initialisation")
	}

	// Test DeleteValue avec keys nil
	session = &Session{ID: uuid.New()}
	session.DeleteValue("test_key") // Ne devrait pas paniquer

	// Test GetValue avec keys nil
	session = &Session{ID: uuid.New()}
	if value := session.GetValue("test_key"); value != "" {
		t.Error("GetValue devrait retourner une chaîne vide si keys est nil")
	}
}
