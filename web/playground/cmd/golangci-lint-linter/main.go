//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func apply(this js.Value, args []js.Value) any {
	input := args[0].String()
	msg := fmt.Sprintf("WebAssembly Input: %s\n", input)
	fmt.Print(msg)
	return msg
}

func main() {
	// channel is used so we can block on it, otherwise go process will exit when this module is loaded,
	// and we want the exported functions still available for later executions
	c := make(chan struct{}, 0)

	js.Global().Set("apply", js.FuncOf(apply))

	fmt.Println("Hello WebAssembly!")
	<-c
}
