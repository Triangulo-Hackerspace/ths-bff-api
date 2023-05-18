package repository

import (
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/models"
)

//ProjectRepository -> ProjectRepository
type ProjectRepository struct {
	db database.Database
}

// NewProjectRepository : fetching database
func NewProjectRepository(db database.Database) ProjectRepository {
	return ProjectRepository{
		db: db,
	}
}

//Save -> Method for saving project to database
func (p ProjectRepository) Save(project models.Project) error {
	return p.db.DB.Create(&project).Error
}

//FindAll -> Method for fetching all projects from database
func (p ProjectRepository) FindAll(project models.Project, keyword string) (*[]models.Project, int64, error) {
	var projects []models.Project
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Project{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("project.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(project).
		Find(&projects).
		Count(&totalRows).Error

	if len(projects) > 0 {
		for i, _ := range projects {
			err := p.db.DB.Debug().Model(&models.User{}).Where("id = ?", projects[i].UserID).Take(&projects[i].User).Error
			if err != nil {
				return &projects, totalRows, err
			}
		}
	}

	return &projects, totalRows, err
}

//Update -> Method for updating Project
func (p ProjectRepository) Update(project models.Project) error {
	return p.db.DB.Save(&project).Error
}

//Find -> Method for fetching project by id
func (p ProjectRepository) Find(project models.Project) (models.Project, error) {
	var projects models.Project
	err := p.db.DB.
		Debug().
		Model(&models.Project{}).
		Where(&project).
		Take(&projects).Error

	if projects.ID != 0 {
		err := p.db.DB.
			Debug().
			Model(&models.User{}).
			Where("id = ?", projects.UserID).
			Take(&projects.User).Error

		if err != nil {
			return projects, err
		}
	}
	return projects, err
}

//Delete Deletes Project
func (p ProjectRepository) Delete(project models.Project) error {
	return p.db.DB.Delete(&project).Error
}
