package payment_system

import (
	"context"

	"github.com/acblacktea/payment_system/payment_system/service/wallet"
)

type PaymentSystemService interface {
	CreateAccount(ctx context.Context, accountID int64, initBalance float64, currency string) error
	GetAccount(ctx context.Context, accountID int64) (*wallet.WalletInfo, error)
	SubmitTransaction(ctx context.Context, SourceAccountID, destinationAccountID int64, amount float64) error
}

type PaymentSystemServiceImpl struct {
	WalletService wallet.WalletService
}

func CreatePaymentSystemService(walletService wallet.WalletService) PaymentSystemService {
	return &PaymentSystemServiceImpl{
		WalletService: walletService,
	}
}

func (s *PaymentSystemServiceImpl) CreateAccount(ctx context.Context, accountID int64, initBalance float64, currency string) error {
	return s.WalletService.CreateAccount(ctx, accountID, initBalance, currency)
}

func (s *PaymentSystemServiceImpl) GetAccount(ctx context.Context, accountID int64) (*wallet.WalletInfo, error) {
	return s.WalletService.GetAccount(ctx, accountID)
}

func (s *PaymentSystemServiceImpl) SubmitTransaction(ctx context.Context, SourceAccountID, destinationAccountID int64, amount float64) error {
	return s.WalletService.SubmitTransaction(ctx, SourceAccountID, destinationAccountID, amount)
}
