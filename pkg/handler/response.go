package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/len3fun/money-tracker/pkg/logger"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
