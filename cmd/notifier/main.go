package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"tezos/missedEventsNotifier/internal/configs"
	"tezos/missedEventsNotifier/internal/scheduling"
	"tezos/missedEventsNotifier/pkg/api"
)

func main() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()
	config, err := configs.GetConfig("./config/config.yaml")
	if err != nil {
		log.Fatalln("Failed to read config")
	}
	tzApi := api.NewApi(config.Host, config.Delegate)
	scheduler := scheduling.NewScheduler(tzApi)
	scheduler.EndorsementsWg().Add(2)
	scheduler.ScheduleEndorsements()
	scheduler.BakingsWg().Add(2)
	scheduler.ScheduleBakings()
	scheduler.BakingsWg().Wait()
	scheduler.EndorsementsWg().Wait()
}
