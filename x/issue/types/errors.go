// nolint
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultCodespace                   = "issue"
	CodeUnknownIssue            uint32 = 1
	CodeIssueExists             uint32 = 2
	CodeOwnerMismatch           uint32 = 3
	CodeInvalidOwnerAddress     uint32 = 4
	CodeInvalidDenom            uint32 = 5
	CodeInvalidAmount           uint32 = 6
	CodeInvalidCoinDecimals     uint32 = 7
	CodeInvalidTotalSupply      uint32 = 8
	CodeInvalidDescription      uint32 = 9
	CodeInvalidSymbol           uint32 = 10
	CodeInvalidFreezeOp         uint32 = 11
	CodeInvalidFeature          uint32 = 12
	CodeInvalidInput            uint32 = 13
	CodeInvalidIssueFee         uint32 = 14
	CodeInvalidMintFee          uint32 = 15
	CodeInvalidBurnFee          uint32 = 16
	CodeInvalidBurnFromFee      uint32 = 17
	CodeInvalidFreezeFee        uint32 = 18
	CodeInvalidUnFreezeFee      uint32 = 19
	CodeInvalidTransferOwnerFee uint32 = 20
	CodeInvalidSenderAddress    uint32 = 21
	CodeInvalidRecipientAddress uint32 = 22
	CodeInvalidFromAddress      uint32 = 23
	CodeInvalidSpenderAddress   uint32 = 24
	CodeInvalidMinterAddress    uint32 = 25
	CodeInvalidBurnerAddress    uint32 = 26
	CodeInvalidFreezerAddress   uint32 = 27
	CodeInvalidHolderAddress    uint32 = 28

	CodeAmountLowerAllowance uint32 = 21
	CodeNotEnoughFee         uint32 = 22
	CodeCannotMint           uint32 = 23
	CodeCannotBurnOwner      uint32 = 24
	CodeCannotBurnHolder     uint32 = 25
	CodeCannotBurnFrom       uint32 = 26
	CodeCannotFreeze         uint32 = 27
	CodeNotTransferOut       uint32 = 28
	CodeNotTransferIn        uint32 = 29
)

var (
	// ErrInvalidIssueParams to doc
	ErrInvalidIssueParams = sdkerrors.Register(DefaultCodespace, CodeInvalidInput, "invalid issue params")

	// ErrIssueAlreadyExists to doc
	ErrIssueAlreadyExists = sdkerrors.Register(DefaultCodespace, CodeIssueExists, "issue already exists")

	// ErrUnknownIssue to doc
	ErrUnknownIssue = sdkerrors.Register(DefaultCodespace, CodeUnknownIssue, "invalid issue denom")

	// ErrInvalidOwnerAddress to doc
	ErrInvalidOwnerAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidOwnerAddress, "invalid owner address")

	// ErrInvalidSenderAddress to doc
	ErrInvalidSenderAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidSenderAddress, "invalid sender address")

	// ErrInvalidRecipientAddress to doc
	ErrInvalidRecipientAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidRecipientAddress, "invalid recipient address")

	// ErrInvalidFromAddress to doc
	ErrInvalidFromAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidFromAddress, "invalid from address")

	// ErrInvalidSpenderAddress to doc
	ErrInvalidSpenderAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidSpenderAddress, "invalid spender address")

	// ErrInvalidMinterAddress to doc
	ErrInvalidMinterAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidMinterAddress, "invalid minter address")

	// ErrInvalidBurnerAddress to doc
	ErrInvalidBurnerAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidBurnerAddress, "invalid burner address")

	// ErrInvalidFreezerAddress to doc
	ErrInvalidFreezerAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidFreezerAddress, "invalid freezer address")

	// ErrInvalidHolderAddress to doc
	ErrInvalidHolderAddress = sdkerrors.Register(DefaultCodespace, CodeInvalidHolderAddress, "invalid holder address")

	// ErrOwnerMismatch to doc
	ErrOwnerMismatch = sdkerrors.Register(DefaultCodespace, CodeOwnerMismatch, "owner mismatch with token")

	// ErrNotEnoughFee to doc
	ErrNotEnoughFee = sdkerrors.Register(DefaultCodespace, CodeNotEnoughFee, "not enough fee")

	// ErrCoinDecimalsMaxValueNotValid to doc
	ErrCoinDecimalsMaxValueNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimals, fmt.Sprintf("decimals max value is %d", CoinDecimalsMaxValue))

	// ErrCoinDecimalsMultipleNotValid to doc
	ErrCoinDecimalsMultipleNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimals, fmt.Sprintf("decimals must be a multiple of %d", CoinDecimalsMultiple))

	// ErrCoinTotalSupplyMaxValueNotValid to doc
	ErrCoinTotalSupplyMaxValueNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidTotalSupply, fmt.Sprintf("total supply max value is %s", CoinMaxTotalSupply.String()))

	// ErrCoinDescriptionNotValid to doc
	ErrCoinDescriptionNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidDescription, "description is not valid json")

	// ErrCoinDescriptionMaxLengthNotValid to doc
	ErrCoinDescriptionMaxLengthNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidDescription, fmt.Sprintf("Description max length is %d", CoinDescriptionMaxLength))

	// ErrInvalidSymbol to doc
	ErrInvalidSymbol = sdkerrors.Register(DefaultCodespace, CodeInvalidSymbol, "invalid symbol")

	// ErrCannotMint to doc
	ErrCannotMint = sdkerrors.Register(DefaultCodespace, CodeCannotMint, "cannot mint")

	// ErrCannotBurnOwner to doc
	ErrCannotBurnOwner = sdkerrors.Register(DefaultCodespace, CodeCannotBurnOwner, "cannot burn with owner")

	// ErrCannotBurnHolder to doc
	ErrCannotBurnHolder = sdkerrors.Register(DefaultCodespace, CodeCannotBurnHolder, "cannot burn with holder")

	// ErrCannotBurnFrom to doc
	ErrCannotBurnFrom = sdkerrors.Register(DefaultCodespace, CodeCannotBurnFrom, "cannot burn from")

	// ErrCannotFreeze to doc
	ErrCannotFreeze = sdkerrors.Register(DefaultCodespace, CodeCannotFreeze, "cannot freeze")
)

// ErrCoinDecimalsMultipleNotValid to doc
func ErrAmountGreaterThanAllowance(amt sdk.Coin, allowance sdk.Coin) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeAmountLowerAllowance, fmt.Sprintf("amount greater than allowance %s > %s", amt.String(), allowance.String()))
}

// ErrInvalidAmount to doc
func ErrInvalidAmount(amount string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidAmount, fmt.Sprintf("%s is not a valid amount", amount))
}

// ErrInvalidTotalSupply to doc
func ErrInvalidTotalSupply(amount string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidTotalSupply, fmt.Sprintf("total supply invalid %s", amount))
}

// ErrCoinDecimalsMultipleNotValid to doc
func ErrInvalidDenom(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidDenom, fmt.Sprintf("denom invalid %s", denom))
}

// ErrCoinDecimalsMultipleNotValid to doc
func ErrInvalidFreezeOp(op string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFreezeOp, fmt.Sprintf("invalid freeze type %s", op))
}

// ErrInvalidFeature to doc
func ErrInvalidFeature(feature string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFeature, fmt.Sprintf("feature invalid %s", feature))
}

// ErrCannotTransferIn to doc
func ErrCannotTransferIn(accAddress string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeNotTransferIn, fmt.Sprintf("can not transfer in to %s", accAddress))
}

// ErrCannotTransferOut to doc
func ErrCannotTransferOut(accAddress string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeNotTransferOut, fmt.Sprintf("can not transfer out from %s", accAddress))
}

// ErrInvalidIssueFee to doc
func ErrInvalidIssueFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidIssueFee, fmt.Sprintf("invalid issue fee: %s", fee))
}

// ErrInvalidMintFee to doc
func ErrInvalidMintFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidMintFee, fmt.Sprintf("invalid mint fee: %s", fee))
}

// ErrInvalidBurnFee to doc
func ErrInvalidBurnFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidBurnFee, fmt.Sprintf("invalid burn fee: %s", fee))
}

// ErrInvalidBurnFromFee to doc
func ErrInvalidBurnFromFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidBurnFromFee, fmt.Sprintf("invalid burn from fee: %s", fee))
}

// ErrInvalidFreezeFee to doc
func ErrInvalidFreezeFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFreezeFee, fmt.Sprintf("invalid freeze fee: %s", fee))
}

// ErrInvalidUnfreezeFee to doc
func ErrInvalidUnfreezeFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidUnFreezeFee, fmt.Sprintf("invalid unfreeze fee: %s", fee))
}

// ErrInvalidTransferOwnerFee to doc
func ErrInvalidTransferOwnerFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidTransferOwnerFee, fmt.Sprintf("invalid transfer owner fee: %s", fee))
}
