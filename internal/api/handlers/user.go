// Package handlers provides HTTP handlers for the API endpoints.
package handlers

import (
	"encoding/json"
	"example/api/internal/services"
	"net/http"
	"strconv"
)

// UserHandler handles HTTP requests related to user operations.
// It contains a reference to the user service that implements the business logic.
type UserHandler struct {
	service *services.UserService
}

// NewUserhandler creates a new instance of UserHandler with the provided user service.
// It returns a pointer to the newly created UserHandler.
func NewUserhandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register handles POST /users endpoint.
// @Summary Create a new user
// @Description Create a new user with the provided name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body object true "User object"
// @Success 201 {object} map[string]int
// @Failure 400 {string} string
// @Router /users [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := h.service.Register(input.Name, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// List handles GET /users endpoint.
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users := h.service.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// FindByID handles GET /users/{id} endpoint.
// @Summary Get user by ID
// @Description Retrieve a specific user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users/{id} [get]
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete handles DELETE /users/{id} endpoint.
// @Summary Delete user
// @Description Delete a user by their ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if !h.service.Delete(id) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
