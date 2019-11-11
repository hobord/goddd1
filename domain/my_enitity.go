package domain

import (
	"fmt"

	uuid "github.com/google/uuid"
)

// MyEntity struct definition
type MyEntity struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewMyEntity initialize MyEntity
func NewMyEntity(title string) (*MyEntity, error) {
	id := uuid.New()
	if title == "" {
		return nil, fmt.Errorf("Invalid title")
	}

	return &MyEntity{
		ID:    id.String(),
		Title: title,
	}, nil
}
