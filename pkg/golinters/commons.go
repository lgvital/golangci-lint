package golinters

import "github.com/lgvital/golangci-lint/pkg/logutils"

// linterLogger must be use only when the context logger is not available.
var linterLogger = logutils.NewStderrLog("linter")
