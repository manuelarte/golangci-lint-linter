package cmd

import (
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/manuelarte/golangci-lint-linter-go/internal"
	"github.com/manuelarte/golangci-lint-linter-go/internal/linters"
)

//nolint:gochecknoglobals // color red
var errorColor = color.New(color.FgRed)

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

	cmd.Printf("Linting: %q\n", path)

	allReports := make([]internal.Report, 0)

	for _, linter := range getAllLinters() {
		linterReports := linter.Lint(golangci)

		if len(linterReports) > 0 {
			if fixer, ok := linter.(linters.Fixer); ok {
				errFix := fixer.Fix(golangci)
				if errFix != nil {
					errorMsg := errorColor.Sprintf("Error applying fix: %s", errFix)
					cmd.PrintErrf("%s\n", errorMsg)
					os.Exit(1)
				}
			}
		}

		allReports = append(allReports, linterReports...)
	}

	cmd.Printf("Found: %d issue(s)\n", len(allReports))

	for _, report := range allReports {
		errorMsg := errorColor.Sprintf("%s", report)
		cmd.PrintErrf("%s\n", errorMsg)
	}

	if len(allReports) > 0 {
		os.Exit(1)
	}
}

func readYamlFile(path string) (internal.Golangci, error) {
	input, errReadingFile := os.ReadFile(filepath.Clean(path))
	if errReadingFile != nil {
		return internal.YamlGolangci{}, errReadingFile
	}

	return internal.Parse(input)
}

func getAllLinters() []linters.Linter {
	return []linters.Linter{
		linters.NewLintersAlphabetical(),
		linters.NewSettingsAlphabetical(),
		linters.NewDisabledLintersWithReason(),
	}
}

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "local"
	}

	return info.Main.Version
}
