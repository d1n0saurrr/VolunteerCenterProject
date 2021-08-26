package models

import "database/sql"

type User struct {
	Id          int          `json:"-" db:"id"`
	Username    string       `json:"username" binding:"required"`
	Password    string       `json:"password" binding:"required"`
	IsAdmin     sql.NullBool `json:"is_admin" db:"is_admin"`
	IdVolunteer int          `json:"id_volunteer" db:"id_volunteer"`
}
