package query

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

func pathQueryIssueAllowances(owner sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowances, denom, owner.String())
}

func getIssueAllowances(cliCtx context.CLIContext, owner sdk.AccAddress, denom string) ([]byte, int64, error) {
	return cliCtx.QueryWithData(pathQueryIssueAllowances(owner, denom), nil)
}

// getQueryCmdAllowances implements the query issue command.
func getQueryCmdAllowances(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowances [owner] [denom]",
		Args:  cobra.ExactArgs(2),
		Short: "Query allowances",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			denom := args[1]

			res, _, err := getIssueAllowances(cliCtx, owner, denom)
			if err != nil {
				return err
			}

			var allowances types.Allowances
			cdc.MustUnmarshalJSON(res, &allowances)

			return cliCtx.PrintOutput(&allowances)
		},
	}

	return cmd
}
