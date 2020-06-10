package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgTransfer(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransfer) (*sdk.Result, error) {
	if err := k.Transfer(ctx, msg.FromAddress, msg.ToAddress, msg.Amount); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
