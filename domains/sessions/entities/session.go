package entities

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID
	RefreshToken string    `gorm:"type:varchar(255)"`
	IsRevoked    int       `gorm:"default:1"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	ExpiredAt    time.Time `gorm:"type:timestamp"`
}
