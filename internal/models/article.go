package models

import (
	"time"
)

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (post *Article) Validate() (string, bool) {

	if post.Title == "" {
		return "Title should be on the payload", false
	}

	if post.Content == "" {
		return "Content should be on the payload", false
	}

	return "Payload is valid", true
}
