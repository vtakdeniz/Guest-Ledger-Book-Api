//go:build provider
// +build provider

package main

import (
	"fmt"
	"guestLedgerBookApi/comments"
	"guestLedgerBookApi/database"
	"log"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

type Settings struct {
	Host                       string
	ConsumerName               string
	ProviderName               string
	PactURL                    string
	PublishVerificationResults bool
	FailIfNoPactsFound         bool
	DisableToolValidityCheck   bool
	BrokerBaseURL              string
	BrokerToken                string
	ProviderVersion            string
	PactFileWriteMode          string
}

func (s *Settings) InitSettings() {
	s.Host = "127.0.0.1"
	s.ConsumerName = "guestLedgerBookClient"
	s.ProviderName = "guestLedgerBookApi"
	s.PactURL = "https://mediterranean.pactflow.io/pacts/provider/ShoppingCartApi/consumer/ShoppingCartClient/version/1"
	s.PublishVerificationResults = true
	s.FailIfNoPactsFound = true
	s.DisableToolValidityCheck = true
	s.BrokerBaseURL = "https://mediterranean.pactflow.io"
	s.BrokerToken = "C05m5gQduXO-0fFGPYN6mw"
	s.ProviderVersion = "1.0.0"
	s.PactFileWriteMode = "merge"
}

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	dbConfig := database.DbConfig{
		DbType: "sqlite3",
		Db:     "fakeprovider.db",
	}
	seed := []comments.Comment{
		{
			Email:   "test@gmail.com",
			Content: "test",
		}, {
			Email:   "test2@gmail.com",
			Content: "test2",
		},
	}
	go StartServerWithSeed(port, dbConfig, seed)

	settings := Settings{}
	settings.InitSettings()

	pact := dsl.Pact{
		Consumer:                 settings.ConsumerName,
		Provider:                 settings.ProviderName,
		Host:                     settings.Host,
		DisableToolValidityCheck: settings.DisableToolValidityCheck,
	}
	verifyRequest := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://%s:%d", settings.Host, port),
		PactURLs:                   []string{settings.PactURL},
		BrokerURL:                  settings.BrokerBaseURL,
		BrokerToken:                settings.BrokerToken,
		FailIfNoPactsFound:         settings.FailIfNoPactsFound,
		PublishVerificationResults: settings.PublishVerificationResults,
		ProviderVersion:            settings.ProviderVersion,
		StateHandlers: map[string]types.StateHandler{
			"i have a list of comments": func() error {
				return nil
			},
			"i will delete comments": func() error {
				return nil
			},
			"i will post a comment": func() error {
				return nil
			},
		},
	}

	_, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove("fakeprovider.db")
	if err != nil {
		log.Fatal(err)
	}
}
