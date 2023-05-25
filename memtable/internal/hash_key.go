package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashes string using SHA256 and returns hex representation
func HashStringToSHA256(key []byte) string {
	hash := sha256.New()
	hash.Write(key)
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
