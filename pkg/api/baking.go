package api

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

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

func CheckBake(tzapi API, b *Bake) bool {
	block, err := tzapi.(*api).tzkt.GetBlock(uint64(b.Level))
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
