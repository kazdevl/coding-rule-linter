package analyzer_test

import (
	"testing"

	"github.com/KazuwoKiwame12/coding-rule-linter/pkg/analyzer"
)

func Test_applyVarNamingRule(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "failed: number exists",
			input: "sample12",
		},
		{
			name:  "failed: symnol exsits",
			input: "sample!",
		},
		{
			name:  "success",
			input: "sample",
			want:  true,
		},
		{
			name:  "success: blanck name",
			input: "_",
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := analyzer.ExportApplyVarNamingRule(test.input)
			if ok != test.want {
				t.Errorf("unmatched error: result is %v, want is %v", ok, test.want)
			}
		})
	}
}

func Test_applyConstNamingRule(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "failed: lower case letters exist",
			input: "SaMple",
		},
		{
			name:  "failed: number exists",
			input: "SAMPLE12",
		},
		{
			name:  "failed: symnol exsits",
			input: "SAMPLE!",
		},
		{
			name:  "success",
			input: "SAMPLE",
			want:  true,
		},
		{
			name:  "success: underscore exists",
			input: "SMAPLE_SAMPLE",
			want:  true,
		},
		{
			name:  "success: firt letter is underscore",
			input: "_SAMPLE",
			want:  true,
		},
		{
			name:  "success: blanck name",
			input: "_",
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := analyzer.ExportApplyConstNamingRule(test.input)
			if ok != test.want {
				t.Errorf("unmatched error: result is %v, want is %v", ok, test.want)
			}
		})
	}
}

func Test_applyTestFuncNamingRule(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "failed: \"Test_\" does not exist",
			input: "sample",
		},
		{
			name:  "failed: \"_\" does not exist between \"Test\" and \"FuncName\"",
			input: "TestSample",
		},
		{
			name:  "failed: first letter, \"t\" is lower case",
			input: "test_sample",
		},
		{
			name:  "failed: func name does not exist",
			input: "Test_",
		},
		{
			name:  "failed: symnol exsits",
			input: "Test_!",
		},
		{
			name:  "success",
			input: "Test_Sample",
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := analyzer.ExportApplyTestFuncNamingRule(test.input)
			if ok != test.want {
				t.Errorf("unmatched error: result is %v, want is %v", ok, test.want)
			}
		})
	}
}
