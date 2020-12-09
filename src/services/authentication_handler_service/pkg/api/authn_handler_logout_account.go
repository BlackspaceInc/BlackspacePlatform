package api

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if s.IsNotAuthenticated(w, r) {
		return
	}

	ctx := r.Context()
	s.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	// hit authn log out endpoint and return
	// we delete the session stored in the authentication service redis store
	var (
		begin = time.Now()
		took  = time.Since(begin)
		f = func() error {
			if err := s.authnClient.LogOutAccount(); err != nil {
				s.logger.ErrorM(err,"status of logout")
				return err
			}
			return nil
		}
	)

	// TODO: perform this operation in a circuit breaker, and trace this
	if err := s.RemoteOperationAndInstrument(f, constants.LOGOUT_ACCOUNT, &took); err != nil {
		s.logger.ErrorM(err, "failed to perform logout account")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.JSONResponse(w, r, http.StatusOK)
}
