package api

import (
	"net/http"

	"github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/helper"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// hit authn log out endpoint and return
	// we delete the session stored in the authentication service redis store
	authErr := s.authnClient.Handler.LogOut()

	if authErr != nil {
		if _, err := helper.ProcessAggregatedErrors(w, authErr); err != nil {
			s.logger.ErrorM(err, "failed to perform log out request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	s.JSONResponse(w, r, http.StatusOK)
}
