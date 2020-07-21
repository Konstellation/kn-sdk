package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func Issues(ctx sdk.Context, k keeper.Keeper, data []byte) ([]byte, error) {
	var params types.IssuesParams
	if err := k.GetCodec().UnmarshalJSON(data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	issues := k.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(k.GetCodec(), issues)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
