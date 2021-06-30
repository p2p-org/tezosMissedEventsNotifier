package api

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	endorsementResponse = "[{\"level\":1540093,\"delegate\":\"tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb\"," +
		"\"slots\":[14,31],\"estimated_time\":\"2021-07-02T04:41:18Z\"},{\"level\":1540095,\"delegate\"" +
		":\"tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb\",\"slots\":[29],\"estimated_time\":\"2021-07-02T04:43:18Z\"}]"
)

var delegate = "tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb"
var api1 = NewApi("https://mainnet-tezos.giganode.io/", delegate, 375)
var api2 = NewApi("https://rpc.tzbeta.net", delegate, 375)

func Test_api_GetCurrentBlock(t *testing.T) {
	b1, err := api1.GetCurrentBlock()
	assert.NoError(t, err)
	b2, err := api2.GetCurrentBlock()
	assert.NoError(t, err)
	assert.Equal(t, b1.Hash, b2.Hash)
}

func Test_api_Parsing(t *testing.T) {
	var endorsements []Endorsement
	err := json.NewDecoder(strings.NewReader(endorsementResponse)).Decode(&endorsements)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(endorsements))
}

func Test_api_GetEndorsements(t *testing.T) {
	b1, err := api1.GetEndorsements()
	assert.NoError(t, err)
	assert.NotZero(t, len(b1))
}
