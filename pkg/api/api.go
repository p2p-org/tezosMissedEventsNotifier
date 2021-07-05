package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"blockwatch.cc/tzstats-go"
)

const (
	blockByLevelSuffix = "/chains/main/blocks/head~%d?min_date=\"%s\""
	currentBlockSuffix = "/chains/main/blocks/head"
	blockByHashSuffix  = "/chains/main/blocks/%s"
	listSuffixFormat   = "/chains/main/blocks/head/helpers/%s_rights"

	endorsing = "endorsing"
	baking    = "baking"

	tezosTimeFormat = "2006-01-02T15:04:05Z"
)

// API provides interface to work with Tezos RPC API
type API interface {
	GetCurrentBlock() (b *Block, err error)
	GetEndorsements() ([]Endorsement, error)
	GetBakes() ([]Bake, error)
	GetBlockByHeight(height int) (b *Block, err error)
	// GetBlock(estTime time.Time, level int) (b *Block, err error)
}

type api struct {
	baseURL  string
	delegate string
	cycle    int
	client   *tzstats.Client
}

func (a *api) GetBlockByHash(hash string) (b *Block, err error) {
	url := "https://" + path.Join(a.baseURL, fmt.Sprintf(blockByHashSuffix, hash))
	log.Println(url)
	resp, err := http.Get(url)
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

func (a *api) GetBlockByHeight(height int) (b *Block, err error) {
	return a.GetBlockByHash(strconv.Itoa(height))
}

func (a *api) GetCurrentBlock() (b *Block, err error) {
	resp, err := http.Get("https://" + a.baseURL + currentBlockSuffix)
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
	b, err := a.GetCurrentBlock()
	if err != nil {
		return nil, err
	}
	a.cycle = b.Metadata.LevelInfo.Cycle
	log.Printf("Current cycle %d\n", a.cycle)
	req, err := http.NewRequest("GET", "https://"+a.baseURL+fmt.Sprintf(listSuffixFormat, what), nil)
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
	resp, err := a.getList(baking)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&bakes)
	if err != nil {
		return nil, err
	}
	return bakes, nil
}

// NewAPI is an API constructor
func NewAPI(baseURL, delegate string) API {
	c, _ := tzstats.NewClient("https://api.edo.tzstats.com", nil)
	return &api{baseURL: baseURL, delegate: delegate, client: c}
}
