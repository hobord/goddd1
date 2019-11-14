package usecase

import (
	"context"
	"testing"

	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository/mocks"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	mockRepository := &mocks.EntityRepository{}

	returnMockEntity := &domain.Entity{
		ID:    "1",
		Title: "Works",
	}
	mockRepository.On("GetByID", mock.Anything, mock.Anything).Return(returnMockEntity, nil)

	interactor := NewExampleInteractor(mockRepository)
	result, err := interactor.GetByID(context.TODO(), "1")
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, result.ID, "1", "The result ID should be 1")
	assert.Equal(t, result.Title, "Works", "The result ID should be 1")
}
