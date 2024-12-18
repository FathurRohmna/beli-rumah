package domain

import "time"

type House struct {
	ID        string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Address   string  `gorm:"not null" json:"address"`
	Category  string  `gorm:"type:house_category;not null" json:"category"`
	Size      int     `gorm:"not null" json:"size"`
	UnitCount int     `gorm:"not null" json:"unit_count"`

	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
}
