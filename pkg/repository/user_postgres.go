package repository

import (
	"VolunteerCenter/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u UserPostgres) Update(user models.User) error {
	isLastAdmin, err := u.isLastAdmin(user.Id)

	if err != nil {
		return err
	}

	if isLastAdmin {
		return errors.New("can't delete last admin")
	}

	query := fmt.Sprintf("UPDATE %s SET is_Admin = $1 WHERE id = $2", usersTable)

	if _, err := u.db.Query(query, user.IsAdmin, user.Id); err != nil {
		return err
	}

	return nil
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

func (u UserPostgres) Delete(id int) error {
	isLastAdmin, err := u.isLastAdmin(id)

	if err != nil {
		return err
	}

	if isLastAdmin {
		return errors.New("can't delete last admin")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, usersTable)

	_, err = u.db.Query(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserPostgres) isLastAdmin(id int) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE is_admin = true", usersTable)
	var counts []int
	var ids []int

	if err := u.db.Select(&counts, query); err != nil {
		return false, err
	}

	if counts[0] == 1 {
		query = fmt.Sprintf("SELECT id FROM %s WHERE is_admin = true", usersTable)

		if err := u.db.Select(&ids, query); err != nil {
			return false, err
		}

		if ids[0] == id {
			return true, nil
		}
	}

	return false, nil
}
