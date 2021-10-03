package analyzer

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "code-rule-linter is analyzer that detects violations of coding rule"

var Analyzer = &analysis.Analyzer{
	Name:     "codingRuleLinter",
	Doc:      doc,
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
				// TODO
				fmt.Printf("funcdecl: %+v\n", n)
			}
			return
		}
	})
	return nil, nil
}

func validateGenDecl(n *ast.GenDecl) (string, bool) {
	const reportDoc string = "identifier should follow the coding rules"
	for _, spec := range n.Specs {
		switch s := spec.(type) {
		case *ast.ValueSpec:
			for _, n := range s.Names {
				if n.Obj.Kind == ast.Var {
					if !applyVarNamingRule(n.Obj.Name) {
						return fmt.Sprintf("%s %s", n.Obj.Kind, reportDoc), false
					}
				} else if n.Obj.Kind == ast.Con {
					if !applyConstNamingRule(n.Obj.Name) {
						return fmt.Sprintf("%s %s", n.Obj.Kind, reportDoc), false
					}
				}
			}
		}
	}
	return "", true
}
