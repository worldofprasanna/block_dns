package block_dns_test

import (
	"testing"
	"github.com/worldofprasanna/block_dns"
)

func TestBlacklistDomainsForFacebook(t *testing.T) {
	validator := block_dns.NewBlacklistDomain()
	result := validator.IsBlocked("www.facebook.com.")
	if !result {
		t.Errorf("Failed to block '%s'", "facebook")
	}
}

func TestBlacklistDomainsForGoogle(t *testing.T) {
	validator := block_dns.NewBlacklistDomain()
	result := validator.IsBlocked("www.google.com.")
	if result {
		t.Errorf("Failed to allow '%s'", "google")
	}
}