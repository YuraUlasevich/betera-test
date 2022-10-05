package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Apod interface {
	GetImageByDate(ctx context.Context, date string)
	GetImageByDatesRange(ctx context.Context, startDate, endDate string)
	SaveImage(ctx context.Context, image []byte, url, date string) error
}

type Repository struct {
	Apod
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetImageByDate(ctx context.Context, date string) {

}
func (r *Repository) GetImageByDatesRange(ctx context.Context, startDate, endDate string) {

}
func (r *Repository) SaveImage(ctx context.Context, image string, url, date string) (int, error) {
	var id int
	query := "INSERT INTO apod (id, url, img) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, dateStringToInt(date), url, image)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func dateStringToInt(date string) int {
	split := strings.Split(date, "-")
	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])
	result := year*10000 + month*100 + day
	return result
}
