// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentsystem

import (
	"context"
	"errors"
	payment_system "github.com/acblacktea/payment_system/payment_system/kitex_gen/acblacktea/payment_system/payment_system"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateAccount": kitex.NewMethodInfo(
		createAccountHandler,
		newPaymentSystemCreateAccountArgs,
		newPaymentSystemCreateAccountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAccount": kitex.NewMethodInfo(
		getAccountHandler,
		newPaymentSystemGetAccountArgs,
		newPaymentSystemGetAccountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SubmitTransaction": kitex.NewMethodInfo(
		submitTransactionHandler,
		newPaymentSystemSubmitTransactionArgs,
		newPaymentSystemSubmitTransactionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	paymentSystemServiceInfo                = NewServiceInfo()
	paymentSystemServiceInfoForClient       = NewServiceInfoForClient()
	paymentSystemServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return paymentSystemServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return paymentSystemServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return paymentSystemServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PaymentSystem"
	handlerType := (*payment_system.PaymentSystem)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "payment_system",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment_system.PaymentSystemCreateAccountArgs)
	realResult := result.(*payment_system.PaymentSystemCreateAccountResult)
	success, err := handler.(payment_system.PaymentSystem).CreateAccount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSystemCreateAccountArgs() interface{} {
	return payment_system.NewPaymentSystemCreateAccountArgs()
}

func newPaymentSystemCreateAccountResult() interface{} {
	return payment_system.NewPaymentSystemCreateAccountResult()
}

func getAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment_system.PaymentSystemGetAccountArgs)
	realResult := result.(*payment_system.PaymentSystemGetAccountResult)
	success, err := handler.(payment_system.PaymentSystem).GetAccount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSystemGetAccountArgs() interface{} {
	return payment_system.NewPaymentSystemGetAccountArgs()
}

func newPaymentSystemGetAccountResult() interface{} {
	return payment_system.NewPaymentSystemGetAccountResult()
}

func submitTransactionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*payment_system.PaymentSystemSubmitTransactionArgs)
	realResult := result.(*payment_system.PaymentSystemSubmitTransactionResult)
	success, err := handler.(payment_system.PaymentSystem).SubmitTransaction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPaymentSystemSubmitTransactionArgs() interface{} {
	return payment_system.NewPaymentSystemSubmitTransactionArgs()
}

func newPaymentSystemSubmitTransactionResult() interface{} {
	return payment_system.NewPaymentSystemSubmitTransactionResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateAccount(ctx context.Context, req *payment_system.CreateAccountRequest) (r *payment_system.CreateAccountResponse, err error) {
	var _args payment_system.PaymentSystemCreateAccountArgs
	_args.Req = req
	var _result payment_system.PaymentSystemCreateAccountResult
	if err = p.c.Call(ctx, "CreateAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAccount(ctx context.Context, req *payment_system.GetAccountRequest) (r *payment_system.GetAccountResponse, err error) {
	var _args payment_system.PaymentSystemGetAccountArgs
	_args.Req = req
	var _result payment_system.PaymentSystemGetAccountResult
	if err = p.c.Call(ctx, "GetAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SubmitTransaction(ctx context.Context, req *payment_system.SubimitTransactionRequest) (r *payment_system.SubimitTransactionResponse, err error) {
	var _args payment_system.PaymentSystemSubmitTransactionArgs
	_args.Req = req
	var _result payment_system.PaymentSystemSubmitTransactionResult
	if err = p.c.Call(ctx, "SubmitTransaction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
