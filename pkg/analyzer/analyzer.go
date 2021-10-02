package analyzer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "code-rule-linter is analyzer that detects violations of coding rule"

var Analyzer = &analysis.Analyzer{
	Name:     "coding-rule-linter",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
