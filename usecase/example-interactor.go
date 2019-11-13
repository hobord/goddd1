package usecase

import (
	"context"

	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository"
)

// ExampleInteractor provides use-case implementation
type ExampleInteractor struct {
	Repository repository.EntityRepository
}

// NewExampleInteractor is create a new example "service" / "interactor"
func NewExampleInteractor(repository *repository.EntityRepository) *ExampleInteractor {
	return &ExampleInteractor{
		Repository: *repository,
	}
}

// Get return entity by id
func (i *ExampleInteractor) Get(ctx context.Context, id string) (*domain.Entity, error) {
	return i.Repository.Get(ctx, id)
}

// GetAll return all entities
func (i *ExampleInteractor) GetAll(ctx context.Context) ([]*domain.Entity, error) {
	return i.Repository.GetAll(ctx)
}

// Save is save to persistent the entity
func (i *ExampleInteractor) Save(ctx context.Context, entity *domain.Entity) error {
	return i.Repository.Save(ctx, entity)
}

// Delete entity from persistent store
func (i *ExampleInteractor) Delete(ctx context.Context, id string) error {
	return i.Repository.Delete(ctx, id)
}
