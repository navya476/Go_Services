package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func createPact() dsl.Pact {
	return dsl.Pact{
		Port:     6666,
		Consumer: "Go_Consumer",
		Provider: "Go_Provider",
	}
}
func TestPact_Provider(t *testing.T) {
	pact := createPact()
	var port string
	var url string
	if os.Getenv("DRP_CF_HTTP_PORT") != "" {
		port = os.Getenv("DRP_CF_HTTP_PORT")
	} else {
		port = "8085"
	}
	if os.Getenv("DRP_CF_HTTP_ADDR") != "" {
		url = os.Getenv("DRP_CF_HTTP_ADDR")
	} else {
		url = "localhost"
	}
	//url := os.Getenv("DRP_CF_HTTP_ADDR")
	// Verify the Provider - Latest Published Pacts for any known consumers
	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s:%s", url, port),
		BrokerURL:                  "http://54.207.30.143:8081",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		Verbose:                    true,
	})

	if err != nil {
		t.Fatal("Error:", err)
	}

}
