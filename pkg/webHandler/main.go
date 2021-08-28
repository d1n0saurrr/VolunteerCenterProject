package webHandler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func (h *WebHandler) mainPage(c *gin.Context) {
	isAdmin, _ := c.Get(userAdmin)

	if isAdmin != nil && isAdmin.(bool) {
		h.header = "ui/html/adminHeader.html"
	} else if isAdmin != nil && !isAdmin.(bool) {
		h.header = "ui/html/header.html"
	}

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
