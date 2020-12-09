package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	utils "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-utilities"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/middleware"
)

// DeleteAccountResponse is struct providing errors tied to delete account operations
type DeleteAccountResponse struct {
	Error error `json:"error"`
}

// Delete account by id request
// swagger:parameters deleteAccount
type DeleteAccountRequest struct {
	// id of the account to delete
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route DELETE /v1/account/delete/{id} deleteAccount
// Delete Account
//
// Deletes an account through the authentication service
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
func (s *Server) deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	s.logger.For(ctx).InfoM("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// start a parent span
	spanCtx, _ := s.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	parentSpan := s.tracer.StartSpan("DeleteAccountRequest", ext.RPCServerOption(spanCtx))
	defer parentSpan.Finish()

	if s.IsNotAuthenticated(w, r) {
		return
	}

	var deleteAccountResp DeleteAccountResponse

	// we extract the user id from the url initially
	authnID, err := s.ExtractIdOperationAndInstrument(r, constants.DELETE_ACCOUNT)
	if utils.HandleError(w, err, http.StatusInternalServerError) {
		s.logger.For(ctx).Error(err, "failed to parse account id from url")
		return
	}

	var (
		begin = time.Now()
		took  = time.Since(begin)
		f     = func() error {
			return s.authnClient.ArchiveAccount(strconv.Itoa(int(authnID)))
		}
	)

	// TODO: perform this operation in a circuit breaker
	if err = s.RemoteOperationAndInstrument(f, constants.DELETE_ACCOUNT, &took, parentSpan.Context()); utils.HandleError(w, err,
		http.StatusInternalServerError) {
		s.logger.For(ctx).Error(err, "failed to archive created account")
		return
	}

	s.logger.For(ctx).Info("Successfully deleted user account", zap.Int("accountID", int(authnID)))
	deleteAccountResp.Error = err
	s.JSONResponse(w, r, deleteAccountResp)
}

func (s *Server) IsNotAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	if !middleware.IsAuthenticated(r.Context()) {
		err := errors.New("user not authenticated")
		s.logger.ErrorM(err, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}
