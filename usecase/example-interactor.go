package usecase

import (
	"context"

	"github.com/hobord/goddd1/domain/entity"
	"github.com/hobord/goddd1/domain/repository"
)

// ExampleInteractorInterface is the interface for example use case
// mockery -name=ExampleInteractorInterface
type ExampleInteractorInterface interface {
	GetByID(ctx context.Context, id string) (*entities.FooEntity, error)
	GetAll(ctx context.Context) ([]*entities.FooEntity, error)
	Save(ctx context.Context, entity *entities.FooEntity) error
	Delete(ctx context.Context, id string) error
}

// ExampleInteractor provides an example use-case implementation
type ExampleInteractor struct {
	FooEntityRepository repository.FooEntityRepository
	// ...Other repositories or interactors
}

// NewExampleInteractor is create a new example "service" / "interactor"
func NewExampleInteractor(repository repository.FooEntityRepository) *ExampleInteractor {
	return &ExampleInteractor{
		FooEntityRepository: repository,
	}
}

// GetByID return entity by id
func (i *ExampleInteractor) GetByID(ctx context.Context, id string) (*entities.FooEntity, error) {
	return i.FooEntityRepository.GetByID(ctx, id)
}

// GetAll return all entities
func (i *ExampleInteractor) GetAll(ctx context.Context) ([]*entities.FooEntity, error) {
	return i.FooEntityRepository.GetAll(ctx)
}

// Save is save to persistent the entity
func (i *ExampleInteractor) Save(ctx context.Context, entity *entities.FooEntity) error {
	return i.FooEntityRepository.Save(ctx, entity)
}

// Delete entity from persistent store
func (i *ExampleInteractor) Delete(ctx context.Context, id string) error {
	return i.FooEntityRepository.Delete(ctx, id)
}
