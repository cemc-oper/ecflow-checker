package main

import (
	"github.com/perillaroc/ecflow-client-go"
	"log"
)

func main() {
	target := "10.40.140.18:31181"

	client := ecflow_client_go.EcflowClient{
		Target: target,
	}

	err := client.Connect()

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	defer client.Close()

	client.CollectNode(
		"nwp_xp",
		"nwpc_pd",
		"login_b01",
		"31071",
		"/meso_post/00/initial")
}
