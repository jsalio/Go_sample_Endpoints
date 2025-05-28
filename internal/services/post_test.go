package services

import (
	"example/api/internal/models"
	"testing"
)

func TestPostService(t *testing.T) {
	// Initialize service
	s := NewPostService()

	// Test Create
	t.Run("Create valid post", func(t *testing.T) {
		id, err := s.Create("Test Post", "This is a test post", 1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if id != 1 {
			t.Errorf("Expected ID 1, got %d", id)
		}
	})

	t.Run("Create post with empty fields", func(t *testing.T) {
		_, err := s.Create("", "This is a test post", 1)
		if err == nil || err.Error() != "title and content are required" {
			t.Errorf("Expected required fields error, got %v", err)
		}

		_, err = s.Create("Test Post", "", 1)
		if err == nil || err.Error() != "title and content are required" {
			t.Errorf("Expected required fields error, got %v", err)
		}
	})

	// Test List
	t.Run("List posts", func(t *testing.T) {
		posts := s.List()
		if len(posts) != 1 {
			t.Errorf("Expected 1 post, got %d", len(posts))
		}
		expected := models.Post{ID: 1, Title: "Test Post", Content: "This is a test post", UserID: 1}
		if posts[0] != expected {
			t.Errorf("Expected post %v, got %v", expected, posts[0])
		}
	})

	// Test FindByID
	t.Run("Find existing post", func(t *testing.T) {
		post, err := s.FindByID(1)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := models.Post{ID: 1, Title: "Test Post", Content: "This is a test post", UserID: 1}
		if post != expected {
			t.Errorf("Expected post %v, got %v", expected, post)
		}
	})

	t.Run("Find non-existent post", func(t *testing.T) {
		_, err := s.FindByID(999)
		if err == nil || err.Error() != "post not found" {
			t.Errorf("Expected post not found error, got %v", err)
		}
	})

	// Test FindByUserID
	t.Run("Find posts by user ID", func(t *testing.T) {
		// Create another post for the same user
		s.Create("Another Post", "This is another test post", 1)

		posts := s.FindByUserID(1)
		if len(posts) != 2 {
			t.Errorf("Expected 2 posts, got %d", len(posts))
		}

		// Create a post for a different user
		s.Create("Different User Post", "This is a post from another user", 2)

		posts = s.FindByUserID(2)
		if len(posts) != 1 {
			t.Errorf("Expected 1 post, got %d", len(posts))
		}
	})

	// Test Delete
	t.Run("Delete existing post", func(t *testing.T) {
		if !s.Delete(1) {
			t.Error("Expected true, got false")
		}
		posts := s.List()
		if len(posts) != 2 {
			t.Errorf("Expected 2 posts after deletion, got %d", len(posts))
		}
	})

	t.Run("Delete non-existent post", func(t *testing.T) {
		if s.Delete(999) {
			t.Error("Expected false, got true")
		}
	})
}
