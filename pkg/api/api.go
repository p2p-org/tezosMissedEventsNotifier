package api

import (
	"encoding/json"
	"fmt"
	"log"
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

func (a *api) getList(what string) (*http.Response, error) {
	req, err := http.NewRequest("GET", a.baseURl+fmt.Sprintf(listSuffixFormat, what), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("cycle", strconv.Itoa(a.cycle))
	q.Add("delegate", a.delegate)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (a *api) GetEndorsements() (endorsements []Endorsement, err error) {
	resp, err := a.getList(endorsing)
	if err != nil {
		return nil, err
	}
	//b, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(b))
	err = json.NewDecoder(resp.Body).Decode(&endorsements)
	if err != nil {
		return nil, err
	}
	return endorsements, nil
}

func (a *api) GetBakes() (bakes []Bake, err error) {
	resp, err := a.getList(endorsing)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&bakes)
	if err != nil {
		return nil, err
	}
	return bakes, nil
}

func NewApi(baseURl, delegate string, cycle int) API {
	return &api{baseURl: baseURl, delegate: delegate, cycle: cycle}
}
