package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	currentBlockSuffix = "/chains/main/blocks/head"
	listSuffixFormat   = "/chains/main/blocks/head/helpers/%s_rights"
	endorsing          = "endorsing"
	baking             = "baking"
)

type API interface {
	GetCurrentBlock() (b *Block, err error)
	GetEndorsements() ([]Endorsement, error)
	GetBakes() ([]Bake, error)
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

func (a *api) getList(what string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", a.baseURl+fmt.Sprintf(listSuffixFormat, what), nil)
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
	return resp.Body, err
}

func (a *api) GetEndorsements() (endorsements []Endorsement, err error) {
	body, err := a.getList(endorsing)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(body).Decode(&endorsements)
	if err != nil {
		return nil, err
	}
	return endorsements, nil
}

func (a *api) GetBakes() (bakes []Bake, err error) {
	body, err := a.getList(endorsing)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(body).Decode(&bakes)
	if err != nil {
		return nil, err
	}
	return bakes, nil
}

func NewApi(baseURl, delegate string, cycle int) API {
	return &api{baseURl: baseURl, delegate: delegate}
}
