package wallet

import (
	"context"

	"github.com/acblacktea/payment_system/payment_system/dal/repo"
	"github.com/acblacktea/payment_system/payment_system/util"
)

func (s *WalletServiceImpl) SubmitTransaction(ctx context.Context, SourceAccountID, destinationAccountID int64, amount float64) error {
	if amount <= 0 {
		return util.ErrInvalidMoneyValue
	}

	wallets, err := s.WalletRepo.GetAccounts(ctx, []int64{SourceAccountID, destinationAccountID})
	if err != nil {
		return err
	}

	if len(wallets) < 2 {
		return util.ErrAccountIDNotExist
	}

	var sourceAccount, destinationAccount WalletInfo
	for _, wallet := range wallets {
		if wallet.ID == SourceAccountID {
			sourceAccount = WalletInfo{
				AccountID:    wallet.ID,
				BalanceValue: wallet.Balance,
				Currency:     wallet.Currency,
			}
		} else {
			destinationAccount = WalletInfo{
				AccountID:    wallet.ID,
				BalanceValue: wallet.Balance,
				Currency:     wallet.Currency,
			}
		}
	}

	if sourceAccount.Currency != destinationAccount.Currency {
		return util.ErrCurrencyTypeMismatch
	}

	amountIntValue, err := convertMoneyToIntegerMoney(amount, sourceAccount.Currency)
	if err != nil {
		return err
	}

	if sourceAccount.BalanceValue < amountIntValue {
		return util.ErrBalanceNotEnough
	}

	return s.WalletRepo.SubmitTransaction(ctx, repo.SubmitTransactionParameters{
		SourceAccountID:      SourceAccountID,
		DestinationAccountID: destinationAccountID,
		SourceBalance:        sourceAccount.BalanceValue - amountIntValue,
		DestinationBalance:   destinationAccount.BalanceValue + amountIntValue,
		Amount:               amountIntValue,
		Currency:             sourceAccount.Currency,
	})
}
