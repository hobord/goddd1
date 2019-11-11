package usecase

import (
	"context"

	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository"
)

// ExampleInteractor provides use-case implementation
type ExampleInteractor struct {
	Repository repository.MyEntityRepository
}

// NewExampleInteractor is create a new example "service" / "interactor"
func NewExampleInteractor(Repository *repository.MyEntityRepository) *ExampleInteractor {
	return &ExampleInteractor{
		Repository: Repository,
	}
}

// Get return entity by id
func (i *ExampleInteractor) Get(ctx context.Context, id string) (*domain.MyEntity, error) {
	return i.Repository.Get(ctx, id)
}

// GetAll return all entities
func (i *ExampleInteractor) GetAll(ctx context.Context) ([]*domain.MyEntity, error) {
	return i.Repository.GetAll(ctx)
}

// Save is save to persintent the entity
func (i *ExampleInteractor) Save(ctx context.Context, entity *domain.MyEntity) error {
	return i.Repository.Save(ctx, entity)
}

// Delete entity from persitnet store
func (i *ExampleInteractor) Delete(ctx context.Context, id string) error {
	return i.Repository.Delete(ctx, id)
}
