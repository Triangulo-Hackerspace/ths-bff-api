package service

import (
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/models"
)

//ProfileService ProfileService struct
type ProfileService struct {
	repository repository.ProfileRepository
}

//NewProfileService : returns the ProfileService struct instance
func NewProfileService(r repository.ProfileRepository) ProfileService {
	return ProfileService{
		repository: r,
	}
}

//Save -> calls Profile repository save method
func (p ProfileService) Save(Profile models.Profile) error {
	return p.repository.Save(Profile)
}

//FindAll -> calls Profile repo find all method
func (p ProfileService) FindAll(Profile models.Profile, keyword string) (*[]models.Profile, int64, error) {
	return p.repository.FindAll(Profile, keyword)
}

// Update -> calls Profilerepo update method
func (p ProfileService) Update(Profile models.Profile) error {
	return p.repository.Update(Profile)
}

// Delete -> calls Profile repo delete method
func (p ProfileService) Delete(id uint) error {
	var Profile models.Profile
	Profile.ID = id
	return p.repository.Delete(Profile)
}

// Find -> calls Profile repo find method
func (p ProfileService) Find(Profile models.Profile) (models.Profile, error) {
	return p.repository.Find(Profile)
}

// FindByUserId -> calls Profile repo find method
func (p ProfileService) FindByUserId(UserId uint64) (models.Profile, error) {
	return p.repository.FindByUserId(UserId)
}
