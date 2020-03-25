package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/crypto/convert-address",
		convertAddressHandlerFn(cliCtx),
	).Methods("GET")
}

func convertAddressHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accAddress sdk.AccAddress
		address := r.URL.Query().Get(client.FlagAddress)
		if len(address) != 0 {
			ad, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			accAddress = ad
		}

		rest.PostProcessResponse(w, cliCtx, sdk.ValAddress(accAddress))
	}
}