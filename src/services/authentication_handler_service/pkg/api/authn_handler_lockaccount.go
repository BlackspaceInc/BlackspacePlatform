package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
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
	authnID, err := s.ExtractIdOperationAndInstrument(r, constants.LOCK_ACCOUNT)
	if err != nil {
		// TODO: emit metrics
		s.logger.ErrorM(err, "failed to parse account id from url")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		begin = time.Now()
		took  = time.Since(begin)
		f = func() error {
			return s.authnClient.LockAccount(strconv.Itoa(int(authnID)))
		}
	)

	// TODO: perform this operation in a circuit breaker, and trace this
	if err = s.RemoteOperationAndInstrument(f, constants.LOCK_ACCOUNT, &took); err != nil {
		s.logger.ErrorM(err, "failed to lock created account")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	lockAccountResp.Error = err
	s.JSONResponse(w, r, lockAccountResp)
}
