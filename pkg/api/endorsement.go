package api

import (
	"time"

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
