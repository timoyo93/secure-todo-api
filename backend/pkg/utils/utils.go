package utils

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// HashPassword hashes a password and returns the hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a given hash with the given password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// TokenGenerator generates a random token with a length of 32
func TokenGenerator() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return strings.ToUpper(fmt.Sprintf("%x", b))
}
