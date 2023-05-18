package service

import (
	"github.com/rogeriofontes/api-go-gin/api/repository"
	"github.com/rogeriofontes/api-go-gin/models"
)

//CommentService CommentService struct
type CommentService struct {
	repository repository.CommentRepository
}

//NewCommentService : returns the CommentService struct instance
func NewCommentService(r repository.CommentRepository) CommentService {
	return CommentService{
		repository: r,
	}
}

//Save -> calls comment repository save method
func (p CommentService) Save(comment models.Comment) error {
	return p.repository.Save(comment)
}

//FindAll -> calls comment repo find all method
func (p CommentService) FindAll(comment models.Comment, keyword string) (*[]models.Comment, int64, error) {
	return p.repository.FindAll(comment, keyword)
}

// Update -> calls commentrepo update method
func (p CommentService) Update(comment models.Comment) error {
	return p.repository.Update(comment)
}

// Delete -> calls comment repo delete method
func (p CommentService) Delete(id uint) error {
	var comment models.Comment
	comment.ID = id
	return p.repository.Delete(comment)
}

// Find -> calls comment repo find method
func (p CommentService) Find(comment models.Comment) (models.Comment, error) {
	return p.repository.Find(comment)
}
