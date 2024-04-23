package main

import (
	"context"
	"fmt"

	payment_system "github.com/acblacktea/payment_system/payment_system/kitex_gen/acblacktea/payment_system/payment_system"
	"github.com/acblacktea/payment_system/payment_system/kitex_gen/base"
	payment_system_service "github.com/acblacktea/payment_system/payment_system/service/payment_system"
	"github.com/acblacktea/payment_system/payment_system/util"
)

// PaymentSystemImpl implements the last service interface defined in the idl.
type PaymentSystemImpl struct {
	PaymentSystemService payment_system_service.PaymentSystemService
}

// CreateAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) CreateAccount(ctx context.Context, req *payment_system.CreateAccountRequest) (resp *payment_system.CreateAccountResponse, err error) {
	if req.Account == nil {
		return &payment_system.CreateAccountResponse{
			BaseResp: &base.BaseResp{
				StatusMessage: util.ErrInvalidRequest.Error(),
				StatusCode:    util.InvalidArgument,
			},
		}, nil
	}

	if err := s.PaymentSystemService.CreateAccount(ctx, req.Account.AccountID, req.Account.Balance, req.Account.Currency); err != nil {
		return &payment_system.CreateAccountResponse{
			BaseResp: &base.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, nil
	}

	return &payment_system.CreateAccountResponse{
		BaseResp: &base.BaseResp{
			StatusCode: util.SuccessCode,
		},
	}, nil
}

// GetAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) GetAccount(ctx context.Context, req *payment_system.GetAccountRequest) (resp *payment_system.GetAccountResponse, err error) {
	account, err := s.PaymentSystemService.GetAccount(ctx, req.AccountID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return &payment_system.GetAccountResponse{
			BaseResp: &base.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, nil
	}

	return &payment_system.GetAccountResponse{
		BaseResp: &base.BaseResp{
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
	fmt.Printf("%f\n", req.Amount)
	if err := s.PaymentSystemService.SubmitTransaction(ctx, req.SourceAccountID, req.DestinationAccountID, req.Amount); err != nil {
		return &payment_system.SubimitTransactionResponse{
			BaseResp: &base.BaseResp{
				StatusMessage: err.Error(),
				StatusCode:    util.InternalErrCode,
			},
		}, nil
	}

	return &payment_system.SubimitTransactionResponse{
		BaseResp: &base.BaseResp{
			StatusCode: util.SuccessCode,
		},
	}, nil
}
