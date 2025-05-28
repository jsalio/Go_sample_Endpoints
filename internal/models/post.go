package models

// Post represents a post entity in the system.
// It contains basic post information such as ID, title, content, and user ID.
type Post struct {
	// ID is the unique identifier for the post
	ID int `json:"id"`
	// Title represents the post's title
	Title string `json:"title"`
	// Content represents the post's content
	Content string `json:"content"`
	// UserID is the ID of the user who created the post
	UserID int `json:"user_id"`
}
