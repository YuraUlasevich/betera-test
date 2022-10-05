package client

import (
	"betera-test/pkg/domain"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApodClient interface {
	GetImageByDate(ctx *gin.Context, date string) (*domain.ApodResp, error)
	GetImageByRange(ctx *gin.Context, startDate, endDate string) (*[]domain.ApodResp, error)
	GetImage(ctx *gin.Context, url string) (string, error)
}

type Client struct {
	ApodClient
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) GetImageByDate(ctx *gin.Context, date string) (*domain.ApodResp, error) {
	url := fmt.Sprintf("%s&date=%s", c.url, date)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	apodResp := &domain.ApodResp{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(apodResp)
	if err != nil {
		return nil, err
	}
	return apodResp, nil
}
func (c *Client) GetImageByRange(ctx *gin.Context, startDate, endDate string) (*[]domain.ApodResp, error) {
	url := fmt.Sprintf("%s&start_date=%s&end_date=%s", c.url, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	apodResp := &[]domain.ApodResp{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(apodResp)
	if err != nil {
		return nil, err
	}
	return apodResp, nil
}

func (c *Client) GetImage(ctx *gin.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(bytes)

	return encoded, nil
}
