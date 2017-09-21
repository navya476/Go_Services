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
	port := os.Getenv("DRP_CF_HTTP_PORT")
	//url := os.Getenv("DRP_CF_HTTP_ADDR")
	// Verify the Provider - Latest Published Pacts for any known consumers
	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%s", port),
		BrokerURL:                  "http://54.207.30.143:8081",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		Verbose:                    true,
	})

	if err != nil {
		t.Fatal("Error:", err)
	}

}
