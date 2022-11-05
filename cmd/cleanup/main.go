package main

import (
	"fmt"
	"os"

	"github.com/hitecherik/hetzner-certbot-dns/pkg/certbot"
	"github.com/hitecherik/hetzner-certbot-dns/pkg/hetzner"
)

func main() {
	apiKey, ok := os.LookupEnv("HETZNER_API_KEY")
	if !ok {
		panic(fmt.Errorf("no hetzner API key provided"))
	}

	h := hetzner.New(apiKey)

	if err := h.DeleteRecord(certbot.Parameters.AuthOutput); err != nil {
		panic(err)
	}
}
