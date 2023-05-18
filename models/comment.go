package models

import (
	"gorm.io/gorm"
)

//Comment Post Model
type Comment struct {
	gorm.Model
	Title    string `gorm:"size:200" json:"title"`
	Comment  string `gorm:"size:3000" json:"comment"`
	Posts    []Post `gorm:"many2many:post_comment"`
	UserId   int    `gorm:"not null"`
	User     User   `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	Approved bool   `json:"approved"`
	IsActive bool   `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (comment *Comment) TableName() string {
	return "comment"
}

//ResponseMap -> response map of Post
func (comment *Comment) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = comment.ID
	resp["title"] = comment.Title
	resp["Comment"] = comment.Comment
	resp["Posts"] = comment.Posts
	resp["created_at"] = comment.CreatedAt
	resp["updated_at"] = comment.UpdatedAt
	return resp

}
