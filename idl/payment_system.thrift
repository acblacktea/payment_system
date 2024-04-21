include "base.thrift"
namespace go acblacktea.payment_system.payment_system


struct Account {
    1: required i64 AccountID
    2: required double Balance
    3: required string Currency
}

struct CreateAccountRequest {
    1: required Account Account,
}

struct CreateAccountResponse {
    255: required base.BaseResp BaseResp,
}

struct GetAccountRequest {
    1: required i64 AccountID
}

struct GetAccountResponse {
    1: optional Account Account,

    255: required base.BaseResp BaseResp,
}

struct SubimitTransactionRequest {
    1: required i64 SourceAccountID,
    2: required i64 DestinationAccountID,
    3: required double Amount
}

struct SubimitTransactionResponse {
    255: required base.BaseResp BaseResp,
}

service PaymentSystem {
    CreateAccountResponse CreateAccount(1: CreateAccountRequest req)
    GetAccountResponse GetAccount(1: GetAccountRequest req)
    SubimitTransactionResponse SubmitTransaction(1: SubimitTransactionRequest req)
}