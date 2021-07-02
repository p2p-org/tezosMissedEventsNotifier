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

type API interface {
	GetCurrentBlock() (b *Block, err error)
	GetEndorsements() ([]Endorsement, error)
	GetBakes() ([]Bake, error)
	// GetBlock(estTime time.Time, level int) (b *Block, err error)
}

type api struct {
	baseURl  string
	delegate string
	cycle    int
	client   *tzstats.Client
}

func (a *api) GetBlockByHash(hash string) (b *Block, err error) {
	url := "https://" + path.Join(a.baseURl, fmt.Sprintf(blockByHashSuffix, hash))
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

//func (a *api) getExpectedBlock(estTime time.Time, d int) (b *Block, err error) {
//	suffix := fmt.Sprintf(blockByLevelSuffix, d, estTime.Format(tezosTimeFormat))
//	url := "https://" + path.Join(a.baseURl, suffix)
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	b = new(Block)
//	err = json.NewDecoder(resp.Body).Decode(b)
//	if err != nil {
//		return nil, err
//	}
//	return b, nil
//}
//
//func abs(x int) int {
//	if x > 0 {
//		return x
//	}
//	return -x
//}
//
//func (a *api) GetBlock(level int) (b *Block, err error) {
//	block, err := a.client.GetBlockHeight(context.TODO(), int64(level), tzstats.NewBlockParams())
//	if err != nil {
//		return nil, err
//	}
//	block.Ops
//	return b, nil
//}

func (a *api) GetCurrentBlock() (b *Block, err error) {
	resp, err := http.Get("https://" + a.baseURl + currentBlockSuffix)
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
	req, err := http.NewRequest("GET", "https://"+a.baseURl+fmt.Sprintf(listSuffixFormat, what), nil)
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

func NewApi(baseURl, delegate string) API {
	c, _ := tzstats.NewClient("https://api.edo.tzstats.com", nil)
	return &api{baseURl: baseURl, delegate: delegate, client: c}
}
