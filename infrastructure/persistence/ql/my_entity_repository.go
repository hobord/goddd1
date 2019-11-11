package persistence

import (
	"context"
	"database/sql"

	_ "github.com/cznic/ql/driver"
	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository"
)

// myEntityRepository Implements repository.MyEntityRepository
type myEntityRepository struct {
	conn *sql.DB
}

// NewMyEntityRepository returns initialized MyEntityRepositoryImpl
func NewMyEntityRepository(conn *sql.DB) repository.MyEntityRepository {
	return &myEntityRepository{conn: conn}
}

func (r *myEntityRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func (r *myEntityRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *myEntityRepository) Get(ctx context.Context, id string) (*domain.MyEntity, error) {
	row, err := r.queryRow(ctx, "SELECT id, title FROM entity WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	u := &domain.MyEntity{}
	err = row.Scan(&u.ID, &u.Title)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *myEntityRepository) GetAll(ctx context.Context) ([]*domain.MyEntity, error) {
	rows, err := r.query(ctx, "SELECT id, title FROM entity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	us := make([]*domain.MyEntity, 0)
	for rows.Next() {
		u := &domain.MyEntity{}
		err = rows.Scan(&u.ID, &u.Title)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func (r *myEntityRepository) Save(ctx context.Context, entity *domain.MyEntity) error {
	stmt, err := r.conn.Prepare("INSERT INTO entity (id, title) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity.ID, entity.Title)
	return err
}

func (r *myEntityRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.conn.Prepare("DELETE FROM entity WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
