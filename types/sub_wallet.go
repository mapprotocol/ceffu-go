package types

type CreatSubWalletRequest struct {
	ParentWalletID string `json:"parentWalletId"`           // parent wallet id
	WalletName     string `json:"walletName,omitempty"`     // Sub Wallet name (Max 20 characters)
	AutoCollection int64  `json:"autoCollection,omitempty"` // Enable auto sweeping to parent wallet; ; 0: Not enable (Default Value), Suitable for API user who required Custody to maintain; asset ledger of each subaccount; ; 1: Enable, Suitable for API user who will maintain asset ledger of each subaccount at; their end.
	RequestID      int64  `json:"requestId"`                // Request identity
	Timestamp      int64  `json:"timestamp"`                // Current Timestamp
}

type GetDepositAddressRequest struct {
	CoinSymbol string `json:"coinSymbol"` // Coin Symbol (in capital letters); Required for Prime wallet; Not required for Qualified; wallet
	Network    string `json:"network"`    // Network symbol
	Timestamp  int64  `json:"timestamp"`  // Current Timestamp in millisecond
	WalletID   int64  `json:"walletId"`   // Sub Wallet id
}

type GetDepositHistoryRequest struct {
	WalletID   int64  `json:"walletId"`             // Prime wallet id or sub wallet id
	CoinSymbol string `json:"coinSymbol,omitempty"` // Coin symbol (in capital letters); All symbols if not specific
	Network    string `json:"network,omitempty"`    // Network symbol; All networks if not specific
	StartTime  int64  `json:"startTime"`            // Start time(timestamp in milliseconds)
	EndTime    int64  `json:"endTime"`              // End time(timestamp in milliseconds)
	PageLimit  int64  `json:"pageLimit"`            // Page limit
	PageNo     int64  `json:"pageNo"`               // Page no
	Timestamp  int64  `json:"timestamp"`            // Current Timestamp in millisecond
}

type TransferRequest struct {
	CoinSymbol   string  `json:"coinSymbol"`   // Coin symbol
	Amount       float64 `json:"amount"`       // Transfer amount
	FromWalletID int64   `json:"fromWalletId"` // From wallet ID
	ToWalletID   int64   `json:"toWalletId"`   // To wallet ID
	RequestID    int64   `json:"requestId"`    // Client request identifier, Client provided Unique Identifier. (Max 70 characters)
	Timestamp    int64   `json:"timestamp"`    // Current timestamp in millisecond
}

// response struct

type CreatSubWalletResponse struct {
	Data struct {
		WalletId          int64  `json:"walletId"`
		WalletIdStr       string `json:"walletIdStr"`
		WalletName        string `json:"walletName"`
		WalletType        uint32 `json:"walletType"`
		ParentWalletId    int64  `json:"parentWalletId"`
		ParentWalletIdStr string `json:"parentWalletIdStr"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetDepositAddressResponse struct {
	Data struct {
		WalletAddress string `json:"walletAddress"`
		Memo          string `json:"memo"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetDepositHistoryResponse struct {
	Data struct {
		Data      []*Transaction `json:"data"`
		TotalPage int            `json:"totalPage"`
		PageNo    int            `json:"pageNo"`
		PageLimit int            `json:"pageLimit"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Transfer struct {
	OrderViewId string `json:"orderViewId"` // Transfer transaction Id
	Status      int32  `json:"status"`      // Status: 10: Pending, 20: Processing, 30: Send success, 99: Failed
	Direction   int32  `json:"direction"`   // Transfer direction: 10: prime wallet->sub wallet, 20: sub wallet->prime wallet, 30: sub wallet-> sub wallet, 40: prime wallet → prime wallet
}

type TransferResponse struct {
	Data    *Transfer `json:"data"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
}
