package node_checker

import (
	"fmt"
	"github.com/perillaroc/ecflow-client-go"
	"github.com/perillaroc/workflow-model-go"
)

type NodeChecker struct {
	Name string
	EcflowServerConfig

	NodePath string

	CheckTasks []NodeCheckTask

	node *workflowmodel.WorkflowNode
}

func (checker *NodeChecker) FetchWorkflowNode() error {
	client := ecflow_client_go.EcflowClient{Target: checker.Target}
	err := client.Connect()
	if err != nil {
		return err
	}

	checker.node, err = client.CollectNode(
		checker.Owner,
		checker.Repo,
		checker.Host,
		checker.Port,
		checker.NodePath)

	return err
}

func (checker *NodeChecker) EvaluateAll() bool {
	hasFitTrigger := false
	for i := range checker.CheckTasks {
		triggerFlag := checker.CheckTasks[i].Evaluate()
		if triggerFlag == EvaluatedFit {
			hasFitTrigger = true
		}
	}
	return hasFitTrigger
}

func (checker *NodeChecker) CheckAll() {
	fmt.Printf("%s:\n", checker.Name)
	for i := range checker.CheckTasks {
		if checker.CheckTasks[i].TriggerFlag == EvaluatedFit {
			checker.CheckTasks[i].Check(checker.node)
			fmt.Printf("  [%s] Checking...Pass\n", checker.CheckTasks[i].Name)
		} else {
			fmt.Printf("  [%s] Ignore\n", checker.CheckTasks[i].Name)
		}
	}
	fmt.Println()
}
