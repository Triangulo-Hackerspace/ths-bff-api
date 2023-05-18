package repository

import (
	"github.com/rogeriofontes/api-go-gin/database"
	"github.com/rogeriofontes/api-go-gin/models"
)

//CommentRepository -> CommentRepository
type CommentRepository struct {
	db database.Database
}

// NewCommentRepository : fetching database
func NewCommentRepository(db database.Database) CommentRepository {
	return CommentRepository{
		db: db,
	}
}

//Save -> Method for saving comment to database
func (p CommentRepository) Save(comment models.Comment) error {
	return p.db.DB.Create(&comment).Error
}

//FindAll -> Method for fetching all comments from database
func (p CommentRepository) FindAll(comment models.Comment, keyword string) (*[]models.Comment, int64, error) {
	var comments []models.Comment
	var totalRows int64 = 0

	queryBuider := p.db.DB.Order("created_at desc").Model(&models.Comment{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			p.db.DB.Where("comment.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(comment).
		Find(&comments).
		Count(&totalRows).Error
	return &comments, totalRows, err
}

//Update -> Method for updating Comment
func (p CommentRepository) Update(comment models.Comment) error {
	return p.db.DB.Save(&comment).Error
}

//Find -> Method for fetching comment by id
func (p CommentRepository) Find(comment models.Comment) (models.Comment, error) {
	var comments models.Comment
	err := p.db.DB.
		Debug().
		Model(&models.Comment{}).
		Where(&comment).
		Take(&comments).Error
	return comments, err
}

//Delete Deletes Comment
func (p CommentRepository) Delete(comment models.Comment) error {
	return p.db.DB.Delete(&comment).Error
}
