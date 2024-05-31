package client

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/mapprotocol/ceffu-go/types"
)

type Wallet interface {
	Withdrawal(ctx context.Context, request *types.WithdrawalRequest) (*types.WithdrawalResponseData, error)
	WithdrawalDetail(ctx context.Context, orderViewID string) (*types.Transaction, error)
	TransferWithExchange(ctx context.Context, request *types.TransferWithExchangeRequest) (*types.Transfer, error)
	TransferDetailWithExchange(ctx context.Context, orderViewID string, walletID int64) (*types.TransferDetail, error)
}

// Withdrawal This method enables the withdrawal of funds from the specified wallet to an external address
// or a Ceffu wallet. The withdrawal endpoint is applicable only to parent Qualified wallet ID or Cosign wallet
// or parent Prime wallet ID. To indicate the destination address, either 'withdrawalAddress'
// or 'ToWalletIdStr' must be provided. If the destination address is a Ceffu wallet address,
// the whitelisted address verification will be bypassed.
//
// IMPORTANT NOTES: The amount field in Withdrawal (v2) endpoint means withdrawal amount excluded network fee in v2,
// that is exact amount receiver will receive. Please use Get Withdrawal History v2
// and Get Withdrawal Detail (v2) together with Withdrawal (v2).
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471332
func (c *client) Withdrawal(ctx context.Context, request *types.WithdrawalRequest) (*types.WithdrawalResponseData, error) {
	request.RequestID = c.RequestID.Generate()
	request.Timestamp = time.Now().UnixMilli()

	ret, err := c.Post(ctx, PathWithdrawal, request)
	if err != nil {
		return nil, NewRequestError(
			PathWithdrawal,
			WithError(err),
		)
	}
	response := types.WithdrawalResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathWithdrawal,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return &response.Data, nil
}

// WithdrawalDetail This method allows to get withdrawal details by orderViewId or requestId
// orderViewId or requestId shall be passed in Request Query.
// Amount field will be included fee if the fee paid in same coin symbol.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471329
func (c *client) WithdrawalDetail(ctx context.Context, orderViewID string) (*types.Transaction, error) {
	request := types.WithdrawalDetailRequest{
		OrderViewID: orderViewID,
		RequestID:   strconv.FormatInt(c.RequestID.Generate(), 10),
		Timestamp:   time.Now().UnixMilli(),
	}

	ret, err := c.Get(ctx, PathWithdrawalDetail, request)
	if err != nil {
		return nil, NewRequestError(
			PathWithdrawalDetail,
			WithError(err),
		)
	}
	response := types.WithdrawalDetailResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathWithdrawalDetail,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data, nil
}

// TransferWithExchange This method allows to transfer assets from Ceffu Prime Wallet to a bound
// Binance Account (To be bound in Web Portal [Wallets > Binance Transfer].
//
// Notes: Currently support from Ceffu to Exchange direction only.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471337
func (c *client) TransferWithExchange(ctx context.Context, request *types.TransferWithExchangeRequest) (*types.Transfer, error) {
	request.RequestID = c.RequestID.Generate()
	request.Timestamp = time.Now().UnixMilli()

	ret, err := c.Post(ctx, PathTransferWithExchange, request)
	if err != nil {
		return nil, NewRequestError(
			PathTransferWithExchange,
			WithError(err),
		)
	}
	response := types.TransferWithExchangeResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathTransferWithExchange,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data, nil
}

// TransferDetailWithExchange This method allows to get transfer details with Exchange by orderViewId or requestId
//
// orderViewId or requestId shall be passed in Request Query.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471330
func (c *client) TransferDetailWithExchange(ctx context.Context, orderViewID string, walletID int64) (*types.TransferDetail, error) {
	request := types.TransferDetailWithExchangeRequest{
		OrderViewID: orderViewID,
		WalletID:    walletID,
		RequestID:   strconv.FormatInt(c.RequestID.Generate(), 10),
		Timestamp:   time.Now().UnixMilli(),
	}

	ret, err := c.Post(ctx, PathTransferDetailWithExchange, request)
	if err != nil {
		return nil, NewRequestError(
			PathTransferDetailWithExchange,
			WithError(err),
		)
	}
	response := types.TransferDetailWithExchangeResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathTransferDetailWithExchange,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data, nil
}
