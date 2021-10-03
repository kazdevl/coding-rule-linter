package analyzer

import (
	"regexp"
)

const (
	VAR_NAME_PATTERN       string = `^[A-Za-z]+$`
	CONST_NAME_PATTERN     string = `^([A-Z_]+)$`
	TEST_FUNC_NAME_PATTERN string = `^Test_[A-Za-z]+$`
)

func applyVarNamingRule(name string) bool {
	if name == "_" {
		return true
	}
	return regexp.MustCompile(VAR_NAME_PATTERN).MatchString(name)
}

func applyConstNamingRule(name string) bool {
	return regexp.MustCompile(CONST_NAME_PATTERN).MatchString(name)
}

func applyTestFuncNamingRule(name string) bool {
	return regexp.MustCompile(TEST_FUNC_NAME_PATTERN).MatchString(name)
}
