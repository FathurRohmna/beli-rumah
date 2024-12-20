package domain

import (
	"beli-tanah/model/web"
	"time"
)

type House struct {
	ID            string            `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Latitude      float64           `gorm:"not null" json:"latitude"`
	Longitude     float64           `gorm:"not null" json:"longitude"`
	Address       string            `gorm:"not null" json:"address"`
	Category      web.HouseCategory `gorm:"type:house_category;not null" json:"category"`
	PricePerMonth float64           `gorm:"not null" json:"price_per_month"`
	UnitCount     int               `gorm:"not null" json:"unit_count"`

	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`

	UserHouseTransactions []UserHouseTransaction `gorm:"foreignKey:HouseID" json:"user_house_transactions"`
}