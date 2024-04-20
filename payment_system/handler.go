package main

import (
	"context"

	payment_system "github.com/acblacktea/payment_system/payment_system/kitex_gen/acblacktea/payment_system/payment_system"
	payment_system_service "github.com/acblacktea/payment_system/payment_system/service/payment_system"
	"github.com/acblacktea/payment_system/payment_system/util"
)

// PaymentSystemImpl implements the last service interface defined in the IDL.
type PaymentSystemImpl struct {
	PaymentSystemService payment_system_service.PaymentSystemService
}

// CreateAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) CreateAccount(ctx context.Context, req *payment_system.CreateAccountRequest) (resp *payment_system.CreateAccountResponse, err error) {
	if req.Account == nil {
		return &payment_system.CreateAccountResponse{
			BaseResp: &payment_system.BaseResp{
				StatusMessage: util.ErrInvalidRequest.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, util.ErrInvalidRequest
	}

	if err := s.PaymentSystemService.CreateAccount(ctx, req.Account.AccountID, req.Account.Balance, req.Account.Currency); err != nil {
		return &payment_system.CreateAccountResponse{
			BaseResp: &payment_system.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, err
	}

	return &payment_system.CreateAccountResponse{
		BaseResp: &payment_system.BaseResp{
			StatusCode: util.SuccessCode,
		},
	}, nil
}

// GetAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) GetAccount(ctx context.Context, req *payment_system.GetAccountRequest) (resp *payment_system.GetAccountResponse, err error) {
	account, err := s.PaymentSystemService.GetAccount(ctx, req.AccountID)
	if err != nil {
		return &payment_system.GetAccountResponse{
			BaseResp: &payment_system.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, err
	}

	return &payment_system.GetAccountResponse{
		BaseResp: &payment_system.BaseResp{
			StatusCode: util.SuccessCode,
		},
		Account: &payment_system.Account{
			AccountID: account.AccountID,
			Balance:   account.Balance,
			Currency:  account.Currency,
		},
	}, nil
}

// SubmitTransaction implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) SubmitTransaction(ctx context.Context, req *payment_system.SubimitTransactionRequest) (resp *payment_system.SubimitTransactionResponse, err error) {
	if err := s.PaymentSystemService.SubmitTransaction(ctx, req.SourceAccountID, req.DestinationAccountID, req.Amount); err != nil {
		return &payment_system.SubimitTransactionResponse{
			BaseResp: &payment_system.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, err
	}

	return &payment_system.SubimitTransactionResponse{
		BaseResp: &payment_system.BaseResp{
			StatusCode: util.SuccessCode,
		},
	}, nil
}
