package goimports

import (
	"golang.org/x/tools/go/analysis"

	"github.com/doitforme/golangci-lint/pkg/config"
	"github.com/doitforme/golangci-lint/pkg/goanalysis"
	"github.com/doitforme/golangci-lint/pkg/goformatters"
	goimportsbase "github.com/doitforme/golangci-lint/pkg/goformatters/goimports"
	"github.com/doitforme/golangci-lint/pkg/golinters/internal"
)

const linterName = "goimports"

func New(settings *config.GoImportsSettings) *goanalysis.Linter {
	a := goformatters.NewAnalyzer(
		internal.LinterLogger.Child(linterName),
		"Checks if the code and import statements are formatted according to the 'goimports' command.",
		goimportsbase.New(settings),
	)

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
