package authentication

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/version"
	"github.com/BlackspaceInc/common/circuitbreaker"
)

const (
	AccountsUrl = "http://localhost:3000/accounts"
	Origin      = "http://localhost"
)

type Authentication struct {
	Origin                           string
	ServiceBaseAddress               string
	AccountsBase                     string
	SessionBase                      string
	Availability                     string
	Import                           string
	RefreshSession                   string
	EnablePrivateEndpointInteraction bool
	Timeout                          time.Duration
	AuthUsername                     string
	AuthPassword                     string
	CircuitBreaker                   *circuitbreaker.CircuitBreaker
}

// Obtain authentication service
func NewAuthenticationService(origin, authPort string, enablePrivateEndpointInteraction bool, httpClientTimeout time.Duration, username, password string, cb *circuitbreaker.CircuitBreaker) *Authentication {
	if origin == "" {
		origin = Origin
	}

	srvAddr := origin + ":" + authPort

	return &Authentication{
		Origin:                           origin,
		ServiceBaseAddress:               srvAddr,
		AccountsBase:                     srvAddr + "/accounts",
		Availability:                     srvAddr + "/accounts/available",
		Import:                           srvAddr + "/accounts/import",
		SessionBase:                      srvAddr + "/session",
		RefreshSession:                   srvAddr + "/session/refresh",
		EnablePrivateEndpointInteraction: enablePrivateEndpointInteraction,
		Timeout:                          httpClientTimeout,
		AuthUsername:                     username,
		AuthPassword:                     password,
		CircuitBreaker:                   cb,
	}
}

func (a *Authentication) SetHeadersAndPerformRequest(request *http.Request, username, password string, isPublic bool) (error, []byte) {
	if !isPublic {
		request.SetBasicAuth(a.AuthUsername, a.AuthPassword)
	}

	request.Header.Set("Origin", a.Origin)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Version", version.VERSION)
	request.Header.Set("X-API-Revision", version.REVISION)

	ctx, cancel := context.WithTimeout(request.Context(), a.Timeout)
	defer cancel()

	// call backend
	request = request.WithContext(ctx)

	resp, err := a.CircuitBreaker.PerformRequest(ctx, "authentication-service", request, nil)
	return err, resp
}

type ParseJwtResponse struct {
	Auth     []string         `json:"aud"`
	AuthTime *jwt.NumericDate `json:"auth_time"`
	Exp      *jwt.NumericDate `json:"exp"`
	Iat      *jwt.NumericDate `json:"iat"`
	Iss      string           `json:"iss"`
	Sub      string           `json:"sub"`
}

func ExtractClaims(tokenStr string) (uint32, error) {
	parsed, err := jose.ParseSigned(tokenStr)
	if err != nil {
		return 0, err
	}

	if parsed == nil || parsed.Signatures == nil || len(parsed.Signatures) == 0 {
		return 0, errors.New("invalid parsed jwt token")
	}

	// since we do not have the private key impossible to verify hence
	// we use an unsafe approach for now and obtain the payload
	payload := parsed.UnsafePayloadWithoutVerification()

	var response ParseJwtResponse
	_ = json.Unmarshal(payload, &response)

	id, err := strconv.Atoi(response.Sub)
	if err != nil {
		return 0, nil
	}
	return uint32(id), nil
}
