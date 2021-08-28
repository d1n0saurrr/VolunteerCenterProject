package repository

import (
	"VolunteerCenter/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
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

	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY end_date DESC`, eventsTable)

	if err := e.db.Select(&items, query); err != nil {
		return items, err
	}

	return items, nil
}

func (e EventPostgres) GetNew() ([]models.Event, error) {
	var items []models.Event

	query := fmt.Sprintf(`SELECT * FROM %s WHERE end_date >= $1 ORDER BY end_date DESC`, eventsTable)

	if err := e.db.Select(&items, query, time.Now()); err != nil {
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

	query := fmt.Sprintf(`SELECT id, name, description, start_date, end_date
		FROM %s JOIN %s ON %s.event_id = %s.id
		WHERE vol_id = $1`, volsAndEvents, eventsTable, volsAndEvents, eventsTable)

	if err := e.db.Select(&items, query, volId); err != nil {
		return items, err
	}

	return items, nil
}

func (e EventPostgres) GetOldVolEvents(volId int) ([]models.Event, error) {
	var items []models.Event

	query := fmt.Sprintf(`SELECT id, name, description, start_date, end_date
		FROM %s JOIN %s ON %s.event_id = %s.id
		WHERE vol_id = $1 AND end_date < $2
		ORDER BY end_date DESC`, volsAndEvents, eventsTable, volsAndEvents, eventsTable)

	if err := e.db.Select(&items, query, volId, time.Now()); err != nil {
		return items, err
	}

	return items, nil
}

func (e EventPostgres) RegisterVol(volId int, eventId int) error {
	var count []int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE vol_id = $1 AND event_id = $2", volsAndEvents)
	err := e.db.Select(&count, query, volId, eventId)

	if err != nil {
		return err
	}

	if count[0] != 0 {
		return errors.New("can't register twice volunteer to event")
	}

	query = fmt.Sprintf("INSERT INTO %s VALUES ($1, $2)", volsAndEvents)
	_, err = e.db.Query(query, volId, eventId)

	return err
}
