package authentication

// Credential encapsulates user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TokenId stores a reference to a returned token id from the authn open source authentication service
type TokenId struct {
	Result TokenResult `json:"result"`
}

// TokenResult encompasses the returned token id json response from the auth service
type TokenResult struct {
	Token string `json:"id_token"`
}

// AuthnAccount witholds account details returned from the auth service
type AuthAccount struct {
	Result AccountResult `json:"result"`
}

// AccountResult represents the account records stored in the auth service
type AccountResult struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Locked   bool   `json:"locked"`
	Deleted  bool   `json:"deleted"`
}

// UsernamePayload models the user's username in json
type UsernamePayload struct {
	Username string `json:"username"`
}

// UpdateAccount captures the set of modifiable fields
type UpdateAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Locked   bool   `json:"locked"`
}

type UpdateAccountResponse struct {
	Result IdResult `json:"result"`
}

type IdResult struct {
	Id string `json:"id_token"`
}

type LoginAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomError struct {
	AuthErrorMsg *AuthError `json:"authentication_error"`
	Error        error      `json:"error"`
}

type AuthError struct {
	AuthErrorMessage []ErrorFields `json:"errors"`
}

type ErrorFields struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type JsonKeys struct {
	Result []JsonWebKeys `json:"keys"`
}

type JsonWebKeys struct {
	Use string `json:"use"`
	Alg string `json:"alg"`
	Kty string `json:"kty"`
	Kid string `json:"kid`
	E   string `json:"e"`
	N   string `json:"n"`
}

type JwtConfiguration struct {
	Issuer                 string   `json:"issuer"`
	ResponseTypesSupported []string `json:"response_types_supported"`
	SubjectTypesSupported  []string `json:"subject_types_supported"`
	TokenSigningAlgo       []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported        []string `json:"claims_supported"`
	JwtPublicKeyURI        string   `json:"jwks_uri"`
}
