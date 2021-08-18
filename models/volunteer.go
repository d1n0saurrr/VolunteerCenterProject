package models

import "time"

type Rank int

const (
	Ordinary Rank = iota
	Silver
	Gold
)

type Volunteer struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Patronymic string    `json:"patronymic"`
	BirthDate  time.Time `json:"birth_date"`
	Rank       Rank      `json:"rank"`
}
