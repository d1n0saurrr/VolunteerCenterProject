package models

import "time"

type Volunteer struct {
	Id         int       `json:"-" db:"id"`
	FirstName  string    `json:"first_name" db:"first_name"`
	SecondName string    `json:"second_name" db:"second_name"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
	BirthDate  time.Time `json:"birth_date" db:"birth_date"`
}
