package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/len3fun/money-tracker/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		sources := api.Group("/sources")
		{
			sources.GET("/", h.GetAllSources)
		}
		currencies := api.Group("/currencies")
		{
			currencies.GET("/", h.getAllCurrencies)
			currencies.POST("/", h.CreateCurrency)
		}
	}

	return router
}
