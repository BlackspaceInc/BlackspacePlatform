package helper_test

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/api"
	"github.com/BlackspaceInc/Backend/user-management-service/pkg/helper"
)

func TestValidJsonBody(t *testing.T) {
	// arrange
	loginReq := api.LoginUserRequest{
		Username: "testUser",
		Password: "testPassword",
	}

	jsonStr, err := json.Marshal(&loginReq)
	assert.Empty(t, err)
	req, err := http.NewRequest("POST", "/v1/test/endpoint", bytes.NewBuffer(jsonStr))
	assert.Empty(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.ResponseRecorder{}

	// act
	err = helper.DecodeJSONBody(&resp, req, &api.LoginUserRequest{})

	// assert
	assert.Empty(t, err)
}

func TestDecodeJsonBodyInvalidContentType(t *testing.T) {
	// arrange
	var tests = []struct {
		contentType string
		statusCode  int
	}{
		{"audio/aac", 400},
		{"application/x-abiword", 400},
		{"application/x-freearc", 400},
		{"video/x-msvideo", 400},
		{"application/vnd.amazon.ebook", 400},
		{"application/octet-stream", 400},
		{"image/bmp", 400},
		{"application/x-bzip", 400},
		{"application/x-bzip2", 400},
		{"application/x-csh", 400},
		{"text/css", 400},
		{"text/csv", 400},
		{"application/msword", 400},
		{"application/vnd.openxmlformats-officedocument.wordprocessingml.document", 400},
		{"application/vnd.ms-fontobject", 400},
		{"application/epub+zip", 400},
		{"application/gzip", 400},
		{"image/gif", 400},
		{"text/html", 400},
		{"image/vnd.microsoft.icon", 400},
		{"text/calendar", 400},
		{"image/jpeg", 400},
		{"application/ld+json", 400},
		{"audio/midi", 400},
		{"audio/x-midi", 400},
		{"text/javascript", 400},
		{"audio/mpeg", 400},
		{"video/mpeg", 400},
		{"application/ogg", 400},
		{"audio/opus", 400},
		{"font/otf", 400},
		{"image/png", 400},
		{"application/pdf", 400},
		{"application/x-httpd-php", 400},
		{"application/vnd.ms-powerpoint", 400},
		{"application/vnd.openxmlformats-officedocument.presentationml.presentation", 400},
		{"application/vnd.rar", 400},
		{"application/rtf", 400},
		{"application/x-sh", 400},
		{"image/svg+xml", 400},
		{"application/x-shockwave-flash", 400},
		{"application/x-tar", 400},
		{"image/tiff", 400},
		{"video/mp2t", 400},
		{"font/ttf", 400},
		{"text/plain", 400},
		{"application/vnd.visio", 400},
		{"audio/wav", 400},
		{"audio/webm", 400},
		{"video/webm", 400},
		{"image/webp", 400},
		{"font/woff2", 400},
		{"font/woff", 400},
		{"application/xhtml+xml", 400},
		{"application/xhtml+xml", 400},
		{"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", 400},
		{"application/xml", 400},
		{"application/vnd.mozilla.xul+xml", 400},
		{"application/zip", 400},
		{"video/3gpp", 400},
		{"audio/3gpp", 400},
		{"video/3gpp2", 400},
		{"audio/3gpp2", 400},
		{"application/x-7z-compressed", 400},
	}

	req, resp := generateRequestInitialConditions(t, 20)

	// act & assert
	for _, tt := range tests {
		testname := fmt.Sprintf("TestName: %s_test", tt.contentType)
		t.Run(testname, func(t *testing.T) {
			req.Header.Set("Content-Type", tt.contentType)
			err := helper.DecodeJSONBody(&resp, req, &api.LoginUserRequest{})

			var mr *helper.MalformedRequest
			if errors.As(err, &mr) == false || mr.Status != tt.statusCode {
				t.Errorf("got %d, want %d", mr.Status, tt.statusCode)
			}
		})
	}
}

/*
func TestDecodeJsonBodyInvalidMaxBytesWritten(t *testing.T){
	// arrange
	req, resp := generateRequestInitialConditions(t, helper.MAX_REQUEST_READ_SIZE)

	// act
	err := helper.DecodeJSONBody(&resp, req, &api.LoginUserRequest{})

	// TODO enhance this test
	// assertTestDecodeJsonBodyInvalidMaxBytesWritten
	var mr *helper.MalformedRequest
	if errors.As(err, &mr) == false || !strings.Contains(mr.Msg,helper.BADLY_FORMATTED_JSON) {
		t.Errorf("got %s, want %s", mr.Msg, helper.RESPONSE_BODY_MUST_NOT_BE_LARGER_THAN_1MB)
	}
}

func TestDecodeJsonBodyInvalidRequestBodyTooLarge(t *testing.T){
	// arrange
	req, resp := generateRequestInitialConditions(t, 30)

	// act
	err := helper.DecodeJSONBody(&resp, req, nil)

	// assert
	var mr *helper.MalformedRequest

	// should see an error due to nil destination interface passed as input to decode json body
	// method call
	if errors.As(err, &mr) == false || !strings.Contains(mr.Msg,helper.BADLY_FORMATTED_JSON) {
		t.Errorf("got %s, want %s", mr.Msg, helper.BADLY_FORMATTED_JSON)
	}
}
*/

func TestDecodeJsonBodyInvalidDestInterface(t *testing.T) {
	// arrange
	req, resp := generateRequestInitialConditions(t, 30)

	// act
	err := helper.DecodeJSONBody(&resp, req, &api.LoginUserRequest{})

	// assert
	var mr *helper.MalformedRequest
	if errors.As(err, &mr) == false || !strings.Contains(mr.Msg, helper.BADLY_FORMATTED_JSON) {
		t.Errorf("got %s, want %s", mr.Msg, helper.BADLY_FORMATTED_JSON)
	}
}

type FaultyJson struct {
	TestString string `json:"page sample"`
}

func TestDecodeJsonBodyInvalidJsonSyntax(t *testing.T) {
	// arrange
	testStr := FaultyJson{TestString: "json value for testing"}
	jsonStr, err := json.Marshal(&testStr)
	assert.Empty(t, err)

	req, err := http.NewRequest("POST", "/v1/test/endpoint", bytes.NewBuffer(jsonStr))
	assert.Empty(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.ResponseRecorder{}

	// act
	err = helper.DecodeJSONBody(&resp, req, &api.LoginUserRequest{})

	// assert
	var mr *helper.MalformedRequest
	errors.As(err, &mr)
	if mr.Status != http.StatusBadRequest {
		t.Errorf("got %d, want %d", mr.Status, http.StatusBadRequest)
	}
}

func TestDecodeJsonBodyInvalidEmptyRequest(t *testing.T) {
	req, resp := generateRequestInitialConditions(t, 0)

	// act
	err := helper.DecodeJSONBody(&resp, req, &FaultyJson{})

	// assert
	var mr *helper.MalformedRequest
	errors.As(err, &mr)
	if mr.Status != http.StatusBadRequest || mr.Msg != helper.REQUEST_BODY_MUST_NOT_BE_EMPTY {
		t.Errorf("got status %d, want status %d, got status %s, want status %s", mr.Status,
			http.StatusBadRequest, mr.Msg, helper.REQUEST_BODY_MUST_NOT_BE_EMPTY)
	}
}

func generateRequestInitialConditions(t *testing.T, byteArrayLength int64) (*http.Request,
	httptest.ResponseRecorder) {
	randBytes, err := generateRandomByteArray(byteArrayLength)
	assert.Empty(t, err)
	req, err := http.NewRequest("POST", "/v1/test/endpoint", bytes.NewBuffer(randBytes))
	assert.Empty(t, err)
	resp := httptest.ResponseRecorder{}
	return req, resp
}

func generateRandomByteArray(length int64) ([]byte, error) {
	byteToken := make([]byte, length)
	_, err := rand.Read(byteToken)
	return byteToken, err
}

// invalid cases
// TestDecodeJsonBodyInvalidEmptyContentType
// TestDecodeJsonBodyInvalidUnMarshalledType
// TestDecodeJsonBodyInvalidRequestUnknownField
// TestDecodeJsonBodyInvalidMoreThanOneJsonObject
