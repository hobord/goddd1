package repository

import (
	"context"

	"github.com/hobord/goddd1/domain"
)

// MyEntityRepository interface definition
type MyEntityRepository interface {
	Get(ctx context.Context, id string) (*domain.MyEntity, error)
	GetAll(ctx context.Context) ([]*domain.MyEntity, error)
	Save(ctx context.Context, entity *domain.MyEntity) error
	Delete(ctx context.Context, id string) error
}
