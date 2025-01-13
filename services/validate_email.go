package services

import "net/mail"

func ValidateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)

    return err == nil
}
