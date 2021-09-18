package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/len3fun/money-tracker/internal/models"
	"net/http"
	"time"
)

func (h *Handler) getAllActivities(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	activities, err := h.services.Activity.GetAllActivities(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (h *Handler) createActivity(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Activity
	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.ActivityDate.IsZero() {
		input.ActivityDate = time.Now()
	}

	input.UserId = userId

	err = h.services.Activity.CreateActivity(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Ok")
}
