package webHandler

import (
	"VolunteerCenter/pkg/service"
	"github.com/gin-gonic/gin"
)

type WebHandler struct {
	services *service.Service
	header   string
}

func NewWebHandler(services *service.Service) *WebHandler {
	return &WebHandler{
		services: services,
		header:   "ui/html/authHeader.html",
	}
}

func (h *WebHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static/", "./ui/static/")

	router.GET("/", h.mainPage)
	router.GET("/sign-out", h.signOut)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUpPage)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signInPage)
		auth.POST("/sign-in", h.singIn)
	}

	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/info", h.userInfoPage)
		user.POST("/info", h.changeUserInfo)
	}

	events := router.Group("/events", h.userIdentity)
	{
		events.GET("/new", h.getNewEvents)
		events.POST("/register", h.registerToEvent)
		events.GET("/visited", h.getVisitedEvents)
	}

	admin := router.Group("/admin", h.userIdentity, h.adminIdentity)
	{
		admin.GET("/", h.adminPage)

		vols := admin.Group("/vols")
		{
			vols.GET("/", h.getAllVols)
			vols.POST("/change", h.changeVol)
			vols.POST("/delete", h.deleteVol)
		}

		events := admin.Group("/events")
		{
			events.GET("/", h.getEvents)
			events.POST("/add", h.addEvent)
			events.POST("/delete", h.deleteEvent)
			events.GET("/get_by_vol", h.getVolEvents)
		}
	}

	return router
}
