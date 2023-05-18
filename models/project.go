package models

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

//Project Post Model
type Project struct {
	gorm.Model
	Title   string    `gorm:"size:200" json:"title"`
	Content string    `gorm:"size:6000" json:"content"`
	Slug    string    `gorm:"size:256" json:"slug"`
	Link    string    `gorm:"size:512" json:"link"`
	Date    time.Time `gorm:"autoCreateTime" json:"date"`
	Users   []User    `gorm:"many2many:user_project"`
	//UserId   int       `gorm:"not null"`
	//User     User      `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	User     User   `json:"user"`
	UserID   uint32 `sql:"type:int REFERENCES users(id)" json:"user_id"`
	Approved bool   `json:"approved"`
	IsActive bool   `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (project *Project) TableName() string {
	return "project"
}

//ResponseMap -> response map of Post
func (project *Project) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = project.ID
	resp["title"] = project.Title
	resp["Content"] = project.Content
	resp["Slug"] = slug.Make(project.Title)
	resp["link"] = project.Link
	resp["date"] = project.Date
	resp["Users"] = project.Users
	resp["user"] = project.User
	resp["created_at"] = project.CreatedAt
	resp["updated_at"] = project.UpdatedAt
	return resp

}
