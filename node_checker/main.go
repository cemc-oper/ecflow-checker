package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/perillaroc/ecflow-client-go/node_checker/node_checker"
	"github.com/perillaroc/workflow-model-go"
	"log"
	"time"
)

func main() {
	var err error
	var beginTime, endTime time.Time

	beginTime, err = now.Parse("03:42")
	if err != nil {
		log.Fatalf("beginTime error: %v", err)
	}

	endTime, err = now.Parse("23:59")
	if err != nil {
		log.Fatalf("endTime error: %v", err)
	}

	checker := node_checker.NodeChecker{
		Target:   "10.40.140.18:31181",
		Owner:    "nwp_xp",
		Repo:     "nwpc_pd",
		Host:     "login_b01",
		Port:     "31071",
		NodePath: "/gmf_grapes_gfs_v2.2_post/00",

		TimeTrigger: node_checker.TimeTrigger{
			BeginTime: beginTime,
			EndTime:   endTime,
		},

		CheckItems: []node_checker.NodeCheckItem{
			{
				WorkflowNodeCondition: &workflowmodel.WorkflowNodeStatusCondition{
					Checker: &workflowmodel.NodeStatusInValueChecker{
						ExpectedValues: []workflowmodel.NodeStatus{
							workflowmodel.Submitted,
							workflowmodel.Active,
							workflowmodel.Complete,
						}},
				},
			},
		},
	}

	if !checker.Evaluate() {
		log.Fatalf("time trigger is not fit")
		return
	}

	err = checker.FetchWorkflowNode()
	if err != nil {
		log.Fatalf("fetch node failed: %v", err)
		return
	}

	isFit := checker.IsFit()

	fmt.Printf("isFit = %t", isFit)
}
