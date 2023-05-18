package models

import (
	"gorm.io/gorm"
)

//Post Post Model
type Post struct {
	gorm.Model
	Title    string    `gorm:"size:200" json:"title"`
	Body     string    `gorm:"size:3000" json:"body"`
	Comments []Comment `gorm:"many2many:post_comment"`
	UserId   int       `gorm:"not null"`
	User     User      `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	Approved bool      `json:"approved"`
	IsActive bool      `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (post *Post) TableName() string {
	return "post"
}

//ResponseMap -> response map of Post
func (post *Post) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = post.ID
	resp["title"] = post.Title
	resp["body"] = post.Body
	resp["Comments"] = post.Comments
	resp["created_at"] = post.CreatedAt
	resp["updated_at"] = post.UpdatedAt
	return resp

}
