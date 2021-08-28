package service

import (
	"VolunteerCenter/models"
	"VolunteerCenter/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, bool, error)
}

type User interface {
	GetById(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAll() ([]models.User, error)
	Update(user models.User) error
	SetVolId(userId, volId int) error
	Delete(id int) error
}

type Volunteer interface {
	Create(volunteer models.Volunteer) (int, error)
	Update(volunteer models.Volunteer) error
	GetById(id int) (models.Volunteer, error)
	GetAll() ([]models.Volunteer, error)
	Delete(id int) error
}

type Event interface {
	Create(event models.Event) (int, error)
	GetAll() ([]models.Event, error)
	GetNew() ([]models.Event, error)
	Delete(id int) error
	GetVolEvents(volId int) ([]models.Event, error)
	GetOldVolEvents(volId int) ([]models.Event, error)
	RegisterVol(volId int, eventId int) error
}

type Service struct {
	Authorization
	User
	Volunteer
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Volunteer:     NewVolService(repos.Volunteer),
		Event:         NewEventService(repos.Event),
	}
}
