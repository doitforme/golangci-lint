package gci

import (
	"golang.org/x/tools/go/analysis"

	"github.com/doitforme/golangci-lint/pkg/config"
	"github.com/doitforme/golangci-lint/pkg/goanalysis"
	"github.com/doitforme/golangci-lint/pkg/goformatters"
	gcibase "github.com/doitforme/golangci-lint/pkg/goformatters/gci"
	"github.com/doitforme/golangci-lint/pkg/golinters/internal"
)

const linterName = "gci"

func New(settings *config.GciSettings) *goanalysis.Linter {
	formatter, err := gcibase.New(settings)
	if err != nil {
		internal.LinterLogger.Fatalf("%s: create analyzer: %v", linterName, err)
	}

	a := goformatters.NewAnalyzer(
		internal.LinterLogger.Child(linterName),
		"Checks if code and import statements are formatted, with additional rules.",
		formatter,
	)

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
