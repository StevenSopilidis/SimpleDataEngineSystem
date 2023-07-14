package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashes string using SHA256 and returns hex representation
func HashStringToSHA256(key []byte) string {
	hash := sha256.Sum256([]byte(key))
	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash[:])
}
