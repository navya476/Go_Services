package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestMain(t *testing.T) {
	var url string
	if os.Getenv("DRP_CF_HTTP_ADDR") != "" {
		url = os.Getenv("DRP_CF_HTTP_ADDR")
	} else {
		url = "localhost"
	}

	// Create Pact, connecting to local Daemon
	// Ensure the port matches the daemon port!
	pact := dsl.Pact{
		Port:     6666,
		Consumer: "Go_Consumer",
		Provider: "Go_Provider",
	}
	// Shuts down Mock Service when done
	//defer pact.Teardown()
	//p := dsl.Publisher{}
	//err := p.Publish(types.PublishRequest{
	//	PactURLs:        []string{"./pacts/go_consumer-go_provider.json"},
	//	PactBroker:      "http://54.207.30.143:8081",
	//	ConsumerVersion: "1.0.1",
	//	Tags:            []string{"latest", "dev"},
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	// Set up our interactions. Note we have multiple in this test case!
	pact.
		AddInteraction().
		Given("User exists").
		UponReceiving("A request").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   "/user",
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body:   `{"id": 1,"name": "john"}`,
		})

	// Run the test and verify the interactions.
	if err := pact.Verify(func() error {
		u := fmt.Sprintf("http://%s:%d/user", url, pact.Server.Port)
		req, err := http.NewRequest("GET", u, nil)
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}
		if _, err = http.DefaultClient.Do(req); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	pact.WritePact()

	// Shuts down Mock Service when done
	defer pact.Teardown()
	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{"./pacts/go_consumer-go_provider.json"},
		PactBroker:      "http://54.207.30.143:8081",
		ConsumerVersion: "1.0.1",
		Tags:            []string{"latest", "dev"},
	})
	if err != nil {
		fmt.Println(err)
	}

}
