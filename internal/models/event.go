package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	OwnerID     uint      `gorm:"not null"`
	Owner       User      `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Location    string    `gorm:"not null"`
	Attendees   []Attendee
}
