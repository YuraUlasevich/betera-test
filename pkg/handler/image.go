package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetImage(c *gin.Context) {
	startDate := c.Params.ByName("start_date")
	endDate := c.Params.ByName("end_date")
	date := c.Params.ByName("date")

	if date == "" && endDate == "" && startDate != "" {
		date = time.Now().Format("2006-01-02")
	}

	if date != "" && endDate != "" && startDate == "" {
		date = endDate
	}

	if err := validateDate(startDate, endDate, date); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if date != "" {
		err := h.service.GetImageByDate(c, date)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	err := h.service.GetImageByRange(c, startDate, endDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
			return fmt.Errorf("End date %v is greater than now date %v", endDate, time.Now())
		}
	}

	return nil
}
