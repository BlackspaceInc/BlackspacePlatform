package api

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"k8s.io/klog/v2"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

type CreateAccountRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type CreateAccountResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
}

// Create account request
// swagger:parameters createAccountRequest
type createAccountRequest struct {
	// in: body
	Body struct {
		// user username to create
		// required : true
		Email string `json:"email"`
		// user password to create
		// required : true
		Password string `json:"password"`
	}
}

// Account successfully created
// swagger:response createAccountResponse
type createAccountResponse struct {
	// in: body
	Body struct {
		// account id
		// required: true
		// example: 20
		Id uint32 `json:"id"`
		// error
		// required: true
		// example: account already exists
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/account/create AccountDetails createAccountRequest
//
// Create Account
//
// creates an account object via the authentication service
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
//
//     Security:
//       api_key:
//       oauth: read, write
// responses:
//      200: signUpUserResp
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 500: internalServerError
// creates an account
func (s *Server) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		createAccountReq CreateAccountRequest
	)

	err := helper.DecodeJSONBody(w, r, &createAccountReq)
	if err != nil {
		// TODO: emit a metric
		klog.Error("failed to decode request", zap.Error(err))
		helper.ProcessMalformedRequest(w, err)
		return
	}

	if createAccountReq.Password == "" || createAccountReq.Email == "" {
		// TODO: emit a metric
		errMsg := "invalid input parameters. please specify a username and password"
		klog.Error("invalid input parameters", "error", errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// TODO: emit a metric, and trace this
	authnID, err := s.authnClient.Client.ImportAccount(createAccountReq.Email, createAccountReq.Password, false)
	if err != nil {
		// TODO: emit a metric
		klog.Error("failed to create account via authentication service", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// this is ran in the case any error is encountered when this function returns. We don't want to leave the datastore of the authentication
	// service in an inconsistent state
	defer func() {
		if err != nil {
			// TODO: perform this operation in a circuit breaker, emit a metric, and trace this
			klog.Error("unable to create user account in authentication service. archiving account", zap.Error(err))
			if err = s.authnClient.Client.ArchiveAccount(strconv.Itoa(authnID)); err != nil {
				klog.Error("failed to archive created account", zap.Error(err))
			}
		}
	}()

	response := CreateAccountResponse{Id: uint32(authnID), Error: err}
	s.JSONResponse(w, r, response)
}
