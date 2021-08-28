package webHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userAdmin           = "isAdmin"
)

func (h *WebHandler) userIdentity(c *gin.Context) {
	token, err := c.Cookie("token")

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no token in cookie")
	}

	if token == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty token cookie")
		return
	}

	userId, isAdmin, err := h.services.Authorization.ParseToken(token)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
	c.Set(userAdmin, isAdmin)
	//h.header = "ui/html/header.html"
}

func (h *WebHandler) adminIdentity(c *gin.Context) {
	isAdmin, ok := c.Get(userAdmin)

	if !ok || !isAdmin.(bool) {
		newErrorResponse(c, http.StatusBadGateway, "not admin")
	}
	//h.header = "ui/html/adminHeader.html"
}
