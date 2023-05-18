package models

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

//Meet Post Model
type Meet struct {
	gorm.Model
	Name        string    `gorm:"size:256" json:"name"`
	Description string    `gorm:"size:256" json:"description"`
	Slug        string    `gorm:"size:256" json:"slug"`
	Link        string    `gorm:"size:512" json:"link"`
	Date        time.Time `gorm:"not null" json:"date"`
	Users       []User    `gorm:"many2many:user_meet"`
	UserId      int       `gorm:"not null"`
	User        User      `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	Approved    bool      `json:"approved"`
	IsActive    bool      `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (meet *Meet) TableName() string {
	return "meet"
}

//ResponseMap -> response map of Post
func (meet *Meet) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = meet.ID
	resp["title"] = meet.Name
	resp["Description"] = meet.Description
	resp["Slug"] = slug.Make(meet.Name)
	resp["link"] = meet.Link
	resp["date"] = meet.Date
	resp["Users"] = meet.Users
	resp["created_at"] = meet.CreatedAt
	resp["updated_at"] = meet.UpdatedAt
	return resp

}
