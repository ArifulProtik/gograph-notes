package utils

import "golang.org/x/crypto/bcrypt"

// HashBeforeSave Hashes the Password
func HashBeforeSave(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}
func VerifyPass(hashedpass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass))
}
