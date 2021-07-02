package api

import (
	"context"
	"log"
	"time"

	"blockwatch.cc/tzstats-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Endorsement corresponds to endorsement from the Tezos RPC API
type Endorsement struct {
	Level         int       `json:"level"`
	Delegate      string    `json:"delegate"`
	Slots         []int     `json:"slots"`
	EstimatedTime time.Time `json:"estimated_time,omitempty"`
}

var (
	endorsementsMissed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "endorsements_missed_total",
		Help: "Number of missed endorsements",
	})
)

// CheckEndorsement checks the endorsement and reports problem to Prom
func CheckEndorsement(e *Endorsement, tzapi API) bool {
	client := tzapi.(*api).client
	block, err := client.GetBlockHeight(context.TODO(), int64(e.Level), tzstats.NewBlockParams())
	if err != nil {
		log.Printf("Error while validating endorsement %d", e.Level)
		log.Println(err)
		return false
	}
	for _, op := range block.Ops {
		log.Printf("Discovered delegate %s", op.Delegate.String())
		if op.Type == 9 && op.Delegate.String() == e.Delegate {
			if op.IsSuccess {
				log.Printf("Endorsement %d is successful", e.Level)
				return true
			}
			endorsementsMissed.Inc()
			log.Printf("Endorsement %d is missed", e.Level)
			return false

		}
	}
	return false
}
