package api

import (
	"errors"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

type SignUpUserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type SignUpUserResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
}

// Sign up user request
// swagger:parameters signUpUserReq
type signUpUserRequest struct {
	// in: body
	Body struct {
		// user email address to create
		// required : true
		Email string `json:"email"`
		// user first name
		// required : true
		FirstName string `json:"firstname"`
		// user last name
		// required : true
		LastName string `json:"lastname"`
		// user username to create
		// required : true
		Username string `json:"username"`
		// user password to create
		// required : true
		Password string `json:"password"`
	}
}

// User Successfully signed up
// swagger:response signUpUserResp
type signUpResponse struct {
	// in: body
	Body struct {
		// user account id
		// required: true
		// example: 20
		Id uint32 `json:"id"`
		// error
		// required: true
		// example: user already exists
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/user/signup User signUpUserReq
//
// Sign Up User
//
// creates a user account object in the backend database
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
// creates a user account
func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var (
		signUpUserRequest SignUpUserRequest
		userAccount       *models.User
	)

	err := helper.DecodeJSONBody(w, r, &signUpUserRequest)
	if err != nil {
		helper.ProcessMalformedRequest(w, err)
		return
	}

	// obtain fields of interest from the frontend
	userAccount = new(models.User)
	userAccount.Email = signUpUserRequest.Email
	userAccount.Password = signUpUserRequest.Password
	userAccount.Username = signUpUserRequest.Username
	userAccount.FirstName = signUpUserRequest.Firstname
	userAccount.LastName = signUpUserRequest.Lastname

	if userAccount.Email == "" || userAccount.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and email", http.StatusBadRequest)
		return
	}

	authnID, err := s.CircuitBreaker.PerformServiceRequest(
		r.Context(),
		s.config.AuthenticationServiceName,
		func() (interface{}, error) {
			return s.AuthnClient.Client.ImportAccount(userAccount.Username, userAccount.Password, false)
		}, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func() {
		if err != nil {
			s.logger.Error("unable to create user account in authentication service. archiving account", zap.Error(err))
			err = s.CircuitBreaker.PerformCustomServiceRequest(
				r.Context(),
				s.config.AuthenticationServiceName,
				func() error {
					authnID, ok := authnID.(int)
					if !ok {
						return errors.New("failed type assertations")
					}
					return s.AuthnClient.Client.ArchiveAccount(strconv.Itoa(authnID))
				}, nil)
		}
	}()

	// we still continue account creation even if the call to the authentication service fails
	ID, ok := authnID.(uint32)
	if !ok {
		http.Error(w, "type assertion error", http.StatusBadRequest)
		return
	}
	userAccount.Authnid = ID

	// create the user account in our backend
	createdUserAccount, err := s.db.CreateUser(r.Context(), userAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: once authentication working rethink this design. JWT should come from authentication service
	// return a response
	response := SignUpUserResponse{Id: createdUserAccount.Id, Error: err}
	s.JSONResponse(w, r, response)
}
