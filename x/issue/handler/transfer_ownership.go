package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgTransferOwnership(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransferOwnership) (*sdk.Result, error) {
	// Sub fee from sender
	fee := k.GetParams(ctx).TransferOwnerFee
	if err := k.ChargeFee(ctx, msg.Owner, fee); err != nil {
		return nil, err
	}

	if err := k.TransferOwnership(ctx, msg.Owner, msg.ToAddress, msg.Denom); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
