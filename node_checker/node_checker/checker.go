package node_checker

import (
	"github.com/perillaroc/ecflow-client-go"
	"github.com/perillaroc/workflow-model-go"
	"log"
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
	for i := range checker.CheckTasks {
		if checker.CheckTasks[i].TriggerFlag == EvaluatedFit {
			isFit := checker.CheckTasks[i].Check(checker.node)
			log.Printf("[%s][%s] isFit = %s\n", checker.Name, checker.CheckTasks[i].Name, isFit)
		}
	}

}
