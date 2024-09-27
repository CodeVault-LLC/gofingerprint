package testing

import (
	"encoding/json"
	"net/http"

	"github.com/codevault-llc/gofingerprint"
)

// HelloHandler is a simple API handler that returns a message including fingerprint data.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fingerprintData := r.Header.Get("X-Fingerprint")

	var fingerprint gofingerprint.Fingerprint
	err := json.Unmarshal([]byte(fingerprintData), &fingerprint)
	if err != nil {
		http.Error(w, "Error parsing fingerprint", http.StatusInternalServerError)
		return
	}

	// Respond as JSON with the fingerprint data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fingerprint)
}

// SetupFakeAPI sets up the fake API with the fingerprinting middleware.
func SetupFakeAPI() http.Handler {
	mux := http.NewServeMux()

	// Apply the fingerprinting middleware to the HelloHandler
	mux.Handle("/hello", gofingerprint.FingerprintMiddleware(http.HandlerFunc(HelloHandler)))

	return mux
}
