package stringHelper

import (
	"net/mail"
	"strings"

	"golang.org/x/net/idna"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidHostname(hostname string) bool {
	hostname = strings.ReplaceAll(hostname, "http://", "")
	hostname = strings.ReplaceAll(hostname, "https://", "")

	_, err := idna.Lookup.ToASCII(hostname)
	return err == nil
}
