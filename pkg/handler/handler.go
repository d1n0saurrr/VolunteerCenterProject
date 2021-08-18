package handler

import (
	"VolunteerCenter/pkg/service"
	"github.com/gin-gonic/gin"
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
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api")
	{
		vols := api.Group("/vols")
		{
			vols.GET("/", h.getAllVols)
			vols.GET("/:id", h.getVolById)
		}
	}

	return router
}
