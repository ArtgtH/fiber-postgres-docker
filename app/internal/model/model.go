package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Car struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	ModelName string
	Price     int64
	Country   string
	Producer  string
	Date      time.Time
}
