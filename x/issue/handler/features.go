package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgFeatures(ctx sdk.Context, k keeper.Keeper, msg types.MsgFeatures) (*sdk.Result, error) {
	if err := k.ChangeFeatures(ctx, msg.Owner, msg.Denom, msg.IssueFeatures); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
