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

func ConvertMoneyToIntegerMoney(money float64, currency string) (int64, error) {
	if currency == Dollar {
		return int64(money*100.0 + 0.00001), nil
	}

	return -1, ErrInvalidCurrency
}

func ConvertIntegerMoneyToMoney(moneyIntValue int64, currency string) (float64, error) {
	if currency == Dollar {
		return float64(moneyIntValue) / 100, nil
	}

	return -1, ErrInvalidCurrency
}
