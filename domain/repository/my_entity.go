package repository

import (
	"context"

	"github.com/hobord/goddd1/domain"
)

// MyEntityRepository interface definition
type MyEntityRepository interface {
	// Get return entity by id
	Get(ctx context.Context, id string) (*domain.MyEntity, error)

	// GetAll return all entities
	GetAll(ctx context.Context) ([]*domain.MyEntity, error)

	// Save is save to persintent the entity
	Save(ctx context.Context, entity *domain.MyEntity) error

	// Delete entity from persitnet store
	Delete(ctx context.Context, id string) error
}
