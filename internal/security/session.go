package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// NewSessionToken generates a cryptographically secure random session token.
func NewSessionToken() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// HashSessionToken returns the SHA-256 hash of a session token.
func HashSessionToken(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}
