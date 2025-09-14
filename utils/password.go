package utils

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	// Implement password hashing logic here (e.g., using bcrypt)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err // Placeholder, replace with actual hashed password
}