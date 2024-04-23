package wallet

import (
	"context"
)

func (s *WalletServiceImpl) SubmitTransaction(ctx context.Context, SourceAccountID, destinationAccountID int64, amount float64) error {
	return s.WalletRepo.SubmitTransaction(ctx, SourceAccountID, destinationAccountID, amount)
}
