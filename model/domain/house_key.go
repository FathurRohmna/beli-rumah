package domain

import (
	"time"
)

type HouseKey struct {
	ID            string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TransactionID string `gorm:"type:uuid;not null;" json:"transaction_id"`

	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
}
