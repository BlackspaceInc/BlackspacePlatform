package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/middleware"
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
	if s.IsNotAuthenticated(w, r) {
		return
	}

	var deleteAccountResp DeleteAccountResponse

	// we extract the user id from the url initially
	authnID, err := s.ExtractIdOperationAndInstrument(r, constants.DELETE_ACCOUNT)
	if err != nil {
		s.logger.ErrorM(err, "failed to parse account id from url")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		begin = time.Now()
		took  = time.Since(begin)
		f     = func() error {
			return s.authnClient.ArchiveAccount(strconv.Itoa(int(authnID)))
		}
	)

	// TODO: perform this operation in a circuit breaker, and trace this
	if err = s.RemoteOperationAndInstrument(f, constants.DELETE_ACCOUNT, &took); err != nil {
		s.logger.ErrorM(err, "failed to archive created account")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteAccountResp.Error = err
	s.JSONResponse(w, r, deleteAccountResp)
}

func (s *Server) IsNotAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	if !middleware.IsAuthenticated(r.Context()) {
		err := errors.New("user not authenticated")
		s.logger.ErrorM(err, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}
