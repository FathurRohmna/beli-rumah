package domain

import "time"

type UserHouseTransaction struct {
	ID                string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID            string    `gorm:"type:uuid;not null" json:"user_id"`
	HouseID           string    `gorm:"type:uuid;not null" json:"house_id"`
	TransactionStatus string    `gorm:"type:house_availability;not null;default:'pending'" json:"transaction_status"`
	
	ExpiredAt         time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"updated_at"`
}