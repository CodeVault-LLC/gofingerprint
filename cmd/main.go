package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codevault-llc/gofingerprint"
)

func main() {
	// Create a new http server
	mux := http.NewServeMux()

	// Apply the fingerprinting middleware to the HelloHandler
	mux.Handle("/", gofingerprint.FingerprintMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The Fingerprint adds the required headers to the response.
		// Here is how you can access the fingerprint data.
		fingerprintHeader := r.Header.Get("X-Fingerprint") // The header returned by the middleware

		var fingerprint gofingerprint.Fingerprint                      // The struct to hold the fingerprint data
		err := json.Unmarshal([]byte(fingerprintHeader), &fingerprint) // Parse the header into the struct
		if err != nil {
			http.Error(w, "Error parsing fingerprint", http.StatusInternalServerError)
			return
		}

		// Respond as JSON with the fingerprint data
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fingerprint)
	})))

	// Start the server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
