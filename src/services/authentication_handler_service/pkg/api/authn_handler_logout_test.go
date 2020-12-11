package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/BlackspaceInc/BlackspacePlatform/src/services/authentication_handler_service/pkg/middleware"
)

func TestLogoutAccountHandler(t *testing.T) {
	var testDataInfo = []struct{
		newEmail string
		responseCode int
		errorExpectedToOcurr bool
		shouldCreateAndAuthenticateAccountFirst bool
	}{
		{
			// test case where we have a valid account and lock it
			fmt.Sprintf("test_%s@gmail.com", GenerateRandomString(10)),
			http.StatusOK,
			false,
			true,
		},
		{
			// test case where we have an invalid account and cant lock it
			fmt.Sprintf("test_%s@gmail.com", GenerateRandomString(10)),
			http.StatusBadRequest,
			true,
			false,
		},
	}

	for _, data := range testDataInfo {
		var result CreateAccountResponse
		var authRes LoginAccountResponse
		var err error

		// first we create the account
		if data.shouldCreateAndAuthenticateAccountFirst {
			result, err, authRes = createAndLoginAccountTestUtil(t, result, err, data.newEmail, data.errorExpectedToOcurr, authRes)
		}

		// try the lock operation
		rr, err := LogoutUserAccountRequestTestUtil(authRes.Token, t)

		if data.errorExpectedToOcurr && err == nil {
			t.Errorf("expected error to occur but none did")
		}

		if !data.errorExpectedToOcurr && err != nil {
			t.Errorf("error was not expected to occur - error %s", err.Error())
		}

		// Check the status code is what we expect.
		if status := rr.Code; status != data.responseCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, data.responseCode)
		}
	}
}

func LogoutUserAccountRequestTestUtil(token string, t *testing.T) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("POST", "/v1/account/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	req, rr := generateAuthorizedRequest(req, token)

	srv := NewMockServer()
	handler := http.HandlerFunc(srv.logoutHandler)

	handler.ServeHTTP(rr, req)

	return rr, err
}
