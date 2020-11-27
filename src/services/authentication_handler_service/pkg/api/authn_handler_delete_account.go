package api

import (
	"net/http"
	"strconv"
	"k8s.io/klog/v2"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

// DeleteAccountResponse is struct providing errors tied to delete account operations
type DeleteAccountResponse struct {
	Error error `json:"error"`
}

// Delete account by id request
// swagger:parameters deleteAccount
type DeleteAccountRequest struct {
	// id of the account to delete
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route DELETE /v1/account/delete/{id} deleteAccount
// Delete Account
//
// Deletes an account through the authentication service
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https, ws, wss
//
//     Security:
//       api_key:
//       oauth: read, write
// responses:
//      200: operationResponse
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 401: unAuthorizedError
// 500: internalServerError
// deletes an by account id
func (s *Server) deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	var deleteAccountResp DeleteAccountResponse
	// we extract the user id from the url initially
	// TODO: emit metrics
	authnID, err := helper.ExtractIDFromRequest(r)
	if err != nil{
		// TODO: emit metrics
		klog.Error("failed to parse account id from url", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: perform this operation in a circuit breaker, emit a metric, and trace this
	if err = s.authnClient.Client.ArchiveAccount(strconv.Itoa(int(authnID))); err != nil {
		// TODO: emit metrics
		klog.Error("failed to archive created account", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	deleteAccountResp.Error = err
	s.JSONResponse(w, r, deleteAccountResp)
}
