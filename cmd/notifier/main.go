package main

import (
	"log"
	"strconv"

	"tezos/missedEventsNotifier/internal/configs"
	"tezos/missedEventsNotifier/pkg/api"
)

func main() {
	config, err := configs.GetConfig("./config/config.yaml")
	if err != nil {
		log.Fatalln("Failed to read config")
	}
	cycle, err := strconv.Atoi(config.Cycle)
	if err != nil {
		log.Fatalln(err)
	}
	tzApi := api.NewApi(config.Host, config.Delegate, cycle)

}
