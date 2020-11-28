package api

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"k8s.io/klog/v2"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

// LockAccountResponse is struct providing errors tied to lock account operations
type LockAccountResponse struct {
	Error error `json:"error"`
}

// Lock account request
// swagger:parameters lockAccount
type LockAccountRequest struct {
	// id of the account to lock
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route POST /v1/account/lock/{id} lockAccount
// Lock Account
//
// Locks an account through the authentication service
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
// locks an account by account id
func (s *Server) lockAccountHandler(w http.ResponseWriter, r *http.Request) {
	var lockAccountResp LockAccountResponse
	// we extract the user id from the url initially
	// TODO: emit metrics
	authnID, err := helper.ExtractIDFromRequest(r)
	if err != nil {
		// TODO: emit metrics
		klog.Error("failed to parse account id from url", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: emit a metric, and trace this
	if err = s.authnClient.Client.LockAccount(strconv.Itoa(int(authnID))); err != nil {
		// TODO: emit metrics
		klog.Error("failed to lock created account", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lockAccountResp.Error = err
	s.JSONResponse(w, r, lockAccountResp)
}
