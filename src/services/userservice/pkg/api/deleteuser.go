package api

import (
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
	"net/http"
	"strconv"
)

// OperationResponse is struct providing errors tied to common operations
type OperationResponse struct {
	Error error `json:"error"`
}

// DeleteUserByIdRequest Defines a request to delete user by id
// swagger:parameters deleteUser
type DeleteUserByIdRequest struct {
	// id of the user account to delete
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route DELETE /v1/user/{id} User deleteUser
// Delete User Account
//
// Deletes a user account present in the backend database
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
// deletes a user by account id
func (s *Server) deleteUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var response OperationResponse
	// we extract the user id from the url initially
	id := helper.ExtractIDFromRequest(r)

	// Ensure the user account actually exists
	exist, userAccount, err := s.db.GetUserIfExists(r.Context(), id, "", "")
	if !exist || err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO - refactor and change this
	// Important to first attempt deletion of the credentials through the authentication
	// service and then to the record in the user's service database
	// perform the request to the authentication service in a retry loop
	if err := s.CircuitBreaker.PerformCustomServiceRequest(
		r.Context(),
		s.config.AuthenticationServiceName,
		func() error {
			authnID := int(userAccount.Authnid)
			return s.AuthnClient.Client.ArchiveAccount(strconv.Itoa(authnID))
		}, nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// delete the record in the backend once successfully able to delete from the authentication service
	err = s.db.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Error = err
	s.JSONResponse(w, r, response)
}
