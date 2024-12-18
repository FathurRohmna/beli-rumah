package domain

import "time"

type House struct {
	ID        string    `gorm:"primaryKey"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	Address   string    `gorm:"not null"`
	Category  string    `gorm:"not null"`
	Size      int       `gorm:"not null"`
	UnitCount int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
