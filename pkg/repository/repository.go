package repository

import (
	"VolunteerCenter/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type User interface {
	GetById(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAll() ([]models.User, error)
	SetVolId(userId, volId int) error
}

type Volunteer interface {
	Create(volunteer models.Volunteer) (int, error)
	Update(volunteer models.Volunteer) error
	GetById(id int) (models.Volunteer, error)
	GetAll() ([]models.Volunteer, error)
}

type Event interface {
	Create(event models.Event) (int, error)
	GetAll() ([]models.Event, error)
	Delete(id int) error
	GetVolEvents(volId int) ([]models.Event, error)
	RegisterVol(volId int, eventId int) error
}

type Repository struct {
	Authorization
	User
	Volunteer
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Volunteer:     NewVolPostgres(db),
		Event:         NewEventPostgres(db),
	}
}
