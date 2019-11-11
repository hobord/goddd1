package persistence

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/domain/repository"
)

// EntityRepository Implements repository.EntityRepository
type entityRepository struct {
	conn *sql.DB
}

// NewEntityMysqlRepository returns initialized EntityRepositoryImpl
func NewEntityMysqlRepository(conn *sql.DB) repository.EntityRepository {
	return &entityRepository{conn: conn}
}

func (r *entityRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func (r *entityRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *entityRepository) Get(ctx context.Context, id string) (*domain.Entity, error) {
	row, err := r.queryRow(ctx, "SELECT id, title FROM entity WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	u := &domain.Entity{}
	err = row.Scan(&u.ID, &u.Title)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *entityRepository) GetAll(ctx context.Context) ([]*domain.Entity, error) {
	rows, err := r.query(ctx, "SELECT id, title FROM entity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	us := make([]*domain.Entity, 0)
	for rows.Next() {
		u := &domain.Entity{}
		err = rows.Scan(&u.ID, &u.Title)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func (r *entityRepository) Save(ctx context.Context, entity *domain.Entity) error {
	stmt, err := r.conn.Prepare("INSERT INTO entity (id, title) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, entity.ID, entity.Title)
	return err
}

func (r *entityRepository) Delete(ctx context.Context, id string) error {
	stmt, err := r.conn.Prepare("DELETE FROM entity WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
