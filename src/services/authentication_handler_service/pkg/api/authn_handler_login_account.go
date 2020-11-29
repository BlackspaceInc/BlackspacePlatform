package api

import (
	"net/http"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

type LoginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAccountResponse struct {
	Error error  `json:"error"`
	Token string `json:"token"`
}

// Log into account request
// swagger:parameters loginAccountRequest
type loginAccountRequest struct {
	// in: body
	Body struct {
		// account email to log into
		// required : true
		Email string `json:"email"`
		// account password to log into
		// required : true
		Password string `json:"password"`
	}
}

// Account successfully created
// swagger:response loginAccountResponse
type loginAccountResponse struct {
	// in: body
	Body struct {
		// account auth token
		// required: true
		// example: sjfkhjasgdsfdjsh.ajgsdjkaskfgdkgsafd.afsdjaksjgdfas
		Token string `json:"id"`
		// error
		// required: true
		// example: account already exists
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/account/login loginAccountRequest
//
// Log into account
//
// logs into an account via the authentication service
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
//      200: loginAccountResponse
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 500: internalServerError
// creates an account
func (s *Server) loginAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		loginAccountReq LoginAccountRequest
	)

	err := helper.DecodeJSONBody(w, r, &loginAccountReq)
	if err != nil {
		// TODO: emit a metric
		s.logger.ErrorM(err, "failed to decode request")
		helper.ProcessMalformedRequest(w, err)
		return
	}

	if loginAccountReq.Password == "" || loginAccountReq.Email == "" {
		// TODO: emit a metric
		errMsg := "invalid input parameters. please specify a email and password"
		s.logger.ErrorM(err, "invalid input parameters")
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// TODO: perform this operation in a circuit breaker, emit a metric, and trace this
	token, customErr := s.authnClient.Handler.Login(loginAccountReq.Email, loginAccountReq.Password)
	s.logger.Info("status of login", "err",customErr)
	if _, err := helper.ProcessAggregatedErrors(w, customErr); err != nil {
		s.logger.ErrorM(err, "failed to login user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := LoginAccountResponse{Token: token, Error: err}
	s.JSONResponse(w, r, response)
}
