package handlers

import (
	"encoding/json"
	"example/api/internal/services"
	"net/http"
	"strconv"
)

// PostHandler handles HTTP requests related to post operations.
// It contains a reference to the post service that implements the business logic.
type PostHandler struct {
	service *services.PostService
}

// NewPostHandler creates a new instance of PostHandler with the provided post service.
// It returns a pointer to the newly created PostHandler.
func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

// Create handles POST /posts endpoint.
// @Summary Create a new post
// @Description Create a new post with the provided title, content, and user ID
// @Tags posts
// @Accept json
// @Produce json
// @Param post body object true "Post object"
// @Success 201 {object} map[string]int
// @Failure 400 {string} string
// @Router /posts [post]
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  int    `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id, err := h.service.Create(input.Title, input.Content, input.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// List handles GET /posts endpoint.
// @Summary Get all posts
// @Description Retrieve a list of all posts
// @Tags posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts := h.service.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// FindByID handles GET /posts/{id} endpoint.
// @Summary Get post by ID
// @Description Retrieve a specific post by its ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /posts/{id} [get]
func (h *PostHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// FindByUserID handles GET /users/{id}/posts endpoint.
// @Summary Get posts by user ID
// @Description Retrieve all posts for a specific user
// @Tags posts
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Post
// @Failure 400 {string} string
// @Router /users/{id}/posts [get]
func (h *PostHandler) FindByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	idStr = idStr[:len(idStr)-len("/posts")]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	posts := h.service.FindByUserID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Delete handles DELETE /posts/{id} endpoint.
// @Summary Delete post
// @Description Delete a post by its ID
// @Tags posts
// @Param id path int true "Post ID"
// @Success 204 "No Content"
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /posts/{id} [delete]
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	if !h.service.Delete(id) {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
