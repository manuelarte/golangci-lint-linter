package cmd

import (
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/spf13/cobra"

	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

func RegisterLintCommand() *cobra.Command {
	rootCMD := &cobra.Command{
		Use:     "golangci-lint-linter",
		Short:   "Linter for the configuration files for the well known linter golangci-lint.",
		Long:    "Linter for the configuration files for the well known linter golangci-lint.",
		Args:    cobra.ExactArgs(1),
		Version: version(),
		Run:     run,
	}

	rootCMD.PersistentFlags().BoolP("fix", "f", false, "update file to resolve fixable issues.")

	return rootCMD
}

func run(cmd *cobra.Command, args []string) {
	path := args[0]

	golangci, err := readYamlFile(path)
	if err != nil {
		cmd.PrintErrf("error reading yaml file: %s\n", err)

		return
	}

	cmd.Println(golangci)
}

func readYamlFile(path string) (internal.Golangci, error) {
	input, errReadingFile := os.ReadFile(filepath.Clean(path))
	if errReadingFile != nil {
		return internal.YamlGolangci{}, errReadingFile
	}

	return internal.Parse(input)
}

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "local"
	}

	return info.Main.Version
}
