package services

import (
	"errors"
	"example/api/internal/models"
)

// PostService manages post-related operations such as creation, listing, finding, and deleting posts.
// It maintains an in-memory collection of posts and handles post ID generation.
type PostService struct {
	posts  []models.Post
	nextId int
}

// NewPostService creates and returns a new instance of PostService with initialized fields.
func NewPostService() *PostService {
	return &PostService{
		posts:  make([]models.Post, 0),
		nextId: 1,
	}
}

// Create creates a new post with the given title, content, and user ID.
// Returns the new post's ID and an error if creation fails.
// Creation fails if title or content is empty, or if the user ID doesn't exist.
func (s *PostService) Create(title string, content string, userID int) (int, error) {
	if title == "" || content == "" {
		return 0, errors.New("title and content are required")
	}

	post := models.Post{
		ID:      s.nextId,
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	s.posts = append(s.posts, post)
	s.nextId++
	return post.ID, nil
}

// List returns all posts.
func (s *PostService) List() []models.Post {
	return s.posts
}

// FindByID searches for a post by its ID.
// Returns the post if found, or an error if no post exists with the given ID.
func (s *PostService) FindByID(id int) (models.Post, error) {
	for _, p := range s.posts {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Post{}, errors.New("post not found")
}

// FindByUserID returns all posts for a specific user.
// Returns an empty slice if no posts are found.
func (s *PostService) FindByUserID(userID int) []models.Post {
	var userPosts []models.Post
	for _, p := range s.posts {
		if p.UserID == userID {
			userPosts = append(userPosts, p)
		}
	}
	return userPosts
}

// Delete removes a post with the specified ID from the service.
// Returns true if the post was found and deleted, false otherwise.
func (s *PostService) Delete(id int) bool {
	for i, p := range s.posts {
		if p.ID == id {
			s.posts = append(s.posts[:i], s.posts[i+1:]...)
			return true
		}
	}
	return false
}
