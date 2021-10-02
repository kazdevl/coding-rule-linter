package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"coding-rule-linter/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to ged wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "p")
}
