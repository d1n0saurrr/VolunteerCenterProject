package webHandler

import (
	"VolunteerCenter/pkg/service"
	"github.com/gin-gonic/gin"
)

type WebHandler struct {
	services *service.Service
}

func NewWebHandler(services *service.Service) *WebHandler {
	return &WebHandler{services: services}
}

func (h *WebHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static/", "./ui/static/")

	router.GET("/", h.mainPage)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUpPage)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signInPage)
		auth.POST("/sign-in", h.singIn)
	}

	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/info", h.userInfo)
		user.POST("/info", h.userChange)
	}

	events := router.Group("/events", h.userIdentity)
	{
		events.GET("/new", h.getNewEvents)
		events.POST("/register", h.registerToEvent)
		events.GET("/visited", h.getVisitedEvents)
	}

	admin := router.Group("/admin", h.userIdentity)
	{
		admin.GET("/", h.adminPage)

		vols := admin.Group("/vols")
		{
			vols.GET("/", h.getAllVols)
			vols.GET("/:id", h.getVolById)
		}

		events := admin.Group("/events")
		{
			events.GET("/", h.getEvents)
			events.POST("/add", h.addEvent)
		}
	}

	return router
}
