package internal

type RuleCode string

const (
	GC001 = RuleCode("GC001")
)

type Report struct {
	Rule    RuleCode
	Message string
}
