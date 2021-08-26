package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type UserInfo struct {
	Username   string
	FirstName  string
	SecondName string
	Patronymic string
	BirthDate  string
}

func (h *WebHandler) userInfoPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/user/info.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetById(id.(int))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vol, err := h.services.Volunteer.GetById(user.IdVolunteer)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "info", UserInfo{
		user.Username,
		vol.FirstName,
		vol.SecondName,
		vol.Patronymic,
		vol.BirthDate.Format("2006-01-02"),
	}) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) changeUserInfo(c *gin.Context) {
	t, err := time.Parse("2006-01-02", c.PostForm("birth_date"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetById(id.(int))

	vol := models.Volunteer{
		Id:         user.IdVolunteer,
		FirstName:  c.PostForm("first_name"),
		SecondName: c.PostForm("second_name"),
		Patronymic: c.PostForm("patronymic"),
		BirthDate:  t,
	}

	err = h.services.Volunteer.Update(vol)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/user/info")
}

func (h *WebHandler) getNewEvents(c *gin.Context) {

}

func (h *WebHandler) registerToEvent(c *gin.Context) {

}

func (h *WebHandler) getVisitedEvents(c *gin.Context) {

}
