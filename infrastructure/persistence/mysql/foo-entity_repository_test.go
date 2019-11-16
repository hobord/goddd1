package persistence

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/hobord/goddd1/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestGetEntityByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow("TEST_ID", "my title")

	query := "SELECT id, title FROM entity WHERE id=?"

	prep := mock.ExpectPrepare(query)
	testID := "TEST_ID"
	prep.ExpectQuery().WithArgs(testID).WillReturnRows(rows)

	repository := NewEntityMysqlRepository(db)

	entity, err := repository.GetByID(context.TODO(), testID)

	if err != nil {
		assert.NoError(t, err)
	}
	assert.NotNil(t, entity)
	assert.Equal(t, entity.ID, testID, "The IDs should be the same")
}

func TestCreateEntity(t *testing.T) {
	entity, err := entities.NewFooEntity("test")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectPrepare("INSERT INTO entity").
		ExpectExec().
		WithArgs(entity.ID, entity.Title).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewEntityMysqlRepository(db)

	err = repository.Save(context.TODO(), entity)

	if err != nil {
		assert.NoError(t, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllEntity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow("TEST_ID", "my title1").
		AddRow("TEST_ID2", "my title2").
		AddRow("TEST_ID3", "my title3")

	query := "SELECT id, title FROM entity"

	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WillReturnRows(rows)

	repository := NewEntityMysqlRepository(db)

	entities, err := repository.GetAll(context.TODO())

	if err != nil {
		assert.NoError(t, err)
	}
	assert.NotNil(t, entities)
	assert.Equal(t, len(entities), 3, "The count of results row should be 3")
	assert.Equal(t, entities[0].ID, "TEST_ID", "The first elements id should be \"TEST_ID\" ")
	assert.Equal(t, entities[1].ID, "TEST_ID2", "The second elements id should be \"TEST_ID\" ")
	assert.Equal(t, entities[2].ID, "TEST_ID3", "The third elements id should be \"TEST_ID\" ")
	assert.Equal(t, entities[0].Title, "my title1", "The first elements id should be \"my title1\" ")
	assert.Equal(t, entities[1].Title, "my title2", "The second elements id should be \"my title2\" ")
	assert.Equal(t, entities[2].Title, "my title3", "The third elements id should be \"my title3\" ")
}

// TODO: all other methods test implementation
