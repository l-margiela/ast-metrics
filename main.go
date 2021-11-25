package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

var idCounter int

// uniqueID is a monotonous integer generator.
//
// It would be great to use function name and/or a hint from a comment.
//
// E.g.
//
//   //metrics:RottenTomatoesAPICall
//   http.Get([…]
//
func uniqueID() int {
	idCounter++
	return idCounter
}

// metricsStarter is creating a new variable that holds timestamp at the beginning of a code block.
//
// It renders to the following code:
//
//   metricsStart<id> := time.Now()
//
func metricsStarter(id int) ast.Stmt {
	varName := fmt.Sprintf("metricsStart%d", id)
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: varName,
				Obj:  ast.NewObj(ast.Var, varName),
			},
		},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("time"), Sel: ast.NewIdent("Now")},
			},
		},
		Tok: token.DEFINE,
	}
}

// metricsStarter prints the time measurement
//
// It renders to the following code:
//
//   fmt.Printf("Block time measurement (ID <id>); time: %s", time.Since(metricsStart<id>))
//
func metricsEnder(id int) ast.Stmt {
	varName := fmt.Sprintf("metricsStart%d", id)
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("fmt"),
				Sel: ast.NewIdent("Printf"),
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: `"Block time measurement (ID ` + strconv.Itoa(id) + `); time: %s\n"`,
				},
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("time"),
						Sel: ast.NewIdent("Since"),
					},
					Args: []ast.Expr{
						&ast.Ident{Name: varName},
					},
				},
			},
		},
	}
}

// stmtsWithMetrics wraps ast.Stmt slice in metrics.
func stmtsWithMetrics(stmts []ast.Stmt, id int) []ast.Stmt {
	withStarter := append([]ast.Stmt{metricsStarter(id)}, stmts...)
	return append(withStarter, metricsEnder(id))
}

// blocksWithMetrics walks through the ast.File and wraps blocks and switch case clauses in metrics.
func blocksWithMetrics(f *ast.File) {
	astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		block, ok := n.(*ast.BlockStmt)
		if !ok {
			return true
		}

		id := uniqueID()
		withStarter := append([]ast.Stmt{metricsStarter(id)}, block.List...)
		blockWithMetrics := append(withStarter, metricsEnder(id))

		blockLen := len(blockWithMetrics)
		_, okRet := blockWithMetrics[blockLen-2].(*ast.ReturnStmt)
		if okRet {
			blockWithMetrics[blockLen-1], blockWithMetrics[blockLen-2] = blockWithMetrics[blockLen-2], blockWithMetrics[blockLen-1]
		}

		maybeSwitch := c.Parent()
		_, ok = maybeSwitch.(*ast.SwitchStmt)
		if ok {
			return true
		}

		c.Replace(&ast.BlockStmt{
			List: blockWithMetrics,
		})

		return true
	})
}

// switchCasesWithMetrics wraps switch case clauses in metrics.
func switchCasesWithMetrics(f *ast.File) {
	astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		clause, ok := n.(*ast.CaseClause)
		if !ok {
			return true
		}

		clause.Body = stmtsWithMetrics(clause.Body, uniqueID())

		maybeSwitch := c.Parent()
		_, ok = maybeSwitch.(*ast.SwitchStmt)
		if ok {
			return true
		}

		c.Replace(clause)

		return true
	})
}

func main() {
	fset := token.NewFileSet()
	fpath := "somepkg/main.go"
	file, err := parser.ParseFile(fset, fpath, nil, 0)
	if err != nil {
		panic(err)
	}

	blocksWithMetrics(file)
	switchCasesWithMetrics(file)

	modFile := path.Join(path.Dir(fpath)+"mod", path.Base(fpath))
	fd, err := os.Create(modFile)
	if err != nil {
		panic(err)
	}
	printer.Fprint(fd, fset, file)
}
