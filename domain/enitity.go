package domain

import (
	"fmt"

	uuid "github.com/google/uuid"
)

// Entity struct definition
type Entity struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewEntity initialize MyEntity
func NewEntity(title string) (*Entity, error) {
	id := uuid.New()
	if title == "" {
		return nil, fmt.Errorf("Invalid title")
	}

	return &Entity{
		ID:    id.String(),
		Title: title,
	}, nil
}
