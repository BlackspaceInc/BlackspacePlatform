package api

import (
	"net/http"
	"strconv"
	"k8s.io/klog/v2"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

// UnLockAccountResponse is struct providing errors tied to Unlock account operations
type UnLockAccountResponse struct {
	Error error `json:"error"`
}

// UnLock account request
// swagger:parameters unlockAccount
type UnLockAccountRequest struct {
	// id of the account to unlock
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route POST /v1/account/unlock/{id} unlockAccount
// UnLock Account
//
// UnLocks an account through the authentication service
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
// unlocks an by account id
func (s *Server) unlockAccountHandler(w http.ResponseWriter, r *http.Request) {
	var unlockAccountResp UnLockAccountResponse
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
	if err = s.authnClient.Client.UnlockAccount(strconv.Itoa(int(authnID))); err != nil {
		// TODO: emit metrics
		klog.Error("failed to lock created account", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	unlockAccountResp.Error = err
	s.JSONResponse(w, r, unlockAccountResp)
}

