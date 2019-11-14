package repository

import (
	"context"

	"github.com/hobord/goddd1/domain"
)

// EntityRepository interface definition
type EntityRepository interface {
	// Get return entity by id
	GetByID(ctx context.Context, id string) (*domain.Entity, error)

	// GetAll return all entities
	GetAll(ctx context.Context) ([]*domain.Entity, error)

	// Save is save to persintent the entity
	Save(ctx context.Context, entity *domain.Entity) error

	// Delete entity from persitnet store
	Delete(ctx context.Context, id string) error
}
