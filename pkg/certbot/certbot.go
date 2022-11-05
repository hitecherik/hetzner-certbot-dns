package certbot

import (
	"os"
	"strings"
)

type Authentication struct {
	Domain     string
	Validation string
	AuthOutput string
}

var Parameters Authentication

func init() {
	domain, _ := os.LookupEnv("CERTBOT_DOMAIN")
	validation, _ := os.LookupEnv("CERTBOT_VALIDATION")
	authOutput, _ := os.LookupEnv("CERTBOT_AUTH_OUTPUT")

	Parameters = Authentication{
		Domain:     domain,
		Validation: validation,
		AuthOutput: strings.TrimSpace(authOutput),
	}
}
