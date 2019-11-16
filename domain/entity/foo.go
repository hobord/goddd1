package entities

import (
	"fmt"

	uuid "github.com/google/uuid"
)

// FooEntity struct definition
type FooEntity struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewFooEntity initialize MyEntity
func NewFooEntity(title string) (*FooEntity, error) {
	id := uuid.New()
	if title == "" {
		return nil, fmt.Errorf("Invalid title")
	}

	return &FooEntity{
		ID:    id.String(),
		Title: title,
	}, nil
}
