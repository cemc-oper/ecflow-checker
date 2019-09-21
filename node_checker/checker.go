package node_checker

import (
	"fmt"
	"github.com/perillaroc/ecflow-checker"
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
	client := ecflow_checker.EcflowClient{Target: checker.Target}
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
	fmt.Printf("%s:\n", bold(checker.Name))
	for i := range checker.CheckTasks {
		if checker.CheckTasks[i].TriggerFlag == EvaluatedFit {
			condition := checker.CheckTasks[i].Check(checker.node)
			var mark = blue("☐")
			if condition == ConditionFit {
				mark = green("✔")
			} else if condition == ConditionUnFit {
				mark = red("✗")
			}
			fmt.Printf("[%s] Checking for %s\n", mark, checker.CheckTasks[i].Name)
		} else {
			fmt.Printf("[%s] Checking for %s\n", blue("━"), checker.CheckTasks[i].Name)
		}
	}
}
