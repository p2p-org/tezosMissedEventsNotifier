package api

import (
	"log"
	"strings"
	"time"

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

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CheckEndorsement checks the endorsement and reports problem to Prom
func CheckEndorsement(e *Endorsement, tzapi API) bool {
	b, err := tzapi.GetBlockByHeight(e.Level)
	for err != nil {
		b, err = tzapi.GetBlockByHeight(e.Level)
	}
	for _, ops := range b.Operations {
		for _, op := range ops {
			for _, cont := range op.Contents {
				if strings.HasPrefix(cont.Kind, "endorsement") && cont.Metadata.Delegate == e.Delegate && testEq(e.Slots, cont.Metadata.Slots) {
					log.Printf("Success lvl %d", e.Level)
					return true
				}
			}
		}
	}
	log.Printf("Missed endorsement %d", e.Level)
	endorsementsMissed.Inc()
	return false
}
