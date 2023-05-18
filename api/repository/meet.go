package repository

import (
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/models"
)

//MeetRepository -> MeetRepository
type MeetRepository struct {
	db database.Database
}

// NewMeetRepository : fetching database
func NewMeetRepository(db database.Database) MeetRepository {
	return MeetRepository{
		db: db,
	}
}

//Save -> Method for saving meet to database
func (p MeetRepository) Save(meet models.Meet) error {
	return p.db.DB.Create(&meet).Error
}

//FindAll -> Method for fetching all meets from database
func (p MeetRepository) FindAll(meet models.Meet, keyword string) (*[]models.Meet, int64, error) {
	var meets []models.Meet
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Meet{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("meet.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(meet).
		Find(&meets).
		Count(&totalRows).Error
	return &meets, totalRows, err
}

//Update -> Method for updating Meet
func (p MeetRepository) Update(meet models.Meet) error {
	return p.db.DB.Save(&meet).Error
}

//Find -> Method for fetching meet by id
func (p MeetRepository) Find(meet models.Meet) (models.Meet, error) {
	var meets models.Meet
	err := p.db.DB.
		Debug().
		Model(&models.Meet{}).
		Where(&meet).
		Take(&meets).Error
	return meets, err
}

//Delete Deletes Meet
func (p MeetRepository) Delete(meet models.Meet) error {
	return p.db.DB.Delete(&meet).Error
}
