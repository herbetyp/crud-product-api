package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encoder(p string) string {
	str := sha256.Sum256([]byte(p))

	return fmt.Sprintf("%x", str)
}