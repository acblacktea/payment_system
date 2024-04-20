package util

import (
	"time"

	"github.com/google/uuid"
)

// money currency string
const (
	Dollar          = "USD"
	ChineseYuan     = "CNY"
	SingaporeDollar = "SGD"
	BTC             = "BTC"
	ETH             = "ETH"
	USDT            = "USDT"
	USDC            = "USDC"
)

func GetUniqueID() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	ID := currentTimestamp + int64(uniqueID)
	return ID
}
