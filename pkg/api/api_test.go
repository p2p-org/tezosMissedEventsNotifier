package api

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	endorsementResponse = "[{\"level\":1540093,\"delegate\":\"tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb\"," +
		"\"slots\":[14,31]},{\"level\":1540095,\"delegate\"" +
		":\"tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb\",\"slots\":[29],\"estimated_time\":\"2021-07-02T04:43:18Z\"}]"
)

var delegate = "tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb"
var api1 = NewAPI("mainnet-tezos.giganode.io", delegate)

func Test_api_GetCurrentBlock(t *testing.T) {
	b1, err := api1.GetCurrentBlock()
	assert.NoError(t, err)
	b2, err := api1.GetCurrentBlock()
	assert.NoError(t, err)
	assert.Equal(t, b1.Hash, b2.Hash)
}

func Test_api_Parsing(t *testing.T) {
	var endorsements []Endorsement
	err := json.NewDecoder(strings.NewReader(endorsementResponse)).Decode(&endorsements)
	assert.NoError(t, err)
	log.Println(endorsements[0].EstimatedTime)
	assert.Equal(t, 2, len(endorsements))
}

func Test_api_GetEndorsements(t *testing.T) {
	b1, err := api1.GetEndorsements()
	assert.NoError(t, err)
	assert.NotZero(t, len(b1))
}

func Test_api_GetBlockByHeight(t *testing.T) {
	b, err := api1.GetBlockByHeight(1544010)
	assert.NoError(t, err)
	assert.Equal(t, 1544010, b.Metadata.Level)
}
