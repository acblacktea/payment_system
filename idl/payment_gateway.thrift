include "base.thrift"
namespace go acblacktea.payment_system.payment_gateway
struct Account {
    1: required i64 account_id
    2: required double balance
    3: required string currency
}

struct CreateAccountRequest {
     1: required i64 account_id (api.form="account_id");
     2: required double balance (api.form ="balance");
     3: required string currency (api.form ="currency");
}

struct CreateAccountResponse {
    255: required base.BaseResp BaseResp,
}

struct GetAccountRequest {
    1: required i64 account_id (api.query="account_id")
}

struct GetAccountResponse {
    1: optional Account account,

    255: required base.BaseResp BaseResp,
}

struct SubimitTransactionRequest {
    1: required i64 source_account_id (api.form="source_account_id");
    2: required i64 destination_account_id (api.form="destination_account_id");
    3: required double amount (api.form="amount");
}

struct SubimitTransactionResponse {
    255: required base.BaseResp BaseResp,
}

service PaymentGateway {
    CreateAccountResponse CreateAccount(1: CreateAccountRequest req) (api.post="accounts");
    GetAccountResponse GetAccount(1: GetAccountRequest req) (api.get="accounts");
    SubimitTransactionResponse SubmitTransaction(1: SubimitTransactionRequest req) (api.Post="transactions");
}