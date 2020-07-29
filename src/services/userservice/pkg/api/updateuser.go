package api

import (
	"net/http"
	"strconv"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/models"
)

type UpdateUserRequest struct {
	User models.User
}

type UpdateUserResponse struct {
	Error error
}

// Update user request
// swagger:parameters updateUser
type UpdateUserRequestSwagger struct {
	// user account to create
	// in: body
	Body struct {
		// user to update
		// required: true
		User models.User `json:"result"`
	}
	// user id of account to update
	// in: query
	UserAccountId uint32
}

// OperationResponseSwagger Defines a common operation response
// swagger:response operationResponse
type OperationResponseSwagger struct {
	// in: body
	Body struct {
		// error
		// required: true
		// example: error occured while processing request
		Error error `json:"error"`
	}
}

// swagger:parameters updateUser getUserRequest deleteUser
type userIdParam struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:route PUT /v1/user/{id} User updateUser
//
// Update User Account
//
// Updates a user account present in the backend database
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
//      200: operationResponse
// 400: badRequestError
// 404: notFoundError
// 403: forbiddenError
// 406: genericError
// 401: unAuthorizedError
// 500: internalServerError
// gets a user by account id
func (s *Server) updatedUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		updateUserRequest  UpdateUserRequest
		updateUserResponse UpdateUserResponse
	)

	// attempt to first obtain user as a record should exist in the database
	id := helper.ExtractIDFromRequest(r)

	// attempt to obtain the user based on the id in the database
	_, err := s.db.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// decode the update user request
	err = helper.DecodeJSONBody(w, r, &updateUserRequest)
	if err != nil {
		helper.ProcessMalformedRequest(w, err)
		return
	}

	// assert email and username fields are present.
	if updateUserRequest.User.Email == "" || updateUserRequest.User.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and email", http.StatusBadRequest)
		return
	}

	// validate that all necessary required fields are present
	if err := updateUserRequest.User.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the updated value is the username field by attempting to obtain
	// the user of interest by the current username. If successfully then we
	// know the username field has not been updated. If unsuccessful, then the
	// username does not exist. In such a scenario we perform one additional check
	// (get user by id) in order to fully assert the user does exists in the db
	// prior to performing an update api call to the authentication service
	exist, _, err := s.db.GetUserIfExists(r.Context(), 0, updateUserRequest.User.Username, "")
	if !exist {
		// the username field is indeed being updated
		// must check if the user actually does exist based on the user id
		exist, userAccount, err := s.db.GetUserIfExists(r.Context(), updateUserRequest.User.Id, "", "")
		if !exist {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send the update username request to the authentication service
		if err := s.CircuitBreaker.PerformCustomServiceRequest(
			r.Context(),
			s.config.AuthenticationServiceName,
			func() error {
				authnID := int(userAccount.Authnid)
				return s.AuthnClient.Client.Update(strconv.Itoa(authnID), updateUserRequest.User.Username)
			}, nil); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// attempt user update once we have performed the prior checks
	_, err = s.db.UpdateUser(r.Context(), &updateUserRequest.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateUserResponse.Error = err
	s.JSONResponse(w, r, updateUserResponse)
}

func (s *Server) ExtractJwtFromRequest(w http.ResponseWriter, r *http.Request) (uint32, error) {
	authTokenRes, err := s.ExtractJwtFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	id, err := strconv.Atoi(authTokenRes.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	userID := uint32(id)
	return userID, err
}
