package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *Authentication) SignUp(username, password string) (*AggregatedError, *TokenId) {
	var (
		result TokenId
	)

	credentials := Credentials{
		Username: username,
		Password: password,
	}

	jsonStr, err := json.Marshal(&credentials)

	request, err := http.NewRequest("POST", a.AccountsBase, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	// publically available endpoint
	err, body := a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &result
}

func (a *Authentication) GetAccount(id uint32, username, password string) (*AggregatedError, *AuthAccount) {
	var (
		authAccount AuthAccount
	)

	url := a.AccountsBase + "/" + fmt.Sprint(id)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	// private api endpoint
	err, body := a.SetHeadersAndPerformRequest(request, username, password, false)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err = json.Unmarshal(body, &authAccount)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &authAccount
}

func (a *Authentication) UpdateUsername(username, password string, id uint32) *AggregatedError {
	url := a.AccountsBase + "/" + fmt.Sprint(id)
	jsonStr, err := json.Marshal(&UsernamePayload{Username: username})

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	// private api endpoint
	err, _ = a.SetHeadersAndPerformRequest(request, username, password, false)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	return nil
}

func (a *Authentication) IsUsernameAvailable(username string) (*AggregatedError, bool) {
	jsonStr, err := json.Marshal(&UsernamePayload{Username: username})

	request, err := http.NewRequest("GET", a.Availability, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, false
	}

	// public api endpoint
	err, _ = a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, false
	}

	return nil, true
}

func (a *Authentication) LockOrUnlockAccount(id int, username, password string, lock bool) *AggregatedError {
	var url string

	if lock {
		url = a.AccountsBase + "/" + fmt.Sprint(id) + "/lock"
	} else {
		url = a.AccountsBase + "/" + fmt.Sprint(id) + "/unlock"
	}

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	err, _ = a.SetHeadersAndPerformRequest(request, username, password, false)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	return nil
}

func (a *Authentication) DeleteAccount(username, password string, id uint32) *AggregatedError {
	url := a.AccountsBase + "/" + fmt.Sprint(id)

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	err, _ = a.SetHeadersAndPerformRequest(request, username, password, false)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	return nil
}

func ExtractErrorResponse(resp *http.Response, err error, body []byte, authError AuthError) (*AggregatedError, bool) {
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = json.Unmarshal(body, &authError)
		if err != nil {
			return &AggregatedError{Error: err, AuthErrorMsg: nil}, true
		}
		return &AggregatedError{Error: nil, AuthErrorMsg: &authError}, true
	}
	return nil, false
}

func (a *Authentication) UpdateAccount(username, password string, locked bool) (*AggregatedError, int) {
	var (
		updateAccountResponse UpdateAccountResponse
	)

	jsonStr, err := json.Marshal(&UpdateAccount{Username: username, Password: password, Locked: locked})
	request, err := http.NewRequest("POST", a.Import, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, 0
	}

	err, body := a.SetHeadersAndPerformRequest(request, username, password, false)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, 0
	}

	err = json.Unmarshal(body, &updateAccountResponse)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, 0
	}

	id, err := strconv.Atoi(updateAccountResponse.Result.Id)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, 0
	}

	return nil, id
}

func (a *Authentication) Login(username, password string) (*AggregatedError, *string) {
	var (
		accountResponse UpdateAccountResponse
	)
	jsonStr, err := json.Marshal(&LoginAccount{Username: username, Password: password})
	request, err := http.NewRequest("POST", a.SessionBase, bytes.NewBuffer(jsonStr))
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err, body := a.SetHeadersAndPerformRequest(request, username, password, true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err = json.Unmarshal(body, &accountResponse)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &accountResponse.Result.Id
}

func (a *Authentication) RefreshToken() (*AggregatedError, *TokenId) {
	var (
		result TokenId
	)

	request, err := http.NewRequest("GET", a.RefreshSession, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err, body := a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &result
}

func (a *Authentication) LogOut() *AggregatedError {
	request, err := http.NewRequest("DELETE", a.SessionBase, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	err, _ = a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}
	}

	return nil
}

func (a *Authentication) GetJwks(uri string) (*AggregatedError, *JsonKeys) {
	var jwtKeys JsonKeys
	var url string

	if uri == "" {
		url = a.ServiceBaseAddress + "/jwks"
	} else {
		url = "http://" + uri
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err, body := a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err = json.Unmarshal(body, &jwtKeys)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &jwtKeys
}

func (a *Authentication) GetJwtPublicKey() (*AggregatedError, *JwtConfiguration) {
	var jwtConfig JwtConfiguration
	url := a.ServiceBaseAddress + "/configuration"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	err, _ = a.SetHeadersAndPerformRequest(request, "", "", true)
	if err != nil {
		return &AggregatedError{Error: err, AuthErrorMsg: nil}, nil
	}

	return nil, &jwtConfig
}
