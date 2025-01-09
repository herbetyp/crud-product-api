package services

import (
	"crypto/sha512"
	"fmt"
)

func SHA512Encoder(s string) string {
	stringEncoded := sha512.Sum512([]byte(s))

	return fmt.Sprintf("%x", stringEncoded)
}
