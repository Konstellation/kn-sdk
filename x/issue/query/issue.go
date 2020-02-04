package query

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Issue(ctx sdk.Context, k keeper.Keeper, issueId string) ([]byte, sdk.Error) {
	fmt.Println(issueId)
	return nil, nil
}
