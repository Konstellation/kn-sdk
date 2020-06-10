package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgIssueCreate(ctx sdk.Context, k keeper.Keeper, msg types.MsgIssueCreate) (*sdk.Result, error) {
	// Sub fee from issuer
	fee := k.GetParams(ctx).IssueFee
	if err := k.ChargeFee(ctx, msg.Issuer, fee); err != nil {
		return nil, err
	}

	params, errr := types.NewIssueParams(msg.IssueParams)
	if errr != nil {
		return nil, types.ErrInvalidIssueParams
	}

	ci := k.CreateIssue(ctx, msg.Owner, msg.Issuer, params)
	if err := k.Issue(ctx, ci); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
