package repo

import (
	"context"
	"log"

	"github.com/acblacktea/payment_system/payment_system/dal/model"
	"github.com/acblacktea/payment_system/payment_system/util"
	"gorm.io/gorm"
)

type WalletRepo interface {
	CreateAccount(ctx context.Context, accountID, initBalance int64, currency string) error
	GetAccounts(ctx context.Context, accountIDs []int64) ([]*model.Wallet, error)
	SubmitTransaction(ctx context.Context, params SubmitTransactionParameters) error
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

	return r.dbClient.Table(model.WalletTableName).Create(user).Error
}

func (r *WalletRepoImpl) GetAccounts(ctx context.Context, accountIDs []int64) ([]*model.Wallet, error) {
	wallets := make([]*model.Wallet, len(accountIDs))
	result := r.dbClient.Table(model.WalletTableName).Where("id in ", accountIDs).Find(&wallets)
	return wallets, result.Error
}

type SubmitTransactionParameters struct {
	SourceAccountID      int64
	DestinationAccountID int64
	SourceBalance        int64
	DestinationBalance   int64
	Amount               int64
	Currency             string
}

func (r *WalletRepoImpl) SubmitTransaction(ctx context.Context, params SubmitTransactionParameters) error {
	return r.dbClient.Transaction(func(tx *gorm.DB) error {
		if err := r.updateAccount(ctx, params.SourceAccountID, params.SourceBalance, tx); err != nil {
			log.Printf("update account error account id %d balance %d, err %s\n", params.SourceAccountID, params.SourceBalance, err.Error())
			return err
		}

		if err := r.updateAccount(ctx, params.DestinationAccountID, params.DestinationBalance, tx); err != nil {
			log.Printf("update account error account id %d balance %d, err %s\n", params.DestinationAccountID, params.DestinationBalance, err.Error())
			return err
		}

		if err := r.insertTranslation(ctx, params.SourceAccountID, params.DestinationAccountID, params.Amount, tx); err != nil {
			log.Printf("update account error account id %d balance %d\n", params.DestinationAccountID, params.DestinationBalance)
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

func (r *WalletRepoImpl) updateAccount(ctx context.Context, accountID, balance int64, tx *gorm.DB) error {
	account := model.Wallet{
		ID:      accountID,
		Balance: balance,
	}

	tx = tx.Model(&account).Where("id = ?", accountID).Updates(map[string]interface{}{
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
