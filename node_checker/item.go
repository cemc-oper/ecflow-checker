package node_checker

import "github.com/perillaroc/workflow-model-go"

type NodeCheckItem struct {
	workflowmodel.WorkflowNodeCondition
	ConditionFlag ConditionStatus
}

func (item *NodeCheckItem) CheckCondition(node *workflowmodel.WorkflowNode) {
	flag := item.IsFit(node)
	if flag {
		item.ConditionFlag = ConditionFit
	} else {
		item.ConditionFlag = ConditionUnFit
	}
}
