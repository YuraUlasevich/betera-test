package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Apod interface {
	GetImageByDate(ctx context.Context, date string)
	GetImageByDatesRange(ctx context.Context, startDate, endDate string)
	SaveImage(ctx context.Context, image []byte, url, date string)
}

type Repository struct {
	Apod
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

func (r *Repository) GetImageByDate(ctx context.Context, date string) {

}
func (r *Repository) GetImageByDatesRange(ctx context.Context, startDate, endDate string) {

}
func (r *Repository) SaveImage(ctx context.Context, image []byte, url, date string) {

}
