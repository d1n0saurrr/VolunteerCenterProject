package webHandler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const (
	header = "ui/html/header.html"
	footer = 0
)

func (h *WebHandler) mainPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/main.html", header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "main", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
