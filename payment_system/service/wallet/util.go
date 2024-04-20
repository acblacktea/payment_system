package wallet

import "github.com/acblacktea/payment_system/payment_system/util"

type WalletInfo struct {
	AccountID    int64
	Balance      float64
	Currency     string
	BalanceValue int64
}

func convertMoneyToIntegerMoney(money float64, currency string) (int64, error) {
	if currency == util.Dollar {
		return int64(money * 100), nil
	}

	return -1, util.ErrInvalidCurrency
}

func convertIntegerMoneyToMoney(moneyIntValue int64, currency string) (float64, error) {
	if currency == util.Dollar {
		return float64(moneyIntValue) / 100, nil
	}

	return -1, util.ErrInvalidCurrency
}
