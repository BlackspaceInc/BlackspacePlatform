package api

import (
	"net/http"
	"time"

	utils "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-utilities"
	"github.com/opentracing/opentracing-go"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, parentSpan := s.startRootSpan(r, constants.LOGOUT_ACCOUNT)
	defer parentSpan.Finish()

	if s.IsNotAuthenticated(w, r) {
		return
	}

	// hit authn log out endpoint and return
	// we delete the session stored in the authentication service redis store
	var (
		begin = time.Now()
		took  = time.Since(begin)
		f     = func() error {
			if err := s.authnClient.LogOutAccount(); err != nil {
				s.logger.ErrorM(err, "status of logout")
				return err
			}
			return nil
		}
	)

	ctx = opentracing.ContextWithSpan(ctx, parentSpan)
	// TODO: perform this operation in a circuit breaker
	if err := s.RemoteOperationAndInstrument(ctx, f, constants.LOGOUT_ACCOUNT, &took); utils.HandleError(w, err,
		http.StatusInternalServerError) == true {
		s.logger.For(ctx).Error(err, "failed to perform logout account")
		return
	}

	s.JSONResponse(w, r, http.StatusOK)
}
