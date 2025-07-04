package internal

import "fmt"

type RuleCode string

const (
	GC001 = RuleCode("GC001")
	GC002 = RuleCode("GC002")
	GC013 = RuleCode("GC013")
	GC021 = RuleCode("GC021")
)

type Report struct {
	Rule    RuleCode
	Message string
}

func (r Report) String() string {
	return fmt.Sprintf("%s: %s", r.Rule, r.Message)
}
