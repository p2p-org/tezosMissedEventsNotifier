package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_api_GetCurrentBlock(t *testing.T) {
	api1 := NewApi("https://mainnet.smartpy.io")
	api2 := NewApi("https://rpc.tzbeta.net")

	b1, err := api1.GetCurrentBlock()
	assert.NoError(t, err)
	b2, err := api2.GetCurrentBlock()
	assert.NoError(t, err)
	assert.Equal(t, b1.Hash, b2.Hash)
}
