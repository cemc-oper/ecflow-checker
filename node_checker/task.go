package node_checker

import "github.com/perillaroc/workflow-model-go"

type NodeCheckTask struct {
	Name           string
	Triggers       []Trigger
	NodeCheckItems []NodeCheckItem

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

func (task *NodeCheckTask) Check(node *workflowmodel.WorkflowNode) ConditionStatus {
	if node == nil {
		return UnChecked
	}
	conditionFit := ConditionFit
	for i := range task.NodeCheckItems {
		task.NodeCheckItems[i].CheckCondition(node)
		if task.NodeCheckItems[i].ConditionFlag == ConditionUnFit {
			conditionFit = ConditionUnFit
		}
	}
	return conditionFit
}

func (task *NodeCheckTask) IsAllFit() ConditionStatus {
	conditionFlag := ConditionFit
	for i := range task.NodeCheckItems {
		if task.NodeCheckItems[i].ConditionFlag == ConditionUnFit {
			conditionFlag = ConditionUnFit
		}
	}
	return conditionFlag
}
