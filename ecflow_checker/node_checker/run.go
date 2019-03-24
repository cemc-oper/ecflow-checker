package node_checker

import (
	"fmt"
)

func RunCheckTasks(configFilePath string) {
	config := Config{}
	err := config.ReadConfig(configFilePath)
	if err != nil {
		panic(err)
	}
	config.ReadEcflowServerConfig()

	err = config.ReadCheckTaskList()
	if err != nil {
		panic(err)
	}

	checkTasks := config.Checkers

	for _, checker := range checkTasks {
		RunNodeChecker(checker)
	}
}

func RunNodeChecker(checker *NodeChecker) {
	result := checker.EvaluateAll()

	if result {
		err := checker.FetchWorkflowNode()
		if err != nil {
			fmt.Printf("%s: Fetching node...failed", checker.Name)
			return
		}
	}

	checker.CheckAll()
}
