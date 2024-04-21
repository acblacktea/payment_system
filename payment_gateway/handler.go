package main

import (
	"context"
	payment_system "github.com/acblacktea/payment_system/payment_gateway/kitex_gen/acblacktea/payment_system/payment_system"
)

// PaymentSystemImpl implements the last service interface defined in the IDL.
type PaymentSystemImpl struct{}

// CreateAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) CreateAccount(ctx context.Context, req *payment_system.CreateAccountRequest) (resp *payment_system.CreateAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAccount implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) GetAccount(ctx context.Context, req *payment_system.GetAccountRequest) (resp *payment_system.GetAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// SubmitTransaction implements the PaymentSystemImpl interface.
func (s *PaymentSystemImpl) SubmitTransaction(ctx context.Context, req *payment_system.SubimitTransactionRequest) (resp *payment_system.SubimitTransactionResponse, err error) {
	// TODO: Your code here...
	return
}
