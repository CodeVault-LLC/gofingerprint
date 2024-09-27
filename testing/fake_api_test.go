package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codevault-llc/gofingerprint"
)

func TestHelloHandler(t *testing.T) {
	api := SetupFakeAPI()

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.Header.Set("User-Agent", "TestAgent")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Serve the request using our API
	api.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// JSON response and decode it.
	fingerprint := json.NewDecoder(rr.Body)
	var data gofingerprint.Fingerprint
	err := fingerprint.Decode(&data)
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	fmt.Printf("Fingerprint: %+v\n", data.Hash)
}
