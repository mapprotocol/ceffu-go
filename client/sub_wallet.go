package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mapprotocol/ceffu-go/types"
)

type SubWallet interface {
	CreateSubWallet(ctx context.Context, parentWalletID, walletName string, autoCollection bool) (walletId int64, walletType uint32, err error)
	GetDepositAddress(ctx context.Context, network, symbol string, walletID int64) (string, error)
	GetDepositHistory(ctx context.Context, walletID int64, symbol, network string, startTime, endTime int64, pageNo, pageLimit int64) ([]*types.Transaction, error)
	Transfer(ctx context.Context, symbol string, amount float64, fromWalletID, toWalletID int64) (*types.Transfer, error)
}

// CreateSubWallet This method allows to create Sub Wallet of the requested
// Parent wallet ID (Only Applicable to Parent Wallet (Prime)).
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471342
func (c *client) CreateSubWallet(ctx context.Context, parentWalletID, walletName string, autoCollection bool) (walletId int64, walletType uint32, err error) {
	request := types.CreatSubWalletRequest{
		ParentWalletID: parentWalletID,
		WalletName:     walletName,
		AutoCollection: types.ToAutoCollection(autoCollection),
		RequestID:      c.RequestID.Generate(),
		Timestamp:      time.Now().UnixMilli(),
	}

	ret, err := c.Post(ctx, PathCreateSubWallet, request)
	if err != nil {
		return 0, 0, NewRequestError(
			PathCreateSubWallet,
			WithError(err),
		)
	}
	response := types.CreatSubWalletResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return 0, 0, err
	}
	if response.Code != SuccessCode {
		return 0, 0, NewRequestError(
			PathCreateSubWallet,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data.WalletId, response.Data.WalletType, nil
}

// GetDepositAddress This method allows to get the deposit address of the requested walletId, coinSymbol and network.
// The walletId can be parentWalletId or subWalletId.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471326
func (c *client) GetDepositAddress(ctx context.Context, network, symbol string, walletID int64) (string, error) {
	request := types.GetDepositAddressRequest{
		CoinSymbol: symbol,
		Network:    network,
		Timestamp:  time.Now().UnixMilli(),
		WalletID:   walletID,
	}

	ret, err := c.Get(ctx, PathGetDepositAddress, request)
	if err != nil {
		return "", NewRequestError(
			PathGetDepositAddress,
			WithError(err),
		)
	}
	response := types.GetDepositAddressResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return "", err
	}
	if response.Code != SuccessCode {
		return "", NewRequestError(
			PathGetDepositAddress,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data.WalletAddress, nil
}

// GetDepositHistory This method allows to get deposit history of the requested Wallet Id, coinSymbol and network.
// If PrimeWallet ID provided, returns sub wallet deposit history under the Prime Wallet.
// If SubWallet ID provided, returns specified sub wallet deposit history.
//
// Notes:
// walletId must be provided.
// Please notice the default startTime and endTime to make sure that time interval is within 0-30 days.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471585
func (c *client) GetDepositHistory(ctx context.Context, walletID int64, symbol, network string, startTime, endTime int64, pageNo, pageLimit int64) ([]*types.Transaction, error) {
	request := types.GetDepositHistoryRequest{
		WalletID:   walletID,
		CoinSymbol: symbol,
		Network:    network,
		StartTime:  startTime,
		EndTime:    endTime,
		PageLimit:  pageLimit,
		PageNo:     pageNo,
		Timestamp:  time.Now().UnixMilli(),
	}

	ret, err := c.Get(ctx, PathDepositHistory, request)
	if err != nil {
		return nil, NewRequestError(
			PathDepositHistory,
			WithError(err),
		)
	}
	response := types.GetDepositHistoryResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathDepositHistory,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}

	return response.Data.Data, nil
}

// Transfer This method allows to transfer asset between Sub Wallet and Prime Wallet Restriction:
// Only applicable to Prime wallet structure.
//
// reference: https://apidoc.ceffu.io/apidoc/shared-c9ece2c6-3ab4-4667-bb7d-c527fb3dbf78/api-3471348
func (c *client) Transfer(ctx context.Context, symbol string, amount float64, fromWalletID, toWalletID int64) (*types.Transfer, error) {
	timestamp := time.Now().UnixMilli()
	request := types.TransferRequest{
		CoinSymbol:   symbol,
		Amount:       amount,
		FromWalletID: fromWalletID,
		ToWalletID:   toWalletID,
		RequestID:    c.RequestID.Generate(),
		Timestamp:    timestamp,
	}

	ret, err := c.Post(ctx, PathTransfer, request)
	if err != nil {
		return nil, NewRequestError(
			PathTransfer,
			WithError(err),
		)
	}
	response := types.TransferResponse{}
	if err := json.Unmarshal(ret, &response); err != nil {
		return nil, err
	}
	if response.Code != SuccessCode {
		return nil, NewRequestError(
			PathTransfer,
			WithCode(response.Code),
			WithMessage(response.Message),
		)
	}
	return response.Data, nil
}
