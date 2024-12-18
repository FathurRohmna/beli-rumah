package domain

import "time"

type Transaction struct {
	ID      string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID  string  `gorm:"type:uuid;not null" json:"user_id"`
	OrderID string  `gorm:"type:varchar(255);not null" json:"order_id"`
	Amount  float64 `gorm:"type:float;not null" json:"amount"`
	Status  string  `gorm:"type:varchar(50);default:'pending'" json:"status"`

	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
}
