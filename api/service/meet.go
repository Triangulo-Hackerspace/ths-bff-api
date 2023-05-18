package service

import (
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/models"
)

//MeetService MeetService struct
type MeetService struct {
	repository repository.MeetRepository
}

//NewMeetService : returns the MeetService struct instance
func NewMeetService(r repository.MeetRepository) MeetService {
	return MeetService{
		repository: r,
	}
}

//Save -> calls meet repository save method
func (p MeetService) Save(meet models.Meet) error {
	return p.repository.Save(meet)
}

//FindAll -> calls meet repo find all method
func (p MeetService) FindAll(meet models.Meet, keyword string) (*[]models.Meet, int64, error) {
	return p.repository.FindAll(meet, keyword)
}

// Update -> calls meetrepo update method
func (p MeetService) Update(meet models.Meet) error {
	return p.repository.Update(meet)
}

// Delete -> calls meet repo delete method
func (p MeetService) Delete(id uint) error {
	var meet models.Meet
	meet.ID = id
	return p.repository.Delete(meet)
}

// Find -> calls meet repo find method
func (p MeetService) Find(meet models.Meet) (models.Meet, error) {
	return p.repository.Find(meet)
}
