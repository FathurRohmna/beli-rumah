package domain

import "time"

type UserHouseTransaction struct {
	ID                string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID            string    `gorm:"type:uuid;not null" json:"user_id"`
	HouseID           string    `gorm:"type:uuid;not null" json:"house_id"`
	TransactionStatus string    `gorm:"type:house_availability;not null;default:'pending'" json:"transaction_status"`
	StartDate         time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate           time.Time `gorm:"type:date" json:"end_date"`
	ExpiredAt         time.Time `gorm:"not null" json:"expired_at"`

	House House `gorm:"foreignKey:HouseID"`
}
