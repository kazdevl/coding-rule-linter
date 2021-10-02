package analyzer

var (
	ExportApplyVarNamingRule      func(string) bool = applyVarNamingRule
	ExportApplyConstNamingRule    func(string) bool = applyConstNamingRule
	ExportApplyTestFuncNamingRule func(string) bool = applyTestFuncNamingRule
)
