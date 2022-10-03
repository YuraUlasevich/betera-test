package service

import "betera-test/pkg/repository"

type Apod interface{}

type Service struct {
	Apod
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
