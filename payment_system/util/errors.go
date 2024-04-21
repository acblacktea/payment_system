package util

import "errors"

// sentinel err
var (
	ErrDuplicateAccountID   = errors.New("account id existed")
	ErrBalanceNotEnough     = errors.New("user doesn't have enough balance")
	ErrAccountIDNotExist    = errors.New("account is valid, not exist")
	ErrInvalidCurrency      = errors.New("invalid currency")
	ErrCurrencyTypeMismatch = errors.New("currency type mismatch")
	ErrInvalidMoneyValue    = errors.New("invalid money value")
	ErrInvalidRequest       = errors.New("invalid request, please check")
)

// status code, should refinement
var (
	InternalErrCode int32 = 500
	SuccessCode     int32 = 200
	InvalidArgument int32 = 400
)

const (
	GormExistItem = "Duplicate entry"
)
