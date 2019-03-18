package ecflow_client_go

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/perillaroc/ecflow-client-go/ecflowclient"
	"github.com/perillaroc/workflow-model-go"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
)

type EcflowClient struct {
	Target        string
	Connection    *grpc.ClientConn
	Context       context.Context
	ServiceClient pb.EcflowClientServiceClient
}

func (c *EcflowClient) Connect() error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip"))}

	var err error

	c.Connection, err = grpc.Dial(c.Target, opts...)

	if err != nil {
		return err
	}
	c.ServiceClient = pb.NewEcflowClientServiceClient(c.Connection)
	c.Context = context.Background()
	return nil
}

func (c *EcflowClient) Close() {
	c.Connection.Close()
}

func (c *EcflowClient) CollectStatusRecords(owner string, repo string, host string, port string) {
	r, err := c.ServiceClient.CollectStatusRecords(c.Context, &pb.StatusRequest{
		Owner: owner,
		Repo:  repo,
		Host:  host,
		Port:  port,
	})

	if err != nil {
		log.Fatalf("Could not get status: %v", err)
	}

	log.Printf("Size of records: %d\n", len(r.StatusMap))

	bunch := workflowmodel.Bunch{}
	for path, status := range r.StatusMap {
		bunch.AddNodeStatus(workflowmodel.NodeStatusRecord{Path: path, Status: status})
	}

	for _, suite := range bunch.Children {
		fmt.Printf("%s\n", suite.NodePath())
	}
}

func (c *EcflowClient) CollectNode(
	owner string, repo string, host string, port string, path string) (*workflowmodel.WorkflowNode, error) {
	r, err := c.ServiceClient.CollectNode(c.Context, &pb.NodeRequest{
		Owner: owner,
		Repo:  repo,
		Host:  host,
		Port:  port,
		Path:  path,
	})

	if err != nil {
		err = fmt.Errorf("Could not get status: %v\n", err)
		return nil, err
	}

	if r.ResponseStatus != nil && r.ResponseStatus.HasError {
		err = fmt.Errorf("CollectNode has error: %s\n", r.ResponseStatus.ErrorString)
		return nil, err
	}

	var node workflowmodel.WorkflowNode
	err = json.Unmarshal([]byte(r.Node), &node)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal failed: %v\n", err)
		return nil, err
	}
	return &node, err
}
