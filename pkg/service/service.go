package service

import (
	"betera-test/pkg/client"
	"betera-test/pkg/domain"
	"betera-test/pkg/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal(err)
		return nil, err
	}

	logrus.Info(&apod)

	img, err := s.ApodClient.GetImage(ctx, apod.Url)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	logrus.Info("Get image done")

	id, err := s.Repos.SaveImage(ctx, img, apod.Url, apod.Date)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	fmt.Println(id)
	return apod, nil
}

func (s *Service) GetImageByRange(ctx *gin.Context, startDate, endDate string) (*[]domain.ApodResp, error) {
	apod, err := s.ApodClient.GetImageByRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return apod, nil
}
