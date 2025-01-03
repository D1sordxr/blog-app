package models

import (
	"time"
)

type Base struct {
	ID        uint32     `gorm:"index;primary_key;"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP; not null"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP; not null"`
	DeletedAt *time.Time `gorm:"default:null"`
}
