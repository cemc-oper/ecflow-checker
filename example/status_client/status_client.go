package main

import (
	"github.com/perillaroc/ecflow-checker"
	"log"
)

func main() {
	target := "10.40.140.18:31181"

	client := ecflow_checker.EcflowClient{
		Target: target,
	}

	err := client.Connect()

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	defer client.Close()

	client.CollectStatusRecords(
		"nwp_xp",
		"nwpc_pd",
		"login_b01",
		"31071")
}
