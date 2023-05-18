package repository

import (
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/models"
)

//EventRepository -> EventRepository
type EventRepository struct {
	db database.Database
}

// NewEventRepository : fetching database
func NewEventRepository(db database.Database) EventRepository {
	return EventRepository{
		db: db,
	}
}

//Save -> Method for saving event to database
func (p EventRepository) Save(event models.Event) error {
	return p.db.DB.Create(&event).Error
}

//FindAll -> Method for fetching all events from database
func (p EventRepository) FindAll(event models.Event, keyword string) (*[]models.Event, int64, error) {
	var events []models.Event
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Event{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("event.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(event).
		Find(&events).
		Count(&totalRows).Error
	return &events, totalRows, err
}

//Update -> Method for updating Event
func (p EventRepository) Update(event models.Event) error {
	return p.db.DB.Save(&event).Error
}

//Find -> Method for fetching event by id
func (p EventRepository) Find(event models.Event) (models.Event, error) {
	var events models.Event
	err := p.db.DB.
		Debug().
		Model(&models.Event{}).
		Where(&event).
		Take(&events).Error
	return events, err
}

//Delete Deletes Event
func (p EventRepository) Delete(event models.Event) error {
	return p.db.DB.Delete(&event).Error
}
