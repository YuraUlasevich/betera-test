package repository

type Apod interface{}

type Repository struct {
	Apod
}

func NewRepository() *Repository {
	return &Repository{}
}
