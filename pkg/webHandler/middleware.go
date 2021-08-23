package webHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
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

	userId, err := h.services.Authorization.ParseToken(token)

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}
