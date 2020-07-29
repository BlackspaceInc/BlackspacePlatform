package api

import (
	"net/http"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// hit authn log out endpoint and return
	// we delete the session stored in the authentication service redis store
	authErr := s.AuthnClient.Handler.LogOut()

	if authErr != nil {
		if helper.ProcessAggregatedErrors(w, authErr) {
			return
		}
	}

	s.JSONResponse(w, r, http.StatusOK)
}
