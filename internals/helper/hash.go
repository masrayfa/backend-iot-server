package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) (string, error ){
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}

	hashedBytes := h.Sum(nil)
	hashedPassword := hex.EncodeToString(hashedBytes)

	return hashedPassword, nil 
}