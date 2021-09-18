package handler

import (
	"github.com/gin-gonic/gin"
	moneytracker "github.com/len3fun/money-tracker"
	"net/http"
)

func (h *Handler) CreateCurrency(c *gin.Context) {
	var input moneytracker.Currency
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

	err = h.services.Currency.CreateCurrency(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) getAllCurrencies(c *gin.Context) {
	currencies, err := h.services.Currency.GetAllCurrencies()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, currencies)
}
