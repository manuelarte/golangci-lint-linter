package internal

type RuleCode string

const (
	GC001 = RuleCode("GC001")
	GC002 = RuleCode("GC002")
)

type Report struct {
	Rule    RuleCode
	Message string
}
