package main

import (
	"coding-rule-linter/pkg/analyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(analyzer.Analyzer)
}
