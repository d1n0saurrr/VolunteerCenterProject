package repository

import (
	"VolunteerCenter/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VolPostgres struct {
	db *sqlx.DB
}

func (v VolPostgres) Create(volunteer models.Volunteer) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, second_name, patronymic, birth_date) "+
		"VALUES ($1, $2, $3, $4) RETURNING id", volsTable)
	row := v.db.QueryRow(query, volunteer.FirstName, volunteer.SecondName, volunteer.Patronymic, volunteer.BirthDate)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (v VolPostgres) Update(volunteer models.Volunteer) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, second_name = $2, patronymic = $3, birth_date = $4 WHERE id = $5", volsTable)
	_, err := v.db.Query(query, volunteer.FirstName, volunteer.SecondName, volunteer.Patronymic, volunteer.BirthDate, volunteer.Id)

	if err != nil {
		return err
	}

	return nil
}

func (v VolPostgres) GetById(id int) (models.Volunteer, error) {
	var vol models.Volunteer
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", volsTable)
	err := v.db.Get(&vol, query, id)

	return vol, err
}

func (v VolPostgres) GetAll() ([]models.Volunteer, error) {
	var items []models.Volunteer

	query := fmt.Sprintf(`SELECT * FROM %s`, volsTable)

	if err := v.db.Select(&items, query); err != nil {
		return items, err
	}

	return items, nil
}

func NewVolPostgres(db *sqlx.DB) *VolPostgres {
	return &VolPostgres{db: db}
}
