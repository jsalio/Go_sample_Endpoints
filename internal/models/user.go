package models

// User represents a user entity in the system.
// It contains basic user information such as ID, name, and email.
type User struct {
	// ID is the unique identifier for the user
	ID int `json:"id"`
	// Name represents the user's full name
	Name string `json:"name"`
	// Email is the user's email address
	Email string `json:"email"`
}
