package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Issue(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, error) {
	coin, er := k.GetIssue(ctx, denom)
	if er != nil {
		return nil, er
	}

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
