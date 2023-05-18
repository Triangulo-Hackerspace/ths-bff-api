package models

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

//Event Post Model
type Event struct {
	gorm.Model
	Name        string    `gorm:"size:256" json:"name"`
	Description string    `gorm:"size:256" json:"description"`
	Slug        string    `gorm:"size:256" json:"slug"`
	Link        string    `gorm:"size:512" json:"link"`
	Date        time.Time `gorm:"not null" json:"date"`
	Users       []User    `gorm:"many2many:user_event"`
	UserId      int       `gorm:"not null"`
	User        User      `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	Approved    bool      `json:"approved"`
	IsActive    bool      `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (event *Event) TableName() string {
	return "event"
}

//ResponseMap -> response map of Post
func (event *Event) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = event.ID
	resp["title"] = event.Name
	resp["Description"] = event.Description
	resp["Slug"] = slug.Make(event.Name)
	resp["link"] = event.Link
	resp["date"] = event.Date
	resp["Users"] = event.Users
	resp["created_at"] = event.CreatedAt
	resp["updated_at"] = event.UpdatedAt
	return resp

}
