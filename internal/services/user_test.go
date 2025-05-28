package services

import (
	"example/api/internal/models"
	"testing"
)

func TestUserService(t *testing.T) {
	// Initialize service
	s := NewUserService()

	// Test Register
	t.Run("Register valid user", func(t *testing.T) {
		id, err := s.Register("Alice", "alice@example.com")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if id != 1 {
			t.Errorf("Expected ID 1, got %d", id)
		}
	})

	t.Run("Register duplicate email", func(t *testing.T) {
		_, err := s.Register("Bob", "alice@example.com")
		if err == nil || err.Error() != "email already exists" {
			t.Errorf("Expected email already exists error, got %v", err)
		}
	})

	t.Run("Register empty fields", func(t *testing.T) {
		_, err := s.Register("", "test@example.com")
		if err == nil || err.Error() != "name and email are required" {
			t.Errorf("Expected required fields error, got %v", err)
		}
	})

	// Test List
	t.Run("List users", func(t *testing.T) {
		users := s.List()
		if len(users) != 1 {
			t.Errorf("Expected 1 user, got %d", len(users))
		}
		expected := models.User{ID: 1, Name: "Alice", Email: "alice@example.com"}
		if users[0] != expected {
			t.Errorf("Expected user %v, got %v", expected, users[0])
		}
	})

	// Test FindByID
	t.Run("Find existing user", func(t *testing.T) {
		user, err := s.FindByID(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := models.User{ID: 1, Name: "Alice", Email: "alice@example.com"}
		if user != expected {
			t.Errorf("Expected user %v, got %v", expected, user)
		}
	})

	t.Run("Find non-existent user", func(t *testing.T) {
		_, err := s.FindByID(999)
		if err == nil || err.Error() != "user not found" {
			t.Errorf("Expected user not found error, got %v", err)
		}
	})

	// Test Delete
	t.Run("Delete existing user", func(t *testing.T) {
		if !s.Delete(1) {
			t.Error("Expected true, got false")
		}
		if len(s.List()) != 0 {
			t.Errorf("Expected 0 users, got %d", len(s.List()))
		}
	})

	t.Run("Delete non-existent user", func(t *testing.T) {
		if s.Delete(999) {
			t.Error("Expected false, got true")
		}
	})
}
