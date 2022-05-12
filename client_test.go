package ynrcc

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var client *Client

func init() {
	godotenv.Load()
}

func TestNewClient(t *testing.T) {
	var err error
	client, err = NewClient(Options{
		GatewayURL: os.Getenv("GATEWAY_URL"),
		Cert:       os.Getenv("CERT"),
		PriKey:     os.Getenv("PRI_KEY"),
		BankPubKey: os.Getenv("BANK_PUB_KEY"),
		MerID:      os.Getenv("MER_ID"),
		TemID:      os.Getenv("TEM_ID"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestClient_Execute(t *testing.T) {
	req := &QueryTrxStateRequest{}
	res := &QueryTrxStateResponse{}
	err := client.Execute(req, res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Code)
	assert.NotEqual(t, "100036", res.Code)
}
