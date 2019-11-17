package usecase

import (
	"context"
	"testing"

	entities "github.com/hobord/goddd1/domain/entity"
	"github.com/hobord/goddd1/domain/repository/mocks"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {
	mockRepository := &mocks.FooEntityRepository{}

	fakeID := fake.Sentence()
	fakeTitle := fake.Sentence()
	returnMockEntity := &entities.FooEntity{
		ID:    fakeID,
		Title: fakeTitle,
	}
	mockRepository.On("GetByID", mock.Anything, mock.Anything).Return(returnMockEntity, nil)

	interactor := NewExampleInteractor(mockRepository)
	result, err := interactor.GetByID(context.TODO(), "1")
	if err != nil {
		assert.NoError(t, err)
	}
	assert.Equal(t, result.ID, fakeID, "The result ID should be:"+fakeID)
	assert.Equal(t, result.Title, fakeTitle, "The result ID should be:"+fakeTitle)
}

// TODO: implement all use cases tests
