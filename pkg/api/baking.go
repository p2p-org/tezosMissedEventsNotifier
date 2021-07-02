package api

import (
	"context"
	"log"
	"time"

	"blockwatch.cc/tzstats-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Bake is a struct corresponding to bake opertation from the Tezos RPC API
type Bake struct {
	Level         int       `json:"level"`
	Delegate      string    `json:"delegate"`
	Priority      int       `json:"priority"`
	EstimatedTime time.Time `json:"estimated_time,omitempty"`
}

var (
	bakesMissed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bakes_missed_total",
		Help: "Number of missed bakes",
	})
)

// CheckBake determines if bake was not missed and reports miss to Prom
func CheckBake(tzapi API, b *Bake) bool {
	log.Printf("Checking bake for level %d", b.Level)
	block, err := tzapi.(*api).client.GetBlockHeight(context.TODO(), int64(b.Level), tzstats.NewBlockParams())
	if err != nil {
		log.Println(err)
		return false
	}
	if block.Priority > 0 {
		bakesMissed.Inc()
		log.Printf("bake missed for block %s\n", block.Hash)
		return false
	}
	log.Printf("Success with block %s\n", block.Hash)
	return true
}
