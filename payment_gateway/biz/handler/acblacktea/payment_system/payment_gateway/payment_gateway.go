// Code generated by hertz generator.

package payment_gateway

import (
	"context"
	"fmt"

	payment_gateway "github.com/acblacktea/payment_system/payment_gateway/biz/model/acblacktea/payment_system/payment_gateway"
	"github.com/acblacktea/payment_system/payment_gateway/biz/model/base"
	"github.com/acblacktea/payment_system/payment_gateway/kitex_gen/acblacktea/payment_system/payment_system"
	"github.com/acblacktea/payment_system/payment_gateway/kitex_gen/acblacktea/payment_system/payment_system/paymentsystem"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	client2 "github.com/cloudwego/kitex/client"
)

var client paymentsystem.Client

func InitClient() {
	newClient, err := paymentsystem.NewClient("acblacktea.payment_system.payment_system", client2.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic(err)
	}

	client = newClient
}

// CreateAccount .
// @router accounts [POST]
func CreateAccount(ctx context.Context, c *app.RequestContext) { //ignore_security_alert IDOR
	var err error
	var req payment_gateway.CreateAccountRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	rpcResp, err := client.CreateAccount(ctx, &payment_system.CreateAccountRequest{
		Account: &payment_system.Account{
			AccountID: req.AccountID,
			Balance:   req.Balance,
			Currency:  req.Currency,
		},
	})

	resp := &payment_gateway.CreateAccountResponse{}
	stats := consts.StatusOK
	if err != nil {
		stats = consts.StatusInternalServerError
	}

	if rpcResp != nil && rpcResp.BaseResp != nil {
		resp.BaseResp = &base.BaseResp{
			StatusMessage: rpcResp.BaseResp.StatusMessage,
			StatusCode:    rpcResp.BaseResp.StatusCode,
		}

		if rpcResp.BaseResp.StatusCode != 200 {
			stats = consts.StatusInternalServerError
		}
	}

	c.JSON(stats, resp)
}

// GetAccount .
// @router accounts [GET]
func GetAccount(ctx context.Context, c *app.RequestContext) { //ignore_security_alert IDOR
	var err error
	var req payment_gateway.GetAccountRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("%d\n", req.AccountID)

	rpcResp, err := client.GetAccount(ctx, &payment_system.GetAccountRequest{
		AccountID: req.AccountID,
	})

	resp := &payment_gateway.GetAccountResponse{}
	stats := consts.StatusOK
	if err != nil {
		stats = consts.StatusInternalServerError
	}

	fmt.Printf("%v", resp)
	fmt.Printf("%v", rpcResp.BaseResp)
	if rpcResp != nil && rpcResp.BaseResp != nil {
		resp.BaseResp = &base.BaseResp{
			StatusMessage: rpcResp.BaseResp.StatusMessage,
			StatusCode:    rpcResp.BaseResp.StatusCode,
		}

		if rpcResp.BaseResp.StatusCode != 200 {
			stats = consts.StatusInternalServerError
		}
	}

	if rpcResp != nil && rpcResp.Account != nil {
		resp.Account = &payment_gateway.Account{
			AccountID: rpcResp.Account.AccountID,
			Balance:   rpcResp.Account.Balance,
			Currency:  rpcResp.Account.Currency,
		}
	}

	c.JSON(stats, resp)
}

// SubmitTransaction .
// @router transactions [POST]
func SubmitTransaction(ctx context.Context, c *app.RequestContext) { //ignore_security_alert IDOR
	var err error
	var req payment_gateway.SubimitTransactionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("%f\n", req.Amount)
	rpcResp, err := client.SubmitTransaction(ctx, &payment_system.SubimitTransactionRequest{
		DestinationAccountID: req.DestinationAccountID,
		SourceAccountID:      req.SourceAccountID,
		Amount:               req.Amount,
	})

	resp := &payment_gateway.SubimitTransactionResponse{}
	stats := consts.StatusOK
	if err != nil {
		stats = consts.StatusInternalServerError
	}

	if rpcResp != nil && rpcResp.BaseResp != nil {
		resp.BaseResp = &base.BaseResp{
			StatusMessage: rpcResp.BaseResp.StatusMessage,
			StatusCode:    rpcResp.BaseResp.StatusCode,
		}

		if rpcResp.BaseResp.StatusCode != 200 {
			stats = consts.StatusInternalServerError
		}
	}

	c.JSON(stats, resp)
}
