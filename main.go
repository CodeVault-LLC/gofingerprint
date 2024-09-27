package gofingerprint

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codevault-llc/gofingerprint/internal/hashing"
	"github.com/codevault-llc/gofingerprint/internal/modules"
)

// Fingerprint represents all the data we want to collect.
type Fingerprint struct {
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	IsLocalIP bool   `json:"is_local_ip"`

	Hash string `json:"hash"`
}

// NewFingerprint creates a new Fingerprint object based on the request.
func NewFingerprint(req *http.Request) *Fingerprint {
	// Get the IP
	ipAddress := modules.GetIP(req)

	// Get the User-Agent from the request header
	userAgent := req.UserAgent()
	hashed := hashing.SHA256(ipAddress.IP + userAgent + strconv.FormatBool(ipAddress.IsLocal))

	// Populate the fingerprint struct
	return &Fingerprint{
		IPAddress: ipAddress.IP,
		IsLocalIP: ipAddress.IsLocal,
		UserAgent: userAgent,

		Hash: hashed,
	}
}

// Middleware that adds a fingerprint to the request.
func FingerprintMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create the fingerprint
		fingerprint := NewFingerprint(r)

		// Serialize the fingerprint to JSON
		fingerprintJSON, err := json.Marshal(fingerprint)
		if err != nil {
			http.Error(w, "Error generating fingerprint", http.StatusInternalServerError)
			return
		}

		// Set the serialized fingerprint in the header
		r.Header.Set("X-Fingerprint", string(fingerprintJSON))

		// Continue to the next handler
		next.ServeHTTP(w, r)
	})
}
