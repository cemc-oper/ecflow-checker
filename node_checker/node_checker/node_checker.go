package node_checker

import (
	"github.com/perillaroc/ecflow-client-go"
	"github.com/perillaroc/workflow-model-go"
	"time"
)

type NodeCheckItem struct {
	workflowmodel.WorkflowNodeCondition
	FitFlag bool
}

type TimeTrigger struct {
	BeginTime time.Time
	EndTime   time.Time
}

func (t *TimeTrigger) Evaluate() bool {
	current := time.Now()
	return current.After(t.BeginTime) && current.Before(t.EndTime)
}

type NodeChecker struct {
	Target   string
	Owner    string
	Repo     string
	Host     string
	Port     string
	NodePath string

	TimeTrigger

	CheckItems []NodeCheckItem

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
