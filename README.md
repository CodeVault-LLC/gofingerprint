# Go-Fingerprint

A fingerprinting library for golang applications with a middleware for http servers.

## Usage

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/codevault-llc/gofingerprint"
)

func main() {
  // Create a new http server
	mux := http.NewServeMux()

	// Apply the fingerprinting middleware to the HelloHandler
	mux.Handle("/hello", gofingerprint.FingerprintMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The Fingerprint adds the required headers to the response.
		// Here is how you can access the fingerprint data.
		fingerprintHeader := r.Header.Get("X-Fingerprint") // The header returned by the middleware

		var fingerprint gofingerprint.Fingerprint                      // The struct to hold the fingerprint data
		err := json.Unmarshal([]byte(fingerprintHeader), &fingerprint) // Parse the header into the struct
		if err != nil {
			http.Error(w, "Error parsing fingerprint", http.StatusInternalServerError)
			return
		}

		// Now you can access the fingerprint data
		fmt.Println(fingerprint.Hash) // The hash returned of all the calculations done by the middleware, save this to verify the fingerprint later.

		fmt.Fprintf(w, "Hello, World!")
	})))

	// Start the server
	http.ListenAndServe(":8080", mux)
}
```

## Understanding the Fingerprint

The fingerprint is a simple function that calculates a hash of the requests the user is sending. After making a correct hash, we then send it over to the developer for them to do whatever they want. The hash is often stored in a database at the API, and then verified how many times the hash has been seen and so fourth. Generally just a ID for the user.

When you implement the middleware, we are calculating some different things, such as if its a local ip, the user agent, the time, and the request method. We then hash all of these things together to create a unique hash for the user. This is then sent back to the user in the `X-Fingerprint` header. The value of the header is a stringified JSON object that is equal to the `Fingerprint` struct.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
