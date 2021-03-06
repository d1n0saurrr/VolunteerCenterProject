package service

import (
	"VolunteerCenter/models"
	"VolunteerCenter/pkg/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (e EventService) Create(event models.Event) (int, error) {
	return e.repo.Create(event)
}

func (e EventService) GetAll() ([]models.Event, error) {
	return e.repo.GetAll()
}

func (e EventService) GetNew() ([]models.Event, error) {
	return e.repo.GetNew()
}

func (e EventService) Delete(id int) error {
	return e.repo.Delete(id)
}

func (e EventService) GetVolEvents(volId int) ([]models.Event, error) {
	return e.repo.GetVolEvents(volId)
}

func (e EventService) GetOldVolEvents(volId int) ([]models.Event, error) {
	return e.repo.GetOldVolEvents(volId)
}

func (e EventService) RegisterVol(volId int, eventId int) error {
	return e.repo.RegisterVol(volId, eventId)
}
