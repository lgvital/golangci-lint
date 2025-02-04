package golinters

import (
	"github.com/kunwardeep/paralleltest/pkg/paralleltest"
	"golang.org/x/tools/go/analysis"

	"github.com/lgvital/golangci-lint/pkg/config"
	"github.com/lgvital/golangci-lint/pkg/golinters/goanalysis"
)

func NewParallelTest(settings *config.ParallelTestSettings) *goanalysis.Linter {
	a := paralleltest.Analyzer

	var cfg map[string]map[string]interface{}
	if settings != nil {
		cfg = map[string]map[string]interface{}{
			a.Name: {
				"i": settings.IgnoreMissing,
			},
		}
	}

	return goanalysis.NewLinter(
		"paralleltest",
		"paralleltest detects missing usage of t.Parallel() method in your Go test",
		[]*analysis.Analyzer{paralleltest.Analyzer},
		cfg,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
