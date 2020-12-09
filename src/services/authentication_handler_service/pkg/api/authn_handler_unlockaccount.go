package api

import (
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
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
	if s.IsNotAuthenticated(w, r) {
		return
	}

	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	var unlockAccountResp UnLockAccountResponse
	// we extract the user id from the url initially
	authnID, err := s.ExtractIdOperationAndInstrument(r, constants.UNLOCK_ACCOUNT)
	if err != nil {
		s.logger.ErrorM(err, "failed to parse account id from url")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		begin = time.Now()
		took  = time.Since(begin)
		f = func() error {
			return s.authnClient.UnlockAccount(strconv.Itoa(int(authnID)))
		}
	)

	// TODO: perform this operation in a circuit breaker, and trace this
	if err = s.RemoteOperationAndInstrument(f, constants.UNLOCK_ACCOUNT, &took); err != nil {
		s.logger.ErrorM(err, "failed to unlock created account")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	unlockAccountResp.Error = err
	s.JSONResponse(w, r, unlockAccountResp)
}
