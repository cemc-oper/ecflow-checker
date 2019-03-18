package node_checker

import (
	"log"
)

func RunNodeChecker(checker *NodeChecker) {
	if !checker.Evaluate() {
		log.Printf("%s: time trigger is not fit", checker.Name)
		return
	}

	err := checker.FetchWorkflowNode()
	if err != nil {
		log.Printf("%s: fetch node failed: %v", checker.Name, err)
		return
	}

	isFit := checker.IsFit()

	log.Printf("%s: isFit = %t\n", checker.Name, isFit)
}

func RunCheckTasks() {
	config := Config{}
	err := config.ReadConfig()
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
