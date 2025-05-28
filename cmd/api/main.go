package main

import (
	"example/api/internal/api/handlers"
	"example/api/internal/api/middleware"
	"example/api/internal/services"
	"fmt"
	"log"
	"net/http"

	_ "example/api/docs" // This will be generated

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User Management API
// @version 1.0
// @description A RESTful API for managing users and their posts
// @host localhost:8085
// @BasePath /
func main() {
	fmt.Println("Api restfull")

	userService := services.NewUserService()
	userHandler := handlers.NewUserhandler(userService)

	postService := services.NewPostService()
	postHandler := handlers.NewPostHandler(postService)

	// Create a new mux router
	mux := http.NewServeMux()

	// Swagger documentation endpoint
	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8059/swagger/doc.json"),
	))

	// User endpoints
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.List(w, r)
		case http.MethodPost:
			userHandler.Register(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[len("/users/"):] == "posts" {
			postHandler.FindByUserID(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			userHandler.FindByID(w, r)
		case http.MethodDelete:
			userHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Post endpoints
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			postHandler.List(w, r)
		case http.MethodPost:
			postHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			postHandler.FindByID(w, r)
		case http.MethodDelete:
			postHandler.Delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Apply CORS middleware to all routes
	handler := middleware.CORS(mux)

	log.Println("Server starting on :8059")
	log.Fatal(http.ListenAndServe(":8059", handler))
}
