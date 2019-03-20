package node_checker

import "github.com/perillaroc/workflow-model-go"

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
