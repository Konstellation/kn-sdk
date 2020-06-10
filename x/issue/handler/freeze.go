package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgFreeze(ctx sdk.Context, k keeper.Keeper, msg types.MsgFreeze) (*sdk.Result, error) {
	// Sub fee from sender
	fee := k.GetParams(ctx).FreezeFee
	if err := k.ChargeFee(ctx, msg.Freezer, fee); err != nil {
		return nil, err
	}

	if err := k.Freeze(ctx, msg.Freezer, msg.Holder, msg.Denom, msg.Op); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
