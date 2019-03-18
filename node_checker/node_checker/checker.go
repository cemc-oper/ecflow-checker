package node_checker

import (
	"github.com/perillaroc/ecflow-client-go"
	"github.com/perillaroc/workflow-model-go"
)

type NodeCheckItem struct {
	workflowmodel.WorkflowNodeCondition
	FitFlag bool
}

type EcflowServerConfig struct {
	Target string
	Owner  string
	Repo   string
	Host   string
	Port   string
}

type NodeChecker struct {
	Name string
	EcflowServerConfig

	NodePath string

	Triggers []Trigger

	CheckItems []NodeCheckItem

	node *workflowmodel.WorkflowNode
}

func (checker *NodeChecker) Evaluate() bool {
	flag := true
	for _, trigger := range checker.Triggers {
		if !trigger.Evaluate() {
			flag = false
		}
	}
	return flag
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

func (checker *NodeChecker) IsFit() bool {
	if checker.node == nil {
		return false
	}
	isFit := true
	for _, checkItem := range checker.CheckItems {
		checkItem.FitFlag = checkItem.IsFit(checker.node)
		if !checkItem.FitFlag {
			isFit = false
		}
	}
	return isFit
}
