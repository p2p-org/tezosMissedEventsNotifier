package scheduling

import (
	"log"

	"github.com/jasonlvhit/gocron"

	"tezos/missedEventsNotifier/pkg/api"
)

type scheduler struct {
	api api.API
}

func (s *scheduler) ScheduleEndorsements() {
	endorsements, err := s.api.GetEndorsements()
	if err != nil {
		log.Println("Get endorsements schedule")
		log.Println(err)
	}
	for _, endorsement := range endorsements {
		if endorsement.EstimatedTime.Year() == 1 {
			continue
		}
		gocron.Every(1).From(&endorsement.EstimatedTime).Do(api.CheckEndorsement(&endorsement, s.api))
	}

}

func (s *scheduler) ScheduleBakings() {
	panic("implement me")
}

type Scheduler interface {
	ScheduleEndorsements()
	ScheduleBakings()
}
