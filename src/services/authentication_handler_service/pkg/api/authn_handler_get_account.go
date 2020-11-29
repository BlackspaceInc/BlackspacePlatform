package api

import (
	"net/http"
	"strconv"

	"github.com/keratin/authn-go/authn"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
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
	// TODO: emit metrics
	authnID, err := helper.ExtractIDFromRequest(r)
	if err != nil {
		// TODO: emit metrics
		s.logger.ErrorM(err, "failed to parse account id from url")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: emit a metric, and trace this
	account, err := s.authnClient.Client.GetAccount(strconv.Itoa(int(authnID)))
	if err != nil {
		// TODO: emit metrics
		s.logger.ErrorM(err, "failed to get account")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var getAccountResp = GetAccountResponse{
		Account: account,
		Error:   err,
	}

	s.JSONResponse(w, r, getAccountResp)
}
