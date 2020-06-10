package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/supply/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
}

type CoinKeeper interface {
	GetCoins(sdk.Context, sdk.AccAddress) sdk.Coins
	AddCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, error)
	SubtractCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, error)
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetSendEnabled(sdk.Context) bool
	BlacklistedAddr(sdk.AccAddress) bool
}

type SupplyKeeper interface {
	GetSupply(sdk.Context) exported.SupplyI
	SetSupply(sdk.Context, exported.SupplyI)
	MintCoins(sdk.Context, string, sdk.Coins) error
	BurnCoins(sdk.Context, string, sdk.Coins) error
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) error
	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) error
}
