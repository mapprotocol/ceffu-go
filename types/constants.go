package types

const (
	AutoCollectionDisabled = 0
	AutoCollectionEnabled  = 1
)

const (
	TransactionDirectionDeposit    = 10
	TransactionDirectionWithdrawal = 20
)

const (
	TransferDirectionParentWalletToSubWallet = 10
	TransferDirectionSubWalletToParentWallet = 20
	TransferDirectionSubWalletToSubWallet    = 30
)

const (
	TransactionStatusPending    = 10
	TransactionStatusProcessing = 20
	TransactionStatusSuccess    = 30
	TransactionStatusConfirmed  = 40
	TransactionStatusFailed     = 99
)

const (
	TransferTypeOnChain  = 10
	TransferTypeInternal = 20
)

func ToAutoCollection(autoCollection bool) int64 {
	if autoCollection {
		return AutoCollectionEnabled
	}
	return AutoCollectionDisabled
}
