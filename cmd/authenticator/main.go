package main

import (
	"fmt"
	"os"

	"github.com/hitecherik/hetzner-certbot-dns/pkg/certbot"
	"github.com/hitecherik/hetzner-certbot-dns/pkg/hetzner"
)

const ttl = 60

func main() {
	apiKey, ok := os.LookupEnv("HETZNER_API_KEY")
	if !ok {
		panic(fmt.Errorf("no hetzner API key provided"))
	}

	h := hetzner.New(apiKey)
	if err := h.FetchZones(); err != nil {
		panic(err)
	}

	domain := fmt.Sprintf("_acme-challenge.%v", certbot.Parameters.Domain)
	recordId, err := h.CreateRecord(domain, "TXT", ttl, certbot.Parameters.Validation)
	if err != nil {
		panic(err)
	}

	fmt.Println(recordId)
}
