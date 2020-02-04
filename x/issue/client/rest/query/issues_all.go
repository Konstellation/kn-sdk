package query

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

func pathQueryIssuesAll() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssuesAll)
}

// HTTP request handler to query all issues
func issuesAllHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, height, err := cliCtx.QueryWithData(pathQueryIssuesAll(), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
