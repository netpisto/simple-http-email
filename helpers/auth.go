package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}