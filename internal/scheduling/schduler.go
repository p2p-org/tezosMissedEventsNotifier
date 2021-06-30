package scheduling

import (
	"log"
	"time"

	"tezos/missedEventsNotifier/pkg/api"
)

type Scheduler interface {
	ScheduleEndorsements()
	ScheduleBakings()
}

type scheduler struct {
	api api.API
}

func (s *scheduler) ScheduleEndorsements() {
	endorsements, err := s.api.GetEndorsements()
	log.Println("Get endorsements schedule")
	if err != nil {
		log.Println(err)
	}
	var lastPoint time.Time
	for _, endorsement := range endorsements {
		if endorsement.EstimatedTime.Year() == 1 {
			continue
		}
		point := endorsement.EstimatedTime.Add(time.Second)
		lastPoint = point
		time.AfterFunc(point.Sub(time.Now()), func() {
			api.CheckEndorsement(&endorsement, s.api)
		})
	}
	time.AfterFunc(lastPoint.Sub(time.Now()), func() {
		s.ScheduleEndorsements()
	})
}

func (s *scheduler) ScheduleBakings() {
	bakes, err := s.api.GetEndorsements()
	log.Println("Get bakes schedule")
	if err != nil {
		log.Println(err)
	}
	var lastPoint time.Time
	for _, bake := range bakes {
		if bake.EstimatedTime.Year() == 1 {
			continue
		}
		point := bake.EstimatedTime.Add(time.Second)
		lastPoint = point
		time.AfterFunc(point.Sub(time.Now()), func() {
			b, err := s.api.GetCurrentBlock()
			if err != nil {
				log.Println(err)
			}
			api.CheckBlock(b)
		})
	}
	time.AfterFunc(lastPoint.Sub(time.Now()), func() {
		s.ScheduleBakings()
	})
}

func NewScheduler(tzapi api.API) Scheduler {
	return &scheduler{api: tzapi}
}
