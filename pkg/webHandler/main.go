package webHandler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const (
	footer = 0
)

func (h *WebHandler) mainPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/main.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "main", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) signOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "", true, true)
	h.header = "ui/html/authHeader.html"
	c.Redirect(http.StatusFound, "/")
}
