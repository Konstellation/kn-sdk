// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/auth/ante
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/auth/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/auth/types
package issue

import (
	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	//FeeCollectorName              = types.FeeCollectorName
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)

var (
	// functions aliases
	//NewAccountKeeper                  = keeper.NewAccountKeeper
	RegisterCodec               = types.RegisterCodec
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	ValidateGenesis             = types.ValidateGenesis
	NewParams                   = types.NewParams
	ParamKeyTable               = types.ParamKeyTable
	DefaultParams               = types.DefaultParams
	GetGenesisStateFromAppState = types.GetGenesisStateFromAppState

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper = keeper.Keeper
	//BaseAccount                      = types.BaseAccount
	//NodeQuerier                      = types.NodeQuerier
	//AccountRetriever                 = types.AccountRetriever
	GenesisState = types.GenesisState
	Params       = types.Params
	//QueryAccountParams               = types.QueryAccountParams
)
