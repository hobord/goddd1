package persistence

import (
	"context"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/cznic/ql/driver"
	"github.com/stretchr/testify/assert"
)

func TestGetEntity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow("TEST_ID", "my title")

	query := "select id, title from entity where id=?"

	prep := mock.ExpectPrepare(query)
	testID := "TEST_ID"
	prep.ExpectQuery().WithArgs(testID).WillReturnRows(rows)

	repository := NewMyEntityRepository(db)

	entity, err := repository.Get(context.TODO(), testID)

	if err != nil {
		assert.NoError(t, err)
	}
	assert.NotNil(t, entity)
	assert.Equal(t, entity.ID, testID, "The IDs should be the same")
}

// TODO: all other methods test implementation
