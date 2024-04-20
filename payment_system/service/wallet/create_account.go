package wallet

import (
	"context"
	"fmt"
	"strings"

	"github.com/acblacktea/payment_system/payment_system/util"
)

func (s *WalletServiceImpl) CreateAccount(ctx context.Context, accountID int64, initBalance float64, currency string) error {
	if initBalance < 0 {
		return util.ErrInvalidMoneyValue
	}

	balance, err := convertMoneyToIntegerMoney(initBalance, currency)
	if err != nil {
		return err
	}

	err = s.WalletRepo.CreateAccount(ctx, accountID, balance, currency)
	if strings.Contains(err.Error(), util.GormExistItem) {
		return fmt.Errorf("%w: id %d", util.ErrDuplicateAccountID, accountID)
	}

	return err
}
