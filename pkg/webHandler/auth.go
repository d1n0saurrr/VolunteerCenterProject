package webHandler

import (
	"VolunteerCenter/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func (h *WebHandler) signUpPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/auth/signup.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "signup", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *WebHandler) signUp(c *gin.Context) {
	input := models.User{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	if len(input.Username) == 0 || len(input.Password) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Required fields are empty!")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	t, err := time.Parse("2006-01-02", c.PostForm("birth_date"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vol := models.Volunteer{
		FirstName:  c.PostForm("first_name"),
		SecondName: c.PostForm("second_name"),
		Patronymic: c.PostForm("patronymic"),
		BirthDate:  t,
	}

	volId, err := h.services.Volunteer.Create(vol)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.SetVolId(id, volId)

	c.Redirect(http.StatusFound, "/auth/sign-in")
}

func (h *WebHandler) signInPage(c *gin.Context) {
	t, err := template.ParseFiles("ui/html/auth/signin.html", h.header)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if t.ExecuteTemplate(c.Writer, "signin", nil) != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
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

	user, err := h.services.User.GetByUsername(c.PostForm("username"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.IsAdmin.Valid {
		if user.IsAdmin.Bool {
			h.header = "ui/html/adminHeader.html"
		}
	} else {
		h.header = "ui/html/header.html"
	}

	c.SetCookie("token", token, 60*60*2, "/", "", true, true)
	c.Redirect(http.StatusFound, "/")
}
