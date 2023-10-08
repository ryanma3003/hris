package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	base.ID = uuid.NewString()
	return
}
