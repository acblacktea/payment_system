package wallet

import (
	"context"

	"github.com/acblacktea/payment_system/payment_system/util"
)

func (s *WalletServiceImpl) GetAccount(ctx context.Context, accountID int64) (*WalletInfo, error) {
	wallets, err := s.WalletRepo.GetAccounts(ctx, []int64{accountID})
	if err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		return nil, util.ErrAccountIDNotExist
	}

	balanceIntValue, err := convertIntegerMoneyToMoney(wallets[0].Balance, wallets[0].Currency)
	if err != nil {
		return nil, err
	}

	return &WalletInfo{
		AccountID: wallets[0].ID,
		Balance:   balanceIntValue,
		Currency:  wallets[0].Currency,
	}, nil
}
