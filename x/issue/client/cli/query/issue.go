package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/kn-sdk/x/issue/types"
	"github.com/spf13/cobra"
)

func pathQueryIssue(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, denom)
}

func getIssue(cliCtx context.CLIContext, denom string) ([]byte, int64, error) {
	return cliCtx.QueryWithData(pathQueryIssue(denom), nil)
}

// getCmdQueryIssues implements the query issue command.
func getQueryCmdIssue(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Query issue by denom",
		Long:  "Query issue by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denom := args[0]

			// Query the issues
			res, _, err := getIssue(cliCtx, denom)
			if err != nil {
				return err
			}

			var issue types.CoinIssue
			cdc.MustUnmarshalJSON(res, &issue)
			return cliCtx.PrintOutput(&issue)
		},
	}

	cmd.Flags().String(flagAddress, "", "Token owner address")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}
