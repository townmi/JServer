package utils

import "golang.org/x/crypto/scrypt"

var salt string

// Encode s
func Encode(password string) (string, error) {
	salt = "test"
	s, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
