package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/acblacktea/payment_system/payment_system/dal/model"
	"github.com/acblacktea/payment_system/payment_system/util"
	"gorm.io/gorm"
)

type WalletRepo interface {
	CreateAccount(ctx context.Context, accountID, initBalance int64, currency string) error
	GetAccounts(ctx context.Context, accountIDs []int64) ([]*model.Wallet, error)
	SubmitTransaction(ctx context.Context, sourceAccountID, destinationAccountID int64, amount float64) error
}

type WalletRepoImpl struct {
	dbClient *gorm.DB
}

func CreateWalletRepo(dbClient *gorm.DB) WalletRepo {
	return &WalletRepoImpl{
		dbClient: dbClient,
	}
}

func (r *WalletRepoImpl) CreateAccount(ctx context.Context, accountID, initBalance int64, currency string) error {
	user := model.Wallet{
		ID:       accountID,
		Balance:  initBalance,
		Currency: currency,
	}

	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", r.dbClient)
	fmt.Printf("%v\n", r.dbClient.Table(model.WalletTableName))
	return r.dbClient.Table(model.WalletTableName).Create(user).Error
}

func (r *WalletRepoImpl) GetAccounts(ctx context.Context, accountIDs []int64) ([]*model.Wallet, error) {
	wallets := make([]*model.Wallet, len(accountIDs))
	result := r.dbClient.Table(model.WalletTableName).Where("id in ?", accountIDs).Find(&wallets)
	return wallets, result.Error
}

func (r *WalletRepoImpl) SubmitTransaction(ctx context.Context, sourceAccountID, destinationAccountID int64, amount float64) error {
	return r.dbClient.Transaction(func(tx *gorm.DB) error {
		wallets, err := r.GetAccounts(ctx, []int64{sourceAccountID, destinationAccountID})
		if err != nil {
			return err
		}

		if len(wallets) < 2 {
			return util.ErrAccountIDNotExist
		}

		var sourceAccount, destinationAccount *model.Wallet
		for _, wallet := range wallets {
			if wallet.ID == sourceAccountID {
				sourceAccount = wallet
			} else {
				destinationAccount = wallet
			}
		}

		if sourceAccount.Currency != destinationAccount.Currency {
			return util.ErrCurrencyTypeMismatch
		}

		amountIntValue, err := util.ConvertMoneyToIntegerMoney(amount, sourceAccount.Currency)
		if err != nil {
			return err
		}

		if sourceAccount.Balance < amountIntValue {
			return util.ErrBalanceNotEnough
		}

		sourceAccount.Balance -= amountIntValue
		if err := r.updateAccount(ctx, sourceAccount.ID, sourceAccount.Balance, tx); err != nil {
			log.Printf("update account error account id %d balance %d, err %s\n", sourceAccount.ID, sourceAccount.Balance, err.Error())
			return err
		}

		destinationAccount.Balance += amountIntValue
		if err := r.updateAccount(ctx, destinationAccount.ID, destinationAccount.Balance, tx); err != nil {
			log.Printf("update account error account id %d balance %d, err %s\n", destinationAccount.ID, destinationAccount.Balance, err.Error())
			return err
		}

		if err := r.insertTranslation(ctx, sourceAccountID, destinationAccountID, amountIntValue, tx); err != nil {
			log.Printf("insert translation error account id %d %d balance %f\n", sourceAccountID, destinationAccountID, amount)
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

func (r *WalletRepoImpl) updateAccount(ctx context.Context, accountID, balance int64, tx *gorm.DB) error {
	tx = tx.Table(model.WalletTableName).Where("id = ?", accountID).Updates(map[string]interface{}{
		"balance": balance,
	})

	return tx.Error
}

func (r *WalletRepoImpl) insertTranslation(ctx context.Context, sourceAccountID, destinationAccountID, amount int64, tx *gorm.DB) error {
	transaction := model.Transaction{
		ID:                   util.GetUniqueID(),
		SourceAccountID:      sourceAccountID,
		DestinationAccountID: destinationAccountID,
		Amount:               amount,
	}

	return tx.Table(model.TransactionTableName).Create(transaction).Error
}
