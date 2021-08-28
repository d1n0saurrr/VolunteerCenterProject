package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (h *WebHandler) adminPage(c *gin.Context) {

}

func (h *WebHandler) getEvents(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/admin/events.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := h.services.Event.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	for i := 0; i < len(events); i++ {
		events[i].Start = events[i].StartDate.Format("02.01.2006")
		events[i].End = events[i].EndDate.Format("02.01.2006")
	}

	if t.ExecuteTemplate(c.Writer, "events", events) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) addEvent(c *gin.Context) {
	start, err := time.Parse("2006-01-02", c.PostForm("start_date"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	end, err := time.Parse("2006-01-02", c.PostForm("end_date"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	event := models.Event{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Location:    c.PostForm("location"),
		StartDate:   start,
		EndDate:     end,
	}

	_, err = h.services.Event.Create(event)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/admin/events")
}

func (h *WebHandler) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Event.Delete(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/admin/events")
}

func (h *WebHandler) getVolEvents(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	events, err := h.services.Event.GetVolEvents(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	for i := 0; i < len(events); i++ {
		events[i].Start = events[i].StartDate.Format("02.01.2006")
		events[i].End = events[i].EndDate.Format("02.01.2006")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"events": events,
	})
}
