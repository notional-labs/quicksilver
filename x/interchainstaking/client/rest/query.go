package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ingenuity-build/quicksilver/x/interchainstaking/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/interchainstaking/parameters",
		queryParamsHandlerFn(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		"/interchainstaking/zones",
		queryZonesHandlerFn(clientCtx),
	).Methods("GET")

	r.HandleFunc(
		"/interchainstaking/zones/{chainId}/depositAddress",
		queryZoneDepositAddressHandlerFn(clientCtx),
	).Methods("GET")
}

func queryParamsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryZonesHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRegisteredZonesInfo)

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryZoneDepositAddressHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRegisteredZoneDepositAddress)
		vars := mux.Vars(r)

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		query := types.QueryDepositAccountForChainRequest{ChainId: vars["chainId"]}
		bz, err := clientCtx.LegacyAmino.MarshalJSON(&query)

		res, height, err := clientCtx.QueryWithData(route, bz)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
