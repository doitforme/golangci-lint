package internal

import "github.com/doitforme/golangci-lint/pkg/logutils"

// FormatterLogger must be used only when the context logger is not available.
var FormatterLogger = logutils.NewStderrLog(logutils.DebugKeyFormatter)
