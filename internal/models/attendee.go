package models

import "gorm.io/gorm"

type Attendee struct {
	gorm.Model
	UserID  uint  `gorm:"not null"`
	User    User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	EventID uint  `gorm:"not null"`
	Event   Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}
