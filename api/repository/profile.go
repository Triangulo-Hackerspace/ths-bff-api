package repository

import (
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/models"
)

//ProfileRepository -> ProfileRepository
type ProfileRepository struct {
	db database.Database
}

// NewProfileRepository : fetching database
func NewProfileRepository(db database.Database) ProfileRepository {
	return ProfileRepository{
		db: db,
	}
}

//Save -> Method for saving Profile to database
func (p ProfileRepository) Save(Profile models.Profile) error {
	return p.db.DB.Create(&Profile).Error
}

//FindAll -> Method for fetching all Profiles from database
func (p ProfileRepository) FindAll(Profile models.Profile, keyword string) (*[]models.Profile, int64, error) {
	var Profiles []models.Profile
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Profile{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("Profile.Bio LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(Profile).
		Find(&Profiles).
		Count(&totalRows).Error

	if len(Profiles) > 0 {
		for i, _ := range Profiles {
			err := p.db.DB.Debug().Model(&models.User{}).Where("id = ?", Profiles[i].UserID).Take(&Profiles[i].User).Error
			if err != nil {
				return &Profiles, totalRows, err
			}
		}
	}

	return &Profiles, totalRows, err
}

//Update -> Method for updating Profile
func (p ProfileRepository) Update(Profile models.Profile) error {
	return p.db.DB.Save(&Profile).Error
}

//Find -> Method for fetching Profile by id
func (p ProfileRepository) Find(Profile models.Profile) (models.Profile, error) {
	var Profiles models.Profile
	err := p.db.DB.
		Debug().
		Model(&models.Profile{}).
		Where(&Profile).
		Take(&Profiles).Error

	if Profiles.ID != 0 {
		err := p.db.DB.
			Debug().
			Model(&models.User{}).
			Where("id = ?", Profiles.UserID).
			Take(&Profiles.User).Error

		if err != nil {
			return Profiles, err
		}
	}
	return Profiles, err
}

func (p ProfileRepository) FindByUserId(UserId uint64) (models.Profile, error) {
	var Profiles models.Profile

	err := p.db.DB.
		Debug().
		Model(&models.Profile{}).
		Where("user_id = ?", UserId).
		Take(&Profiles).Error

	if err != nil {
		return Profiles, err
	}

	if Profiles.ID != 0 {
		err := p.db.DB.
			Debug().
			Model(&models.User{}).
			Where("id = ?", Profiles.UserID).
			Take(&Profiles.User).Error

		if err != nil {
			return Profiles, err
		}
	}

	return Profiles, err
}

//Delete Deletes Profile
func (p ProfileRepository) Delete(Profile models.Profile) error {
	return p.db.DB.Delete(&Profile).Error
}
