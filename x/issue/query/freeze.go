package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Freeze(ctx sdk.Context, k keeper.Keeper, denom string, holder string) ([]byte, error) {
	holderAddress, _ := sdk.AccAddressFromBech32(holder)
	freeze := k.GetFreeze(ctx, denom, holderAddress)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freeze)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
