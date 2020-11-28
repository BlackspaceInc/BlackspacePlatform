package api

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"k8s.io/klog/v2"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

type UpdateAccountRequest struct {
	Email string
}

type UpdateAccountResponse struct {
	Error error
}

// Update account request
// swagger:parameters updateAccount
type UpdateAccountRequestSwagger struct {
	// user account to update
	// in: body
	Body struct {
		// Email to update
		// required: true
		Email string `json:"email"`
	}
	// id of account to update
	// in: query
	AccountId uint32
}

// Common operation response
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

// swagger:parameters updateAccount
type accountIdParam struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:route PUT POST /v1/account/update/{id} updateAccount
//
// Update Account
//
// Updates an account through the authentication service
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
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
// updates an account credentials in the authentication service
func (s *Server) updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		updateAccountReq  UpdateAccountRequest
		updateAccountResp UpdateAccountResponse
	)

	// TODO: emit metrics
	authnID, err := helper.ExtractIDFromRequest(r)
	if err != nil {
		// TODO: emit metrics
		klog.Error("failed to parse account id from url", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: emit metrics
	// decode the update user request
	if err := helper.DecodeJSONBody(w, r, &updateAccountReq); err != nil {
		// TODO: emit metrics
		klog.Error("failed to decode request body", zap.Error(err))
		helper.ProcessMalformedRequest(w, err)
		return
	}

	// assert password and email field is present.
	if updateAccountReq.Email == "" {
		// TODO: emit metrics
		errMsg := "invalid input parameters. please specify a email"
		klog.Error(errMsg, zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: emit metrics, and trace
	if err := s.authnClient.Client.Update(strconv.Itoa(int(authnID)), updateAccountReq.Email); err != nil {
		klog.Error("failed to update the account through the authentication service", "id", authnID, zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	klog.Info("successfully updated account through authentication service", "id", authnID)
	updateAccountResp.Error = err
	s.JSONResponse(w, r, updateAccountResp)
}
