package service

import (
	"betera-test/pkg/client"
	"betera-test/pkg/domain"
	"betera-test/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Apod interface {
	GetImageByDate(ctx *gin.Context, date string) (domain.ApodResp, error)
	GetImageByRange(ctx *gin.Context, startDate, endDate string) ([]domain.ApodResp, error)
}

type Service struct {
	Apod
	ApodClient *client.Client
	Repos      *repository.Repository
}

func NewService(repos *repository.Repository, client *client.Client) *Service {
	return &Service{ApodClient: client, Repos: repos}
}

func (s *Service) GetImageByDate(ctx *gin.Context, date string) (*domain.ApodResp, error) {
	apod, err := s.ApodClient.GetImageByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	return apod, nil
}

func (s *Service) GetImageByRange(ctx *gin.Context, startDate, endDate string) (*[]domain.ApodResp, error) {
	apod, err := s.ApodClient.GetImageByRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return apod, nil
}
