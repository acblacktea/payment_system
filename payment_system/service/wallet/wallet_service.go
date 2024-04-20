package wallet

import (
	"context"

	"github.com/acblacktea/payment_system/payment_system/dal/repo"
)

type WalletService interface {
	CreateAccount(ctx context.Context, accountID int64, initBalance float64, currency string) error
	GetAccount(ctx context.Context, accountID int64) (*WalletInfo, error)
	SubmitTransaction(ctx context.Context, SourceAccountID, destinationAccountID int64, amount float64) error
}

type WalletServiceImpl struct {
	WalletRepo repo.WalletRepo
}

func CreateWalletService(walletRepo repo.WalletRepo) WalletService {
	return &WalletServiceImpl{
		WalletRepo: walletRepo,
	}
}
