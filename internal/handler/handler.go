package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/len3fun/money-tracker/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.logRequest)
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity, h.logRequest)
	{
		sources := api.Group("/sources")
		{
			sources.GET("/", h.getAllSources)
			sources.POST("/", h.createSource)
		}
		currencies := api.Group("/currencies")
		{
			currencies.GET("/", h.getAllCurrencies)
			currencies.POST("/", h.createCurrency)
		}
		activities := api.Group("/activities")
		{
			activities.GET("/", h.getAllActivities)
			activities.POST("/", h.createActivity)
		}
	}

	return router
}
