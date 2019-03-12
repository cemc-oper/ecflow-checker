package main

import (
	"context"
	"fmt"
	pb "github.com/perillaroc/eclfow-client-go/ecflowclient"
	"github.com/perillaroc/workflow-model-go"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip"))}

	conn, err := grpc.Dial("10.28.32.114:31181", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewEcflowClientServiceClient(conn)

	ctx := context.Background()

	r, err := c.CollectStatusRecords(ctx, &pb.StatusRequest{
		Owner: "nwp_xp",
		Repo:  "nwpc_pd_bk",
		Host:  "10.40.143.18",
		Port:  "31071",
	})

	if err != nil {
		log.Fatalf("Could not get status: %v", err)
	}

	log.Printf("Size of records: %d\n", len(r.StatusMap))

	bunch := workflowmodel.Bunch{}
	for path, status := range r.StatusMap {
		bunch.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: path, Status: status})
	}

	fmt.Printf("bunch children: %v", bunch.Children)
}
