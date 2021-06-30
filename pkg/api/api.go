package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	currentBlockSuffix = "/chains/main/blocks/head"
	endorsementsSuffix = "/chains/main/blocks/head/helpers/endorsing_rights"
)

type API interface {
	GetCurrentBlock() (b *Block, err error)
	GetEndorsements() ([]Endorsement, error)
}

type api struct {
	baseURl  string
	delegate string
	cycle    int
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

func (a *api) GetEndorsements() (endorsements []Endorsement, err error) {
	req, err := http.NewRequest("GET", a.baseURl+endorsementsSuffix, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("cycle", strconv.Itoa(a.cycle))
	q.Add("delegate", a.delegate)
	req.URL.RawQuery = q.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(endorsements)
	if err != nil {
		return nil, err
	}
	return endorsements, nil
}

func NewApi(baseURl, delegate string, cycle int) API {
	return &api{baseURl: baseURl, delegate: delegate}
}
