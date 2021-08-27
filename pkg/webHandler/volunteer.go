package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
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

}

func (h *WebHandler) getVolById(c *gin.Context) {

}
