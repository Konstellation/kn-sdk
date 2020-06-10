package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgDecreaseAllowance(ctx sdk.Context, k keeper.Keeper, msg types.MsgDecreaseAllowance) (*sdk.Result, error) {
	if err := k.DecreaseAllowance(ctx, msg.Owner, msg.Spender, msg.Amount); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
