package persistence

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hobord/goddd1/domain/entity"
	"github.com/hobord/goddd1/domain/repository"
)

// FooEntityRepository Implements repository.FooEntityRepository
type FooEntityRepository struct {
	conn *sql.DB
}

// NewEntityMysqlRepository returns initialized FooEntityRepositoryImpl
func NewEntityMysqlRepository(conn *sql.DB) repository.FooEntityRepository {
	return &FooEntityRepository{conn: conn}
}

func (r *FooEntityRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func (r *FooEntityRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *FooEntityRepository) GetByID(ctx context.Context, id string) (*entities.FooEntity, error) {
	row, err := r.queryRow(ctx, "SELECT id, title FROM entity WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	entity := &entities.FooEntity{}
	err = row.Scan(&entity.ID, &entity.Title)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *FooEntityRepository) GetAll(ctx context.Context) ([]*entities.FooEntity, error) {
	rows, err := r.query(ctx, "SELECT id, title FROM entity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*entities.FooEntity, 0)

	for rows.Next() {
		entity := &entities.FooEntity{}
		err = rows.Scan(&entity.ID, &entity.Title)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}

	return results, nil
}

func (r *FooEntityRepository) Save(ctx context.Context, entity *entities.FooEntity) error {
	stmt, err := r.conn.Prepare("INSERT INTO entity (id, title) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity.ID, entity.Title)
	return err
}

func (r *FooEntityRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.conn.Prepare("DELETE FROM entity WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
