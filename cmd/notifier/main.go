package main

import (
	"fmt"
	"os"
	"tezos/missedEventsNotifier/internal/configs"

	client "github.com/goat-systems/go-tezos/v4/rpc"
)

func main() {
	apiLink, err := configs.GetApiLink("../config/config.yaml")
	if err != nil {
		fmt.Printf("failed to locate config %v", err)
	}
	rpc, err := client.New(apiLink)
	if err != nil {
		fmt.Printf("failed tp connect to network: %v", err)
	}

	resp, cycle, err := rpc.Cycle(50)
	if err != nil {
		fmt.Printf("failed to get (%s) cycle: %s\n", resp.Status(), err.Error())
		os.Exit(1)
	}
	fmt.Println(cycle)
}
