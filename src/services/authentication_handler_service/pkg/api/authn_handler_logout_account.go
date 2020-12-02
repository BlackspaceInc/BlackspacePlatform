package api

import (
	"net/http"
	"time"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/constants"
	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// hit authn log out endpoint and return
	// we delete the session stored in the authentication service redis store
	var (
		begin = time.Now()
		took  = time.Since(begin)
		f = func() error {
			if authErr := s.authnClient.Handler.LogOut(); authErr != nil {
				s.logger.Info("status of logout", "err", authErr)
				if _, err := helper.ProcessAggregatedErrors(w, authErr); err != nil {
					return err
				}
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
