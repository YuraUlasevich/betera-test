package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetImage(c *gin.Context) {
	startDate := c.Query("start")
	endDate := c.Query("end")
	date := c.Query("date")

	if date == "" && endDate == "" && startDate == "" {
		date = time.Now().Format("2006-01-02")
	}

	if date == "" && endDate == "" && startDate != "" {
		endDate = time.Now().Format("2006-01-02")
	}

	if date == "" && endDate != "" && startDate == "" {
		date = endDate
	}

	if err := validateDate(startDate, endDate, date); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if startDate == endDate && date == "" {
		date = startDate
	}

	if date != "" {
		apod, err := h.service.GetImageByDate(c, date)
		if err != nil {
			logrus.Fatal(err)
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, &apod)
		return
	}

	apod, err := h.service.GetImageByRange(c, startDate, endDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &apod)
}

func validateDate(startDate, endDate, date string) error {
	if date != "" && startDate != "" && endDate != "" {
		return errors.New("Too many params")
	}

	if startDate != "" {

		startDateTime, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return err
		}

		endDateTime, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return err
		}

		if endDateTime.Before(startDateTime) {
			return fmt.Errorf("Start date %v is greater than end date %v", startDate, endDate)
		}
		if endDateTime.After(time.Now()) {
			return fmt.Errorf("Date must be between 1995-06-16 and %v.", time.Now())
		}
	}

	if date != "" {
		dateTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return err
		}

		if dateTime.After(time.Now()) {
			return fmt.Errorf("Date must be between 1995-06-16 and %v.", time.Now())
		}
	}

	return nil
}
