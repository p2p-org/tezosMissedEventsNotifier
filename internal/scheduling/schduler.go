package scheduling

import (
	"log"
	"sync"
	"time"

	"tezos/missedEventsNotifier/pkg/api"
)

type Scheduler interface {
	ScheduleEndorsements()
	ScheduleBakings()
	EndorsementsWg() *sync.WaitGroup
	BakingsWg() *sync.WaitGroup
}

type scheduler struct {
	api            api.API
	endorsementsWg sync.WaitGroup
	bakingsWg      sync.WaitGroup
}

func (s *scheduler) EndorsementsWg() *sync.WaitGroup {
	return &s.endorsementsWg
}

func (s *scheduler) BakingsWg() *sync.WaitGroup {
	return &s.bakingsWg
}

func (s *scheduler) ScheduleEndorsements() {
	log.Println("Get endorsements schedule")
	endorsements, err := s.api.GetEndorsements()
	log.Println("Got endorsements schedule")
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
		s.endorsementsWg.Add(1)
		go func() {
			time.AfterFunc(point.Sub(time.Now()), func() {
				api.CheckEndorsement(&endorsement, s.api)
			})
			s.endorsementsWg.Done()
		}()
	}
	go func() {
		time.AfterFunc(lastPoint.Sub(time.Now()), func() {
			s.ScheduleEndorsements()
		})
	}()
}

func (s *scheduler) ScheduleBakings() {
	log.Println("Get bakes schedule")
	bakes, err := s.api.GetBakes()
	log.Println("Got bakes schedule")
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
		s.bakingsWg.Add(1)
		go func() {
			time.AfterFunc(point.Sub(time.Now()), func() {
				b, err := s.api.GetCurrentBlock()
				if err != nil {
					log.Println(err)
				}
				api.CheckBlock(b)
				s.bakingsWg.Done()
			})
		}()
	}
	go func() {
		time.AfterFunc(lastPoint.Sub(time.Now()), func() {
			s.ScheduleBakings()
		})
	}()
}

func NewScheduler(tzapi api.API) Scheduler {
	return &scheduler{api: tzapi}
}
