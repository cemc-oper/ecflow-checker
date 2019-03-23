package node_checker

import (
	"log"
)

func RunCheckTasks() {
	config := Config{}
	err := config.ReadConfig("./dist/conf/nwpc_op.yml")
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
	if !checker.EvaluateAll() {
		log.Printf("%s: all triggers is not fit", checker.Name)
		return
	}

	err := checker.FetchWorkflowNode()
	if err != nil {
		log.Printf("%s: fetch node failed: %v", checker.Name, err)
		return
	}

	checker.CheckAll()
}
