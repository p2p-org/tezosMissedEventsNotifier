package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var delegate = "tz1P2Po7YM526ughEsRbY4oR9zaUPDZjxFrb"
var api1 = NewApi("https://mainnet.smartpy.io", delegate, 375)
var api2 = NewApi("https://rpc.tzbeta.net", delegate, 375)

func Test_api_GetCurrentBlock(t *testing.T) {
	b1, err := api1.GetCurrentBlock()
	assert.NoError(t, err)
	b2, err := api2.GetCurrentBlock()
	assert.NoError(t, err)
	assert.Equal(t, b1.Hash, b2.Hash)
}

func Test_api_GetEndorsements(t *testing.T) {
	b1, err := api1.GetEndorsements()
	assert.NoError(t, err)
	b2, err := api2.GetEndorsements()
	assert.NoError(t, err)
	assert.Equal(t, b1, b2)
}
