package repository

import (
	"VolunteerCenter/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func (e EventPostgres) Create(event models.Event) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, location, start_date, end_date) "+
		"values ($1, $2, $3, $4, $5) RETURNING id", eventsTable)
	row := e.db.QueryRow(query, event.Name, event.Description, event.Location, event.StartDate, event.EndDate)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (e EventPostgres) GetAll() ([]models.Event, error) {
	var items []models.Event

	query := fmt.Sprintf(`SELECT * FROM %s`, eventsTable)

	if err := e.db.Select(&items, query); err != nil {
		return items, err
	}

	return items, nil
}

func (e EventPostgres) Delete(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE event_id = $1`, volsAndEvents)

	rows, err := e.db.Query(query, id)

	if err != nil {
		return err
	}

	err = rows.Close()

	if err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, eventsTable)

	rows, err = e.db.Query(query, id)

	if err != nil {
		return err
	}

	err = rows.Close()

	if err != nil {
		return err
	}

	return nil
}

func (e EventPostgres) GetVolEvents(volId int) ([]models.Event, error) {
	var items []models.Event

	query := fmt.Sprintf(`SELECT * FROM %s WHERE vol_id = $1`, volsAndEvents)

	if err := e.db.Select(&items, query, volId); err != nil {
		return items, err
	}

	return items, nil
}

func (e EventPostgres) RegisterVol(volId int, eventId int) error {
	query := fmt.Sprintf("INSERT INTO %s (vol_id, event_id) VALUES ($1, $2))", volsAndEvents)
	row := e.db.QueryRow(query, volId, eventId)

	return row.Err()
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}
