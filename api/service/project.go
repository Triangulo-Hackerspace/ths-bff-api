package service

import (
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/models"
)

//ProjectService ProjectService struct
type ProjectService struct {
	repository repository.ProjectRepository
}

//NewProjectService : returns the ProjectService struct instance
func NewProjectService(r repository.ProjectRepository) ProjectService {
	return ProjectService{
		repository: r,
	}
}

//Save -> calls project repository save method
func (p ProjectService) Save(project models.Project) error {
	return p.repository.Save(project)
}

//FindAll -> calls project repo find all method
func (p ProjectService) FindAll(project models.Project, keyword string) (*[]models.Project, int64, error) {
	return p.repository.FindAll(project, keyword)
}

// Update -> calls projectrepo update method
func (p ProjectService) Update(project models.Project) error {
	return p.repository.Update(project)
}

// Delete -> calls project repo delete method
func (p ProjectService) Delete(id uint) error {
	var project models.Project
	project.ID = id
	return p.repository.Delete(project)
}

// Find -> calls project repo find method
func (p ProjectService) Find(project models.Project) (models.Project, error) {
	return p.repository.Find(project)
}
