package repository

import (
	"VolunteerCenter/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u UserPostgres) GetById(id int) (models.User, error) {
	var item models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, usersTable)

	if err := u.db.Get(&item, query, id); err != nil {
		return item, err
	}

	return item, nil
}

func (u UserPostgres) GetByUsername(username string) (models.User, error) {
	var item models.User

	query := fmt.Sprintf(`SELECT * FROM %s WHERE username = $1`, usersTable)

	if err := u.db.Get(&item, query, username); err != nil {
		return item, err
	}

	return item, nil
}

func (u UserPostgres) GetAll() ([]models.User, error) {
	var items []models.User

	query := fmt.Sprintf(`SELECT * FROM %s`, usersTable)

	if err := u.db.Select(&items, query); err != nil {
		return items, err
	}

	return items, nil
}

func (u UserPostgres) SetVolId(userId, volId int) error {
	query := fmt.Sprintf(`UPDATE %s SET id_volunteer = $1 WHERE id = $2`, usersTable)

	res, err := u.db.Query(query, volId, userId)

	if err != nil {
		return err
	}

	if err = res.Close(); err != nil {
		return err
	}

	return nil
}
