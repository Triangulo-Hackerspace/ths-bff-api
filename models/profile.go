package models

import (
	"gorm.io/gorm"
)

//Comment Post Model
type Profile struct {
	gorm.Model
	//UserID uint32 `gorm:"not null" json:"user_id,omitempty"`
	//User   User   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	//User      User   `gorm:"foreignkey:UserId;association_foreignkey:ID" json:"user_id,omitempty"`
	Bio       string `gorm:"size:2000" json:"bio"`
	Image     string `gorm:"size:300" json:"image"`
	User      User   `json:"user"`
	UserID    uint32 `sql:"type:int REFERENCES users(id)" json:"user_id"`
	Following bool   `json:"following"`
	Approved  bool   `json:"approved"`
	IsActive  bool   `json:"is_active"`
}

// TableName method that returns tablename of Post model
func (profile *Profile) TableName() string {
	return "profile"
}

//ResponseMap -> response map of Post
func (profile *Profile) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = profile.ID
	resp["user"] = profile.User
	resp["bio"] = profile.Bio
	resp["image"] = profile.Image
	resp["following"] = profile.Following
	resp["created_at"] = profile.CreatedAt
	resp["updated_at"] = profile.UpdatedAt
	return resp
}
