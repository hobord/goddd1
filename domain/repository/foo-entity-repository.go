package repository

import (
	"context"

	entities "github.com/hobord/goddd1/domain/entity"
)

// FooEntityRepository interface definition
// mockery -name=FooEntityRepository
type FooEntityRepository interface {
	// Get return entity by id
	GetByID(ctx context.Context, id string) (*entities.FooEntity, error)

	// GetAll return all FooEntities
	GetAll(ctx context.Context) ([]*entities.FooEntity, error)

	// Save is save to persistent the FooEntity
	Save(ctx context.Context, entity *entities.FooEntity) error

	// Delete FooEntity from persistent store
	Delete(ctx context.Context, id string) error
}
