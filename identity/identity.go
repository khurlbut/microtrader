package identity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateRandomID generates a random, unique string ID of specified length.
// Length here refers to the length of the final string, not the number of bytes.
func GenerateRandomID(length int) (string, error) {
	byteLength := length / 2
	bytes := make([]byte, byteLength)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return hex.EncodeToString(bytes), nil
}
