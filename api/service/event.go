package service

import (
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/models"
)

//EventService EventService struct
type EventService struct {
	repository repository.EventRepository
}

//NewEventService : returns the EventService struct instance
func NewEventService(r repository.EventRepository) EventService {
	return EventService{
		repository: r,
	}
}

//Save -> calls event repository save method
func (p EventService) Save(event models.Event) error {
	return p.repository.Save(event)
}

//FindAll -> calls event repo find all method
func (p EventService) FindAll(event models.Event, keyword string) (*[]models.Event, int64, error) {
	return p.repository.FindAll(event, keyword)
}

// Update -> calls eventrepo update method
func (p EventService) Update(event models.Event) error {
	return p.repository.Update(event)
}

// Delete -> calls event repo delete method
func (p EventService) Delete(id uint) error {
	var event models.Event
	event.ID = id
	return p.repository.Delete(event)
}

// Find -> calls event repo find method
func (p EventService) Find(event models.Event) (models.Event, error) {
	return p.repository.Find(event)
}
