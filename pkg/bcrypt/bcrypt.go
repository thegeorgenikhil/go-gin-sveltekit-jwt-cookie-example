package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// VerifyPassword hashes the password entered by the user and compares it with the hashed password in the database
func VerifyPassword(hashedPassword string, passwordEntered string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(passwordEntered))
	return err == nil
}
