package analyzer

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const DOC = "code-rule-linter is analyzer that detects violations of coding rule"

var Analyzer = &analysis.Analyzer{
	Name:     "codingRuleLinter",
	Doc:      DOC,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.GenDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.GenDecl:
			problem, ok := validateGenDecl(n)
			if ok {
				return
			}
			pass.Reportf(n.Pos(), problem)
		case *ast.FuncDecl:
			if strings.HasSuffix(pass.Fset.File(n.Pos()).Name(), "_test.go") {
				problem, ok := validateFuncDecl(n)
				if ok {
					return
				}
				pass.Reportf(n.Pos(), problem)
			}
			return
		}
	})
	return nil, nil
}

func validateGenDecl(n *ast.GenDecl) (string, bool) {
	const REPORT_TEXT string = "should follow the coding rules"
	for _, spec := range n.Specs {
		switch s := spec.(type) {
		case *ast.ValueSpec:
			for _, n := range s.Names {
				if n.Obj.Kind == ast.Var {
					if !applyVarNamingRule(n.Obj.Name) {
						return fmt.Sprintf("%s(%s identifier) %s", n.Obj.Name, n.Obj.Kind, REPORT_TEXT), false
					}
				} else if n.Obj.Kind == ast.Con {
					if !applyConstNamingRule(n.Obj.Name) {
						return fmt.Sprintf("%s(%s identifier) %s", n.Obj.Name, n.Obj.Kind, REPORT_TEXT), false
					}
				}
			}
		}
	}
	return "", true
}

func validateFuncDecl(n *ast.FuncDecl) (string, bool) {
	if n.Type.Results != nil { // the func that return value is not target
		return "", true
	}

	if len(n.Type.Params.List) != 1 { // the func that has some arguments(different type) is not target
		return "", true
	}

	if len(n.Type.Params.List[0].Names) != 1 { // the func that has some arguments(same type) is not target
		return "", true
	}

	starExpr, ok := n.Type.Params.List[0].Type.(*ast.StarExpr)
	if !ok {
		return "", true
	}

	argSelector, ok := starExpr.X.(*ast.SelectorExpr)
	if !ok {
		return "", true
	}

	argType := fmt.Sprintf("%s.%s", argSelector.X.(*ast.Ident).Name, argSelector.Sel.Name)
	if argType != "testing.T" { // the func that does't have the testing.T arg-type is not target
		return "", true
	}

	if !applyTestFuncNamingRule(n.Name.Name) {
		return fmt.Sprintf("%s(test func name) should follow the coding rules", n.Name.Name), false // TODO 現状と理想を提示する
	}
	return "", true
}
