package domain

import "time"

type User struct {
	ID           string  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string  `gorm:"type:varchar(255);not null" json:"name"`
	Email        string  `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password     string  `gorm:"type:text;not null" json:"password"`
	WalletAmount float64 `gorm:"not null" json:"wallet_amount"`

	CreatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
}
