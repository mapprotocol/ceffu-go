package types

type WithdrawalRequest struct {
	Amount             string `json:"amount"`                       // withdrawal amount
	CoinSymbol         string `json:"coinSymbol"`                   // coin symbol
	Memo               string `json:"memo,omitempty"`               // memo/address tag
	Network            string `json:"network"`                      // network symbol
	WalletID           int64  `json:"walletId"`                     // wallet id
	WithdrawalAddress  string `json:"withdrawalAddress"`            // withdrawal address or to wallet id str  must have one
	ToWalletIDStr      string `json:"toWalletIdStr"`                // to wallet id str  or withdrawal address must have one
	CustomizeFeeAmount string `json:"customizeFeeAmount,omitempty"` // User-specified fee  , now support eth
	RequestID          int64  `json:"requestId"`                    // Unique Identifier
	Timestamp          int64  `json:"timestamp"`                    // Current Timestamp in millisecond
}

type WithdrawalDetailRequest struct {
	OrderViewID string `json:"orderViewId,omitempty"` // Withdrawal Transaction Id
	RequestID   string `json:"requestId,omitempty"`   // Client request identifier: Universal Unique identifier provided by the client side.
	Timestamp   int64  `json:"timestamp"`             // Current Timestamp in millisecond
}

type TransferWithExchangeRequest struct {
	Amount         string `json:"amount"`               // Transfer Amount
	CoinSymbol     string `json:"coinSymbol,omitempty"` // Coin symbol
	Direction      int64  `json:"direction,omitempty"`  // Transfer direction,; 10: custody->exchange; 20: exchange->custody
	ExchangeCode   int64  `json:"exchangeCode"`         // Exchange code, 10: binance
	ExchangeUserID string `json:"exchangeUserId"`       // Binance UID
	ParentWalletID int64  `json:"parentWalletId"`       // Parent Wallet Id; (Only applicable to Parent Shared Wallet)
	Status         int64  `json:"status,omitempty"`     // Status
	RequestID      int64  `json:"requestId"`            // Unique Identifier
	Timestamp      int64  `json:"timestamp"`            // Current Timestamp in millisecond
}

type TransferDetailWithExchangeRequest struct {
	OrderViewID string `json:"orderViewId,omitempty"` // Transfer transaction ID
	RequestID   string `json:"requestId,omitempty"`   // Client request identifier: Universal Unique identifier provided by the client side.
	Timestamp   int64  `json:"timestamp"`             // Current timestamp in millisecond
	WalletID    int64  `json:"walletId"`              // Wallet ID
}

// response struct

type CreatePrimeWalletRequestResponse struct {
	Data struct {
		WalletId    int64  `json:"walletId"`
		WalletIdStr string `json:"walletIdStr"`
		WalletName  string `json:"walletName"`
		WalletType  int    `json:"walletType"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type WithdrawalResponseData struct {
	OrderViewId  string `json:"orderViewId"`
	Status       int    `json:"status"`
	TransferType int    `json:"transferType"`
}

type WithdrawalResponse struct {
	Data    WithdrawalResponseData `json:"data"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
}

type WithdrawalDetailResponse struct {
	Code    string       `json:"code"`    // response code, '000000' when successed, others represent there some error occured
	Data    *Transaction `json:"data"`    // response data, maybe null
	Message string       `json:"message"` // detail of response, when code != '000000', it's detail of error
}

type Transaction struct {
	OrderViewID  string  `json:"orderViewId"`
	TxID         string  `json:"txId"` // transaction id (Only Applicable to on-chain transfer)
	TransferType int64   `json:"transferType"`
	Direction    int64   `json:"direction"`
	FromAddress  string  `json:"fromAddress"`
	ToAddress    string  `json:"toAddress"`
	Network      string  `json:"network"`
	CoinSymbol   string  `json:"coinSymbol"`
	Amount       string  `json:"amount"`
	FeeSymbol    string  `json:"feeSymbol"`
	FeeAmount    string  `json:"feeAmount"`
	Status       int64   `json:"status"`
	Memo         *string `json:"memo"`
	TxTime       string  `json:"txTime"`
	WalletStr    string  `json:"walletStr"`
	RequestID    *string `json:"requestId"` // universal unique identifier provided by the client side.
}

type TransferWithExchangeResponse struct {
	Data    *Transfer `json:"data"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
}

type TransferDetail struct {
	Amount         string `json:"amount"`
	CoinSymbol     string `json:"coinSymbol"`
	Direction      int32  `json:"direction"`
	ExchangeCode   int32  `json:"exchangeCode"`
	ExchangeUserID string `json:"exchangeUserId"`
	OrderViewID    string `json:"orderViewId"`
	Status         int32  `json:"status"`
	WalletID       int64  `json:"walletId"`
	CreateTime     int64  `json:"createTime"` // TODO field is exist in response, need to check
	RequestId      string `json:"requestId"`  // TODO field is exist in response, need to check
}

type TransferDetailWithExchangeResponse struct {
	Data    *TransferDetail `json:"data"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
}
