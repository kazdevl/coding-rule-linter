package analyzer

import (
	"regexp"
)

const (
	varNamePattern      string = `^[a-z]{1}[A-Za-z]+$`
	constNamePattern    string = `^([A-Z_]+)$`
	testFuncNamePattern string = `^Test_[A-Za-z]+$`
)

func applyVarNamingRule(name string) bool {
	if name == "_" {
		return true
	}
	return regexp.MustCompile(varNamePattern).MatchString(name)
}

func applyConstNamingRule(name string) bool {
	return regexp.MustCompile(constNamePattern).MatchString(name)
}

func applyTestFuncNamingRule(name string) bool {
	return regexp.MustCompile(testFuncNamePattern).MatchString(name)
}
