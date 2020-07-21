package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Freezes(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, error) {
	freezes := k.GetFreezes(ctx, denom)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freezes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
