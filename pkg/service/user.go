package service

import (
	"VolunteerCenter/models"
	"VolunteerCenter/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) GetById(id int) (models.User, error) {
	return u.repo.GetById(id)
}

func (u UserService) GetByUsername(username string) (models.User, error) {
	return u.repo.GetByUsername(username)
}

func (u UserService) GetAll() ([]models.User, error) {
	return u.repo.GetAll()
}

func (u UserService) SetVolId(userId, volId int) error {
	return u.repo.SetVolId(userId, volId)
}

func (u UserService) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u UserService) Update(user models.User) error {
	return u.repo.Update(user)
}
