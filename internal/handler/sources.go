package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllSources(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sources, err := h.services.MoneySource.GetAllSources(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sources)
}
