package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

type VolInfo struct {
	User models.User
	Vol  models.Volunteer
}

func (h *WebHandler) getAllVols(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/admin/vols.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := h.services.User.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vols, err := h.services.Volunteer.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	volInfo := make([]VolInfo, len(users))

	for i := 0; i < len(users); i++ {
		volInfo[i].User = users[i]

		for j := 0; j < len(vols); j++ {
			if users[i].IdVolunteer == vols[j].Id {
				volInfo[i].Vol = vols[j]
				break
			}
		}
	}

	if t.ExecuteTemplate(c.Writer, "vols", volInfo) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) changeVol(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = user.IsAdmin.Scan(!user.IsAdmin.Bool)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.Update(user)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/admin/vols")
}

func (h *WebHandler) deleteVol(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("userId"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	volId, err := strconv.Atoi(c.Query("volId"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.Delete(userId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Volunteer.Delete(volId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/admin/vols")
}
