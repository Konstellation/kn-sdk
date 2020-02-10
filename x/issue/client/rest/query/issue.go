package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

func pathQueryIssue(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, denom)
}

// HTTP request handler to query specified issues
func issueHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		res, height, err := cliCtx.QueryWithData(pathQueryIssue(denom), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
