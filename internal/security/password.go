package security

import "golang.org/x/crypto/bcrypt"

// Hash hashes a plain-text password using bcrypt.
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare compares a bcrypt hash with a plain-text password.
// It returns nil if they match.
func ComparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
