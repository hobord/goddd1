package domain

import "fmt"

// MyEntity struct definition
type MyEntity struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// NewMyEntity initialize MyEntity
func NewMyEntity(title string) (*MyEntity, error) {
	if title == "" {
		return nil, fmt.Errorf("Invalid title")
	}

	return &MyEntity{
		Title: title,
	}, nil
}
