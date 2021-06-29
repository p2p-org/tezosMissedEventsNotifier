package main

import (
	"fmt"
	"os"

	client "github.com/goat-systems/go-tezos/v4/rpc"
)

func main() {
	rpc, err := client.New("http://127.0.0.1:8732")
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
