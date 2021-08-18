package models

type Organizer struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	NameOfContact string `json:"name_of_contact"`
	Email         string `json:"email"`
}
