package node_checker

import "github.com/perillaroc/workflow-model-go"

type TriggerStatus uint

const (
	Unknown = iota
	Fit
	UnFit
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

type NodeCheckTask struct {
	Name       string
	Triggers   []Trigger
	CheckItems []NodeCheckItem

	TriggerFlag TriggerStatus
}

func (task *NodeCheckTask) Evaluate() bool {
	flag := false
	for _, trigger := range task.Triggers {
		if trigger.Evaluate() {
			task.TriggerFlag = Fit
			flag = true
		}
	}
	if !flag {
		task.TriggerFlag = UnFit
	}
	return flag
}

func (task *NodeCheckTask) IsFit(node *workflowmodel.WorkflowNode) bool {
	if node == nil {
		return false
	}
	isFit := true
	for _, checkItem := range task.CheckItems {
		checkItem.FitFlag = checkItem.IsFit(node)
		if !checkItem.FitFlag {
			isFit = false
		}
	}
	return isFit
}
