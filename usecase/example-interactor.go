package usecase

import (
	"context"

	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository"
)

// ExampleInteractorInterface is the interface for example use case
// mockery -name=ExampleInteractorInterface
type ExampleInteractorInterface interface {
	GetByID(ctx context.Context, id string) (*domain.Entity, error)
	GetAll(ctx context.Context) ([]*domain.Entity, error)
	Save(ctx context.Context, entity *domain.Entity) error
	Delete(ctx context.Context, id string) error
}

// ExampleInteractor provides an example use-case implementation
type ExampleInteractor struct {
	EntityRepository repository.EntityRepository
	// ...Other repositories or interactors
}

// NewExampleInteractor is create a new example "service" / "interactor"
func NewExampleInteractor(repository repository.EntityRepository) *ExampleInteractor {
	return &ExampleInteractor{
		EntityRepository: repository,
	}
}

// GetByID return entity by id
func (i *ExampleInteractor) GetByID(ctx context.Context, id string) (*domain.Entity, error) {
	return i.EntityRepository.GetByID(ctx, id)
}

// GetAll return all entities
func (i *ExampleInteractor) GetAll(ctx context.Context) ([]*domain.Entity, error) {
	return i.EntityRepository.GetAll(ctx)
}

// Save is save to persistent the entity
func (i *ExampleInteractor) Save(ctx context.Context, entity *domain.Entity) error {
	return i.EntityRepository.Save(ctx, entity)
}

// Delete entity from persistent store
func (i *ExampleInteractor) Delete(ctx context.Context, id string) error {
	return i.EntityRepository.Delete(ctx, id)
}
