// Package services provides business logic implementations for the application.
package services

import (
	"errors"
	"example/api/internal/models"
)

// UserService manages user-related operations such as registration, listing, finding, and deleting users.
// It maintains an in-memory collection of users and handles user ID generation.
type UserService struct {
	users  []models.User
	nextId int
}

// NewUserService creates and returns a new instance of UserService with initialized fields.
func NewUserService() *UserService {
	return &UserService{
		users:  make([]models.User, 0),
		nextId: 1,
	}
}

// Register creates a new user with the given name and email.
// Returns the new user's ID and an error if registration fails.
// Registration fails if name or email is empty, or if the email already exists.
func (service *UserService) Register(name string, email string) (int, error) {
	if name == "" || email == "" {
		return 0, errors.New("name and email are required")
	}

	for _, u := range service.users {
		if email == u.Email {
			return 0, errors.New("email already exists")
		}
	}

	user := models.User{
		ID:    service.nextId,
		Name:  name,
		Email: email,
	}

	service.users = append(service.users, user)
	service.nextId++
	return user.ID, nil
}

// List returns all registered users.
func (s *UserService) List() []models.User {
	return s.users
}

// FindByID searches for a user by their ID.
// Returns the user if found, or an error if no user exists with the given ID.
func (s *UserService) FindByID(id int) (models.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// Delete removes a user with the specified ID from the service.
// Returns true if the user was found and deleted, false otherwise.
func (s *UserService) Delete(id int) bool {
	for i, u := range s.users {
		if u.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return true
		}
	}
	return false
}
