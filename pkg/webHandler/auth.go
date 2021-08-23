package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func (h *WebHandler) signUpPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/signup.html", header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "signup", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) signInPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/signin.html", header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "signin", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) signUp(c *gin.Context) {
	input := models.User{Name: c.PostForm("name"), Username: c.PostForm("username"), Password: c.PostForm("password")}

	if len(input.Name) == 0 || len(input.Username) == 0 || len(input.Password) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Required fields are empty!")
		return
	}

	_, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusFound, "/auth/sign-in")
}

func (h *WebHandler) singIn(c *gin.Context) {
	if len(c.PostForm("username")) == 0 || len(c.PostForm("password")) == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "Required fields are empty!")
		return
	}

	token, err := h.services.Authorization.GenerateToken(c.PostForm("username"), c.PostForm("password"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	/*c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})*/

	c.SetCookie("token", token, 60*60*2, "/", "", true, true)
	c.Redirect(http.StatusFound, "/user/info")
}
