package api

import (
	"context"
	"log"
	"time"

	"blockwatch.cc/tzstats-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

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
			} else {
				endorsementsMissed.Inc()
				log.Printf("Endorsement %d is missed", e.Level)
				return false
			}
		}
	}
	return false
	//m := make(map[int]bool)
	//for _, item := range e.Slots {
	//	m[item] = false
	//}
	//
	//block, err := tzapi.(*api).tzkt.GetBlock(uint64(b.Level))
	//if err != nil {
	//	log.Println(err)
	//	return false
	//}

	//block, err := api.GetCurrentBlock()
	//if err != nil {
	//	log.Println(err)
	//	log.Printf("Endorsement at %v failed", e.EstimatedTime)
	//	return false
	//}
	//
	//for _, collection := range block.Operations {
	//	for _, operation := range collection {
	//		for _, item := range operation.Contents {
	//			if item.Kind == "endorsement" {
	//				if _, ok := m[item.Slot]; ok {
	//					m[item.Slot] = true
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//for _, value := range m {
	//	if !value {
	//		log.Printf("Endorsement at level %d failed", e.Level)
	//		endorsementsMissed.Inc()
	//		return false
	//	}
	//}
	//log.Printf("Endorsement at level %d fsuccessful", e.Level)
	//return true
	return true
}
