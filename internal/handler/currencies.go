package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/len3fun/money-tracker/internal/models"
	"net/http"
)

func (h *Handler) createCurrency(c *gin.Context) {
	var input models.Currency
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = input.Validate()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Currency.CreateCurrency(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllCurrencies(c *gin.Context) {
	currencies, err := h.services.Currency.GetAllCurrencies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, currencies)
}
