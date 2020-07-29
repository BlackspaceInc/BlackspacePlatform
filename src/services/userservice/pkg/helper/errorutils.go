package helper

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/BlackspaceInc/Backend/user-management-service/pkg/authentication"
)

// ProcessAggregatedErrors processes aggregated errors
func ProcessAggregatedErrors(w http.ResponseWriter, aggregatedErr *authentication.AggregatedError) bool {
	if aggregatedErr.Error != nil {
		err := aggregatedErr.Error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	} else if aggregatedErr.AuthErrorMsg != nil {
		jsonStr, err := json.Marshal(&aggregatedErr.AuthErrorMsg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return true
		}

		http.Error(w, string(jsonStr), http.StatusBadRequest)
		return true
	}
	return false
}

// ProcessAggregatedErrorsInTest processes aggregated errors while running and writing unit test cases
func ProcessAggregatedErrorsInTest(t *testing.T, aggregatedErr *authentication.AggregatedError) bool {
	if aggregatedErr.Error != nil {
		err := aggregatedErr.Error
		t.Errorf("Expected empty error field to be returned'. Got '%v'", err)
		return true
	} else if aggregatedErr.AuthErrorMsg != nil {
		jsonStr, err := json.Marshal(&aggregatedErr.AuthErrorMsg)
		if err != nil {
			t.Errorf("Expected empty error field to be returned'. Got '%v'", err)
			return true
		}

		t.Errorf("Expected empty error field to be returned'. Got '%v'", string(jsonStr))
		return true
	}
	return false
}
