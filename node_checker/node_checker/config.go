package node_checker

import (
	"fmt"
	"github.com/perillaroc/workflow-model-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type ConfigDict map[interface{}]interface{}
type ConfigArray []interface{}

type Config struct {
	EcflowConfig *EcflowServerConfig
	Config       ConfigDict
	Checkers     []*NodeChecker
}

func (config *Config) ReadConfig() error {
	data, err := ioutil.ReadFile("./dist/conf/nwpc_op.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &config.Config)
	if err != nil {
		return err
	}
	return nil
}

func (config *Config) ReadEcflowServerConfig() {
	g := config.Config["global"]
	globalSection := g.(ConfigDict)
	ecflowServerSection := globalSection["ecflow_server"].(ConfigDict)

	config.EcflowConfig = &EcflowServerConfig{
		Target: ecflowServerSection["grpc_target"].(string),
		Owner:  ecflowServerSection["owner"].(string),
		Repo:   ecflowServerSection["repo"].(string),
		Host:   ecflowServerSection["host"].(string),
		Port:   fmt.Sprintf("%d", ecflowServerSection["port"].(int)),
	}
}

func (config *Config) ReadCheckTaskList() error {
	taskListSection := config.Config["task_list"].([]interface{})
	for _, item := range taskListSection {
		taskConfig := item.(ConfigDict)

		checker := &NodeChecker{}
		checker.Name = taskConfig["name"].(string)
		checker.NodePath = taskConfig["node_path"].(string)

		if config.EcflowConfig != nil {
			checker.EcflowServerConfig = *config.EcflowConfig
		}

		config.readCheckTasks(taskConfig["check_tasks"].([]interface{}), checker)
		config.Checkers = append(config.Checkers, checker)
	}

	return nil
}

func (config *Config) readCheckTasks(checkTasksConfig ConfigArray, checker *NodeChecker) {
	for _, checkTaskConfig := range checkTasksConfig {
		checkTaskConfigDict := checkTaskConfig.(ConfigDict)
		checkTask := &NodeCheckTask{}

		checkTask.Name = checkTaskConfigDict["name"].(string)

		triggersConfig := checkTaskConfigDict["trigger"].([]interface{})
		config.readTriggers(triggersConfig, checkTask)

		checkListConfig := checkTaskConfigDict["check_list"].([]interface{})
		config.readCheckList(checkListConfig, checkTask)
		checker.CheckTasks = append(checker.CheckTasks, *checkTask)
	}
}

func (config *Config) readTriggers(triggersConfig ConfigArray, checkTask *NodeCheckTask) {
	for _, triggerConfig := range triggersConfig {
		triggerConfigDict := triggerConfig.(ConfigDict)
		triggerType := triggerConfigDict["type"]
		if triggerType == "time" {
			beginTime, beginTimeErr := ParseClockUTC(triggerConfigDict["begin_time"].(string))
			if beginTimeErr != nil {
				panic(beginTimeErr)
			}

			endTime, endTimeErr := ParseClockUTC(triggerConfigDict["end_time"].(string))
			if endTimeErr != nil {
				panic(endTimeErr)
			}
			trigger := &TimeTrigger{
				BeginTime: beginTime,
				EndTime:   endTime,
			}
			checkTask.Triggers = append(checkTask.Triggers, trigger)
		}
	}
}

func (config *Config) readCheckList(checkListConfig ConfigArray, checkTask *NodeCheckTask) {
	for _, item := range checkListConfig {
		checkConfig := item.(ConfigDict)
		checkType := checkConfig["type"].(string)
		if checkType == "variable" {
			variableCheckItem, err := config.readVariableCheck(checkConfig)
			if err != nil {
				panic(err)
			}
			checkTask.CheckItems = append(checkTask.CheckItems, *variableCheckItem)
		} else if checkType == "status" {
			statusCheckItem, err := config.readStatusCheck(checkConfig)
			if err != nil {
				panic(err)
			}
			checkTask.CheckItems = append(checkTask.CheckItems, *statusCheckItem)
		}
	}
}

func (config *Config) readStatusCheck(checkConfig ConfigDict) (*NodeCheckItem, error) {
	var err error = nil
	var checkItem *NodeCheckItem = nil
	valueSection := checkConfig["value"].(ConfigDict)

	valueOperator := valueSection["operator"]
	valueFields := valueSection["fields"]

	if valueOperator == "equal" {
		value := valueFields.(string)
		checkItem = &NodeCheckItem{
			WorkflowNodeCondition: &workflowmodel.WorkflowNodeStatusCondition{
				Checker: &workflowmodel.NodeStatusEqualValueChecker{
					ExpectedValue: workflowmodel.GetNodeStatus(value),
				},
			},
			FitFlag: false,
		}

	} else if valueOperator == "in" {
		values := valueFields.([]interface{})

		var statusArray []workflowmodel.NodeStatus
		for _, statusValue := range values {
			statusArray = append(statusArray, workflowmodel.GetNodeStatus(statusValue.(string)))
		}
		checkItem = &NodeCheckItem{
			WorkflowNodeCondition: &workflowmodel.WorkflowNodeStatusCondition{
				Checker: &workflowmodel.NodeStatusInValueChecker{
					ExpectedValues: statusArray,
				},
			},
			FitFlag: false,
		}

	} else {
		err = fmt.Errorf("valueOperator not supported: %v", valueOperator)
	}
	return checkItem, err
}

func (config *Config) readVariableCheck(checkConfig ConfigDict) (*NodeCheckItem, error) {
	var err error = nil
	var checkItem *NodeCheckItem = nil
	name := checkConfig["name"].(string)
	valueSection := checkConfig["value"].(ConfigDict)

	valueType := valueSection["type"].(string)
	valueOperator := valueSection["operator"]
	valueFields := valueSection["fields"]

	if valueOperator == "equal" {
		value := valueFields.(string)
		if valueType == "date" {
			if value == "current" {
				value = time.Now().Format("20060102")
			}
		}

		checkItem = &NodeCheckItem{
			WorkflowNodeCondition: &workflowmodel.WorkflowNodeVariableCondition{
				VariableName: name,
				Checker: &workflowmodel.StringEqualValueChecker{
					ExpectedValue: value,
				},
			},
			FitFlag: false,
		}

	} else {
		err = fmt.Errorf("valueOperator not supported: %v\n", valueOperator)
	}
	return checkItem, err
}
