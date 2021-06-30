package api

import (
	"encoding/json"
	"net/http"
)

const (
	currentBlockSuffix = "/chains/main/blocks/head"
)

type API interface {
	GetCurrentBlock() Block
}

type api struct {
	baseURl string
}

func (a *api) GetCurrentBlock() (b *Block, err error) {
	resp, err := http.Get(a.baseURl + currentBlockSuffix)
	if err != nil {
		return nil, err
	}
	b = new(Block)
	err = json.NewDecoder(resp.Body).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func NewApi(baseURl string) *api {
	return &api{baseURl: baseURl}
}
