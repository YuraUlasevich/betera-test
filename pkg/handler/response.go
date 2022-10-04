package handler

import "github.com/gin-gonic/gin"

type ApodResp struct {
	ApodResponses []ApodResponse
}

type ApodResponse struct {
	Date string `json:"date"`
	Url  string `json:"url"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
