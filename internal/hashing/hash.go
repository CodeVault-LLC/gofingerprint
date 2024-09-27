package hashing

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
