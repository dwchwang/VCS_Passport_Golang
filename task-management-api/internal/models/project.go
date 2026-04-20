package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey"           json:"id"`
	Name        string    `gorm:"not null"             json:"name"`
	Description string    `                            json:"description"`
	OwnerID     uint      `gorm:"not null"             json:"owner_id"`
	Owner       *User     `gorm:"foreignKey:OwnerID"   json:"owner,omitempty"`
	Tasks       []Task    `gorm:"foreignKey:ProjectID" json:"tasks,omitempty"`
	CreatedAt   time.Time `                            json:"created_at"`
	UpdatedAt   time.Time `                            json:"updated_at"`
}