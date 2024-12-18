package domain

import "time"

type UserHouseTransaction struct {
	ID                string    `gorm:"primaryKey"`
	UserID            string    `gorm:"not null"`
	HouseID           string    `gorm:"not null"`
	TransactionStatus string    `gorm:"not null;default:'pending'"`
	ExpiredAt         time.Time `gorm:"not null"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}