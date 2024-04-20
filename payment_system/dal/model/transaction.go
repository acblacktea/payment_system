package model

import (
	"time"

	"gorm.io/gorm"
)

const TransactionTableName = "Transaction"

type Transaction struct {
	ID                   int64          `gorm:"primaryKey;column:id"`
	SourceAccountID      int64          `gorm:"column:source_account_id"`
	DestinationAccountID int64          `gorm:"column:destination_account_id"`
	Amount               int64          `gorm:"column:amount"`
	CreatedAt            time.Time      `gorm:"column:created_at"`
	UpdatedAt            time.Time      `gorm:"column:updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at"`
}
