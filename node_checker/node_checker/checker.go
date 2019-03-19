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
	for _, checkTask := range checker.CheckTasks {
		flag := checkTask.Evaluate()
		if flag {
			hasFitTrigger = true
		}
	}
	return hasFitTrigger
}

func (checker *NodeChecker) CheckFitAll() {
	for _, checkTask := range checker.CheckTasks {
		if checkTask.TriggerFlag == Fit {

		}
		isFit := checkTask.IsFit(checker.node)
		log.Printf("[%s][%s] isFit = %t\n", checker.Name, checkTask.Name, isFit)
	}

}
