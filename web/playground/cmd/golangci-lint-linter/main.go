//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/manuelarte/golangci-lint-linter/commands"
)

func apply(this js.Value, args []js.Value) any {
	input := args[0].String()

	outputFixedGolangci, _, err := commands.Lint([]byte(input), true)
	if err != nil {
		// TODO: handle error
		return ""
	}

	o, err := outputFixedGolangci.Marshal()
	if err != nil {
		return ""
	}

	msg := fmt.Sprintf("WebAssembly Output: %s\n", o)
	fmt.Print(msg)

	return string(o)
}

func main() {
	// channel is used so we can block on it, otherwise go process will exit when this module is loaded,
	// and we want the exported functions still available for later executions
	c := make(chan struct{}, 0)

	js.Global().Set("apply", js.FuncOf(apply))

	fmt.Println("Hello WebAssembly!")
	<-c
}
