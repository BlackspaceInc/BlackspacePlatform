package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
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

	err := s.DecodeRequestAndInstrument(w, r, &loginAccountReq, constants.LOGIN_ACCOUNT)
	if err != nil {
		s.logger.ErrorM(err, "failed to decode request")
		helper.ProcessMalformedRequest(w, err)
		return
	}

	if loginAccountReq.Password == "" || loginAccountReq.Email == "" {
		s.metrics.InvalidRequestParametersCounter.WithLabelValues(constants.LOGIN_ACCOUNT).Inc()

		errMsg := "invalid input parameters. please specify a email and password"
		s.logger.ErrorM(err, "invalid input parameters")

		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var (
		startTime = time.Now()
		elapsedTime  = time.Since(startTime)
		op = func() (interface{},error) {
			token, aggregatedErr :=  s.authnClient.Handler.Login(loginAccountReq.Email, loginAccountReq.Password)
			s.logger.Info("status of login", "err", aggregatedErr)
			if _, err := helper.ProcessAggregatedErrors(w, aggregatedErr); err != nil {
				return token, err
			}
			return token, nil
		}
	)

	result, err := s.RemoteOperationAndInstrumentWithResult(op, constants.LOGIN_ACCOUNT, &elapsedTime)
	if err != nil {
		s.logger.ErrorM(err, "failed to login user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	token, ok := result.(string)
	if !ok {
		err := errors.New("failed to cast from interface type")
		s.logger.ErrorM(err, "casting error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := LoginAccountResponse{Token: token, Error: err}
	s.JSONResponse(w, r, response)
}
