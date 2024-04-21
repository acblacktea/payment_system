package model

import (
	"time"

	"gorm.io/gorm"
)

const WalletTableName = "Wallet"

type Wallet struct {
	ID        int64          `gorm:"primaryKey;column:id"`
	Balance   int64          `gorm:"column:balance"`
	Currency  string         `gorm:"column:currency"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
