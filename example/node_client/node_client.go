package main

import (
	"fmt"
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

	node, err := client.CollectNode(
		"nwp_xp",
		"nwpc_pd",
		"login_b01",
		"31071",
		"/meso_post/00/initial")
	if err != nil {
		log.Fatalf("collect node error: %v", err)
	}

	fmt.Printf("Node Name: %s\n", node.Name)
}
