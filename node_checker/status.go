package node_checker

type TriggerStatus uint

const (
	UnEvaluated TriggerStatus = iota
	EvaluatedFit
	EvaluatedUnFit
)

func (status TriggerStatus) String() string {
	switch status {
	case UnEvaluated:
		return "UnEvaluated"
	case EvaluatedFit:
		return "EvaluatedFit"
	case EvaluatedUnFit:
		return "EvaluatedUnFit"
	default:
		return "Unknown"
	}
}

type ConditionStatus uint

const (
	UnChecked ConditionStatus = iota
	ConditionFit
	ConditionUnFit
)

func (status ConditionStatus) String() string {
	switch status {
	case UnChecked:
		return "UnChecked"
	case ConditionFit:
		return "ConditionFit"
	case ConditionUnFit:
		return "ConditionUnFit"
	default:
		return "Unknown"
	}
}
