package service

import (
	"betera-test/pkg/client"
	"betera-test/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Apod interface {
	GetImageByDate(ctx *gin.Context, date string) error
	GetImageByRange(ctx *gin.Context, startDate, endDate string) error
}

type Service struct {
	Apod
	ApodClient *client.Client
	Repos      *repository.Repository
}

func NewService(repos *repository.Repository, client *client.Client) *Service {
	return &Service{ApodClient: client, Repos: repos}
}

func (s *Service) GetImageByDate(ctx *gin.Context, date string) error {
	return nil
}

func (s *Service) GetImageByRange(ctx *gin.Context, startDate, endDate string) error {
	return nil
}
