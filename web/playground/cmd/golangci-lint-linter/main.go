//go:build js && wasm

package main

import (
	"fmt"

	"syscall/js"

	"github.com/manuelarte/golangci-lint-linter/commands"
	"github.com/manuelarte/golangci-lint-linter/models"
)

func apply(input string) (string, []models.Report, error) {
	outputFixedGolangci, reports, err := commands.Lint([]byte(input), true)
	if err != nil {
		return "", nil, err
	}

	o, err := outputFixedGolangci.Marshal()
	if err != nil {
		return "", nil, err
	}

	return string(o), reports, nil
}

func applyJS(_ js.Value, args []js.Value) any {
	input := args[0].String()
	output, reports, err := apply(input)

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	reportsStr := make([]string, len(reports))
	for i, report := range reports {
		reportsStr[i] = report.String()
	}

	toReturnReport := ""
	if len(reportsStr) > 0 {
		toReturnReport = reportsStr[0]
	}

	return map[string]any{
		"output":  output,
		"reports": toReturnReport,
		"error":   errStr,
	}
}

func main() {
	// channel is used so we can block on it, otherwise go process will exit when this module is loaded,
	// and we want the exported functions still available for later executions
	c := make(chan struct{})

	js.Global().Set("apply", js.FuncOf(applyJS))

	fmt.Println("WebAssembly linter started")
	<-c
}
