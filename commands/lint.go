package commands

import (
	"fmt"
	"github.com/manuelarte/golangci-lint-linter/models"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/manuelarte/golangci-lint-linter/internal"
	"github.com/manuelarte/golangci-lint-linter/internal/linters"
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

func Lint(input []byte, isFix bool) (models.Golangci, []internal.Report, error) {
	golangci, err := models.Parse(input)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading yaml file: %w", err)
	}

	allReports := make([]internal.Report, 0)

	for _, linter := range getAllLinters() {
		linterReports := linter.Lint(golangci)

		if isFix && len(linterReports) > 0 {
			if fixer, ok := linter.(linters.Fixer); ok {
				errFix := fixer.Fix(golangci)
				if errFix != nil {
					return nil, nil, fmt.Errorf("error applying fix: %w", errFix)
				}
			}
		} else {
			allReports = append(allReports, linterReports...)
		}
	}

	return golangci, allReports, nil
}

//nolint:gocognit // TODO: refactor later
func run(cmd *cobra.Command, args []string) {
	path := args[0]

	isFix, err := cmd.Flags().GetBool("fix")
	if err != nil {
		errorMsg := errorColor.Sprintf("Error reading flag fix: %s\n", err)
		cmd.PrintErrln(errorMsg)

		return
	}

	input, errReadingFile := os.ReadFile(filepath.Clean(path))
	if errReadingFile != nil {
		errorMsg := errorColor.Sprintf("Error reading file: %s\n", errReadingFile)
		cmd.PrintErrf(errorMsg)

		return
	}

	cmd.Printf("Linting: %q\n", path)
	golangciFixed, reports, err := Lint(input, isFix)
	if err != nil {
		errorMsg := errorColor.Sprintf("Error Linting %s: %s\n", path, err)
		cmd.PrintErrf(errorMsg)
		os.Exit(1)
	}

	cmd.Printf("Found: %d issue(s)\n", len(reports))

	if isFix {
		fixedFile, errM := golangciFixed.Marshal()
		if errM != nil {
			errorMsg := errorColor.Sprintf("Error Marshaling file: %s\n", errM)
			cmd.PrintErrf(errorMsg)
			os.Exit(1)
		}

		err = os.WriteFile(path, fixedFile, 0o600)
		if err != nil {
			errorMsg := errorColor.Sprintf("Error writing file: %s\n", err)
			cmd.PrintErrf(errorMsg)
			os.Exit(1)
		}
	}

	for _, report := range reports {
		errorMsg := errorColor.Sprintf("%s", report)
		cmd.PrintErrf("%s\n", errorMsg)
	}

	if len(reports) > 0 {
		os.Exit(1)
	}
}

func getAllLinters() []linters.Linter {
	return []linters.Linter{
		linters.NewLinterFieldsSorted(),
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
