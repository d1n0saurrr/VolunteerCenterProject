package service

import (
	"VolunteerCenter/models"
	"VolunteerCenter/pkg/repository"
)

type VolService struct {
	repo repository.Volunteer
}

func NewVolService(repo repository.Volunteer) *VolService {
	return &VolService{repo: repo}
}

func (v VolService) Create(volunteer models.Volunteer) (int, error) {
	return v.repo.Create(volunteer)
}

func (v VolService) Update(volunteer models.Volunteer) error {
	return v.repo.Update(volunteer)
}

func (v VolService) GetById(id int) (models.Volunteer, error) {
	return v.repo.GetById(id)
}

func (v VolService) GetAll() ([]models.Volunteer, error) {
	return v.repo.GetAll()
}
