include "base.thrift"
namespace go acblacktea.payment_system.payment_gateway
struct Account {
    1: required i64 AccountID
    2: required double Balance
    3: required string currency
}

struct CreateAccountRequest {
     1: required i64 AccountID (api.form="account_id");
     2: required double Balance (api.form ="balance");
     3: required string Currency (api.form ="currency");
}

struct CreateAccountResponse {
    255: required base.BaseResp BaseResp,
}

struct GetAccountRequest {
    1: required i64 AccountID (api.query="account_id")
}

struct GetAccountResponse {
    1: required Account Account,

    255: required base.BaseResp BaseResp,
}

struct SubimitTransactionRequest {
    1: required i64 SourceAccountID (api.form="source_account_id");
    2: required i64 DestinationAccountID (api.form="destination_account_id");
    3: required double Amount (api.form="amount");
}

struct SubimitTransactionResponse {
    255: required base.BaseResp BaseResp,
}

service PaymentGateway {
    CreateAccountResponse CreateAccount(1: CreateAccountRequest req) (api.post="accounts");
    GetAccountResponse GetAccount(1: GetAccountRequest req) (api.get="accounts");
    SubimitTransactionResponse SubmitTransaction(1: SubimitTransactionRequest req) (api.Post="transactions");
}