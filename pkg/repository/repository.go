package repository

import (
	"betera-test/pkg/common"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Apod interface {
	GetImageByDate(ctx context.Context, date string) (ApodDB, error)
	GetImageByDatesRange(ctx context.Context, dates []string) ([]ApodDB, error)
	SaveImage(ctx context.Context, image []byte, url, date string) error
}

type Repository struct {
	Apod
	db *sqlx.DB
}

type ApodDB struct {
	Id  int    `db:"id"`
	Url string `db:"url"`
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetImageByDate(ctx context.Context, date string) (ApodDB, error) {
	var apod ApodDB
	query := "SELECT id, url FROM apod WHERE id = $1"
	if err := r.db.Get(&apod, query, common.DateStringToInt(date)); err != nil {
		return apod, err
	}

	return apod, nil
}
func (r *Repository) GetImageByDatesRange(ctx context.Context, dates []int) ([]ApodDB, error) {
	var apod []ApodDB

	query := "SELECT id, url FROM apod WHERE id = any($1)"
	if err := r.db.Select(&apod, query, pq.Array(dates)); err != nil {
		return nil, err
	}

	return apod, nil
}
func (r *Repository) SaveImage(ctx context.Context, image string, url, date string) (int, error) {
	var id int
	query := "INSERT INTO apod (id, url, img) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, common.DateStringToInt(date), url, image)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
