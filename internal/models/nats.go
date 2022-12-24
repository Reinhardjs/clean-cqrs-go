package models

import "time"

type Message interface {
	Key() string
}

type ArticleCreatedMessage struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ArticleCreatedMessage) Key() string {
	return "article.created"
}
