package gofumpt

import (
	"golang.org/x/tools/go/analysis"

	"github.com/doitforme/golangci-lint/pkg/config"
	"github.com/doitforme/golangci-lint/pkg/goanalysis"
	"github.com/doitforme/golangci-lint/pkg/goformatters"
	gofumptbase "github.com/doitforme/golangci-lint/pkg/goformatters/gofumpt"
	"github.com/doitforme/golangci-lint/pkg/golinters/internal"
)

const linterName = "gofumpt"

func New(settings *config.GofumptSettings) *goanalysis.Linter {
	a := goformatters.NewAnalyzer(
		internal.LinterLogger.Child(linterName),
		"Checks if code and import statements are formatted, with additional rules.",
		gofumptbase.New(settings, settings.LangVersion),
	)

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
