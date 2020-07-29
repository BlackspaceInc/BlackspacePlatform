package api

import (
	"net/http"
	"strconv"

	"github.com/keratin/authn-go/authn"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
)

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JwtToken     string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Error        error  `json:"error"`
}

// Log in user request
// swagger:parameters loginUserRequest
type loginUserRequest struct {
	// in: body
	Body struct {
		// user name
		// required: true
		// example: test-username
		Username string `json:"username"`
		// password
		// required: true
		// example: test-password
		Password string `json:"password"`
	}
}

// Log in user response
// swagger:response loginUserResponse
type loginUserResponse struct {
	// in: body
	Body struct {
		// Jwt Token
		// required: true
		// example: kBxbjzKVDjvasgvds.askdhjaskjdgsagjcdgc.asjdjkasfgdas
		JwtToken string `json:"token"`
		// Refresh Token
		// required: true
		// example: kBxbjzKVDjvasgvds.askdhjaskjdgsagjcdgc.asjdjkasfgdas
		RefreshToken string `json:"refresh_token"`
		// error
		// required: true
		// example: unable to get token
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/user/login User loginUserRequest
//
// Log in user
//
// Logs in a user into the system
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
//      200: loginUserResponse
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 500: internalServerError
// logs in a team into the system
func (s *Server) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		loginUserRequest LoginUserRequest
	)

	// decode the data present in the body
	err := helper.DecodeJSONBody(w, r, &loginUserRequest)
	if err != nil {
		helper.ProcessMalformedRequest(w, err)
		return
	}

	// assert that the password and username fields are not empty
	if loginUserRequest.Password == "" || loginUserRequest.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and password", http.StatusBadRequest)
		return
	}

	// obtain the user by username
	_, user, err := s.db.GetUserIfExists(r.Context(), 0, loginUserRequest.Username, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// assert that the input password and the password of the obtained user from the database match
	if !s.db.ComparePasswords(user.Password, []byte(loginUserRequest.Password)) {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}

	fn := func() (interface{}, error) {
		id := int(user.Authnid)
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

	if account.Deleted || account.Locked {
		http.Error(w, "account locked or deleted", http.StatusBadRequest)
		return
	}

	// TODO - think about sending web tokens from authentication server itself
	/*
		aggErr, idToken := s.AuthnClient.Handler.Login(loginUserRequest.Username, loginUserRequest.Password)
		if aggErr != nil {
			helper.ProcessAggregatedErrors(w, aggErr)
			return
		}
	*/

	token, err := s.GenerateAndSignJwtToken(user.Id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse := LoginResponse{
		JwtToken:     token.AccessToken,
		RefreshToken: token.RefreshToken,
		Error:        nil,
	}

	s.JSONResponse(w, r, loginResponse)
}
