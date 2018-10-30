package utils

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

var salt string

// Encode s
func Encode(password string) (string, error) {
	salt = "test"
	fmt.Println(password)
	s, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(s), nil
}
