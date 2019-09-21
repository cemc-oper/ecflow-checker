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
			fmt.Printf("%s\n", red(fmt.Sprintf("%s: Fetching node...failed", bold(checker.Name))))
			return
		}

		checker.CheckAll()
	} else {
		fmt.Printf("%s: Ignore\n", bold(checker.Name))
	}
}
