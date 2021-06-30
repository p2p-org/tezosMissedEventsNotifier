package api

import (
	"log"
	"time"
)

type Endorsement struct {
	Level         int       `json:"level"`
	Delegate      string    `json:"delegate"`
	Slots         []int     `json:"slots"`
	EstimatedTime time.Time `json:"estimated_time,omitempty"`
}

func CheckEndorsement(e *Endorsement, api API, slots []int) bool {
	m := make(map[int]bool)
	for _, item := range slots {
		m[item] = false
	}
	block, err := api.GetCurrentBlock()
	if err != nil {
		log.Println(err)
		log.Printf("Endorsement at %v failed", e.EstimatedTime)
		return false
	}

	for _, collection := range block.Operations {
		for _, operation := range collection {
			for _, item := range operation.Contents {
				if item.Kind == "endorsement" {
					if _, ok := m[item.Slot]; ok {
						m[item.Slot] = true
					}
				}
			}
		}
	}

	for _, value := range m {
		if !value {
			log.Printf("Endorsement at %v failed", e.EstimatedTime)
			return false
		}
	}
	log.Printf("Endorsement at %v fsuccessful", e.EstimatedTime)
	return true
}
