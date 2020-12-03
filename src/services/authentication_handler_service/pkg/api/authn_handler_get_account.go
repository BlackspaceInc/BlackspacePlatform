package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/keratin/authn-go/authn"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/middleware"
)

// GetAccountResponse is struct providing errors tied to get account operations
type GetAccountResponse struct {
	Account *authn.Account `json:"account"`
	Error   error          `json:"error"`
}

// Get account by id request
// swagger:parameters getAccount
type GetAccountRequest struct {
	// id of the account to obtain
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route GET /v1/account/{id} getAccount
// Get Account
//
// Gets an account through the authentication service
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
func (s *Server) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	// we extract the user id from the url initially
	authnID, err := s.ExtractIdOperationAndInstrument(r, constants.GET_ACCOUNT)
	if err != nil {
		s.logger.ErrorM(err, "failed to parse account id from url")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !middleware.IsAuthenticated(r.Context()){
		s.logger.InfoM("user not authenticated")
	}

	var (
		begin = time.Now()
		took  = time.Since(begin)
		f = func() (interface{}, error){
			return s.authnClient.GetAccount(strconv.Itoa(int(authnID)))
		}
	)

	result, err := s.RemoteOperationAndInstrumentWithResult(f, constants.GET_ACCOUNT, &took)
	if err != nil {
		s.logger.ErrorM(err, "failed to get account")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	account, ok := result.(*authn.Account)
	if !ok {
		s.metrics.CastingOperationFailureCounter.WithLabelValues(constants.GET_ACCOUNT)
		err  := errors.New("failed to cast response to account object")
		s.logger.ErrorM(err, "casting failure")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var getAccountResp = GetAccountResponse{
		Account: account,
		Error:   err,
	}

	s.JSONResponse(w, r, getAccountResp)
}
