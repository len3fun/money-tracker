package handler

import (
	"github.com/gin-gonic/gin"
	moneytracker "github.com/len3fun/money-tracker"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input moneytracker.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
