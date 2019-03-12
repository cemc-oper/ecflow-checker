package main

import (
	"context"
	"fmt"
	pb "github.com/perillaroc/eclfow-client-go/ecflowclient"
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

	r, err := c.CollectNode(ctx, &pb.NodeRequest{
		Owner: "nwp_xp",
		Repo:  "nwpc_pd_bk",
		Host:  "10.40.143.18",
		Port:  "31071",
		Path:  "/meso_post/00/initial",
	})

	if err != nil {
		log.Fatalf("Could not get status: %v", err)
	}

	fmt.Printf("node: %v", r.Node)
}
