package service

import (
	"betera-test/pkg/client"
	"betera-test/pkg/common"
	"betera-test/pkg/domain"
	"betera-test/pkg/repository"
	"fmt"
	"time"

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
	dbres, err := s.Repos.GetImageByDate(ctx, date)
	if err == nil {
		apod := &domain.ApodResp{
			Url:  dbres.Url,
			Date: common.DateIntToString(dbres.Id),
		}
		logrus.Info("Get image from db")
		return apod, nil
	}

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
	startDateInt := common.DateStringToInt(startDate)
	endDateInt := common.DateStringToInt(endDate)
	var dateArr []int
	for i := startDateInt; i <= endDateInt; i++ {
		_, err := time.Parse("2006-01-02", common.DateIntToString(i))
		if err == nil {
			dateArr = append(dateArr, i)
		}
	}
	logrus.Infof("Array of days is %v", dateArr)
	dbres, err := s.Repos.GetImageByDatesRange(ctx, dateArr)
	if err != nil {
		return nil, err
	}

	if len(dbres) != 0 {
		apods := []domain.ApodResp{}
		for _, v := range dbres {
			apod := domain.ApodResp{
				Url:  v.Url,
				Date: common.DateIntToString(v.Id),
			}
			apods = append(apods, apod)
		}

		datesMap := make(map[string]struct{})
		for _, v := range apods {
			datesMap[v.Date] = struct{}{}
		}

		notFoundDates := make([]string, 0)
		for _, v := range dateArr {
			if _, ok := datesMap[common.DateIntToString(v)]; !ok {
				notFoundDates = append(notFoundDates, common.DateIntToString(v))
			}
		}

		if len(notFoundDates) == 0 {
			return &apods, nil
		}

		for _, v := range notFoundDates {
			apod, err := s.GetImageByDate(ctx, v)
			if err != nil {
				return nil, err
			}
			apods = append(apods, *apod)
		}

		return &apods, nil
	}

	apod, err := s.ApodClient.GetImageByRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return apod, nil
}
