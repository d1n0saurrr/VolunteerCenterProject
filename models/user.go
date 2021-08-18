package models

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
	IdVolunteer int    `json:"id_volunteer"`
}
