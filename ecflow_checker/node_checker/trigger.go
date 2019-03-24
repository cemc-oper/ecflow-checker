package node_checker

import "time"

type Trigger interface {
	Evaluate() bool
}

type TimeTrigger struct {
	BeginTime time.Time
	EndTime   time.Time
}

func (t *TimeTrigger) Evaluate() bool {
	current := time.Now().UTC()
	return current.After(t.BeginTime) && current.Before(t.EndTime)
}
