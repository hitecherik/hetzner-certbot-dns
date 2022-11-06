package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

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

	for !checkTXT(domain, certbot.Parameters.Validation) {
		time.Sleep(ttl * time.Second)
	}
}

func checkTXT(domain string, expected string) bool {
	results, err := net.LookupTXT(domain)

	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			return false
		} else {
			panic(err)
		}
	}

	for _, result := range results {
		if result == expected {
			return true
		}
	}

	return false
}
