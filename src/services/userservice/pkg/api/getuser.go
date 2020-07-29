package api

import (
	"net/http"
	"strconv"

	"github.com/keratin/authn-go/authn"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

type GetUserResponse struct {
	User  models.UserORM `json:"user"`
	Error error          `json:"error"`
}

// Get user by id request
// swagger:parameters getUserRequest
type GetUserRequestSwagger struct {
	// user account to create
	// in: body
	Body struct {
		// id of the user account to get
		// in: query
		// required: true
		Id uint32 `json:"result"`
	}
}

// Common operation response
// swagger:response getUserResponse
type GetUserResponseSwagger struct {
	// in: body
	Body struct {
		// error
		// required: true
		// example: error occured while processing request
		Error error `json:"error"`
		// User
		// required: true
		User models.UserORM `json:"user"`
	}
}

// swagger:route GET /v1/user/{id} User getUserRequest
//
// Get User Account By ID
//
// Returns a user account by id
//
//     Consumes:
//     - application/json
//     - application/x-protobuf
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss
//
//     Security:
//       api_key:
//       oauth: read, write
// responses:
//      200: getUserResponse
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 401: unAuthorizedError
// 500: internalServerError
// gets a user by account id
func (s *Server) getUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		userAccount *models.UserORM
		response    GetUserResponse
	)

	id := helper.ExtractIDFromRequest(r)

	userAccount, err := s.db.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtain the account of interest from the authentication service as it
	// is the single source of truth for all authentication records tied to a
	// user
	fn := func() (interface{}, error) {
		id := int(userAccount.Authnid)
		return s.AuthnClient.Client.GetAccount(strconv.Itoa(id))
	}

	authnAccount, err := s.CircuitBreaker.PerformServiceRequest(r.Context(), s.config.AuthenticationServiceName, fn, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	// once the account is obtained, we assert that it is not locked or deleted
	// this is important to enforce that consistency at the record level is witheld
	// throughout the entire backend
	if authnAccount == nil {
		http.Error(w, "account authentication records do not exist", http.StatusNotFound)
		return
	}

	// assertion cast of account
	account, ok := authnAccount.(*authn.Account)
	if !ok {
		http.Error(w, "failed to perform authn account level assertion", http.StatusNotFound)
		return
	}

	if account.Locked {
		http.Error(w, "account is locked", http.StatusNotFound)
		return
	} else if account.Deleted {
		http.Error(w, "account is deleted", http.StatusNotFound)
		return
	}

	response.Error = err
	response.User = *userAccount

	// store the request in the database
	s.JSONResponse(w, r, response)
}
