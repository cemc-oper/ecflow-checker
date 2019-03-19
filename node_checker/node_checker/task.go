package node_checker

import "github.com/perillaroc/workflow-model-go"

type TriggerStatus uint

const (
	UnEvaluated TriggerStatus = iota
	EvaluatedFit
	EvaluatedUnFit
)

func (status TriggerStatus) String() string {
	switch status {
	case UnEvaluated:
		return "UnEvaluated"
	case EvaluatedFit:
		return "EvaluatedFit"
	case EvaluatedUnFit:
		return "EvaluatedUnFit"
	default:
		return "Unknown"
	}
}

type ConditionStatus uint

const (
	UnChecked ConditionStatus = iota
	ConditionFit
	ConditionUnFit
)

func (status ConditionStatus) String() string {
	switch status {
	case UnChecked:
		return "UnChecked"
	case ConditionFit:
		return "ConditionFit"
	case ConditionUnFit:
		return "ConditionUnFit"
	default:
		return "Unknown"
	}
}

type NodeCheckItem struct {
	workflowmodel.WorkflowNodeCondition
	ConditionFlag ConditionStatus
}

type NodeCheckTask struct {
	Name       string
	Triggers   []Trigger
	CheckItems []NodeCheckItem

	TriggerFlag TriggerStatus
}

func (task *NodeCheckTask) Evaluate() TriggerStatus {
	flag := EvaluatedUnFit
	for i := range task.Triggers {
		if task.Triggers[i].Evaluate() {
			task.TriggerFlag = EvaluatedFit
			flag = EvaluatedFit
		}
	}
	task.TriggerFlag = flag
	return flag
}

func (task *NodeCheckTask) IsFit(node *workflowmodel.WorkflowNode) ConditionStatus {
	if node == nil {
		return UnChecked
	}
	conditionFit := ConditionFit
	for i := range task.CheckItems {
		flag := task.CheckItems[i].IsFit(node)
		if !flag {
			conditionFit = ConditionUnFit
		}
		task.CheckItems[i].ConditionFlag = ConditionUnFit
	}
	return conditionFit
}
