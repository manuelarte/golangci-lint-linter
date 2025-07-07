package main

import (
	"fmt"
	"github.com/manuelarte/golangci-lint-linter/commands"
	"os"
)

func main() {
	rootCMD := commands.RegisterLintCommand()

	err := rootCMD.Execute()
	if err != nil {
		//nolint:forbidigo // main
		fmt.Println(err)
		os.Exit(1)
	}
}
