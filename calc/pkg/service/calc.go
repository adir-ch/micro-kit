package service

import (
	"fmt"
	"go/ast"
	"go/parser"
	"reflect"
	"strconv"
	"strings"
)

func eval(expr string) float64 {
	//fs := token.NewFileSet()
	tr, _ := parser.ParseExpr(expr)
	//ast.Print(fs, tr)
	result := traverseTree(tr)
	fmt.Printf("eval: %s = %f\n", expr, result)
	return result
}

func traverseTree(node ast.Node) float64 {
	//fmt.Printf("node is: %s\n", reflect.TypeOf(node))
	switch n := node.(type) {
	case *ast.ParenExpr:
		return traverseTree(n.X)
	case *ast.BinaryExpr:
		switch n.Op.String() {
		case "+":
			return traverseTree(n.X) + traverseTree(n.Y) // TODO: call uService Add(traverseTree(n.X), traverseTree(n.Y))
		case "-":
			return traverseTree(n.X) - traverseTree(n.Y)
		case "*":
			return traverseTree(n.X) * traverseTree(n.Y)
		case "/":
			return traverseTree(n.X) / traverseTree(n.Y)
		default:
			fmt.Printf("not supported operator: %s\n", n.Op)
		}
	case *ast.BasicLit:
		//fmt.Printf("lit: %s\n", n.Value)
		if v, err := strconv.ParseFloat(strings.TrimSpace(n.Value), 64); err != nil {
			fmt.Printf("unable to convert number to float64: %s\n", err)
			return -99 // for now :(
		} else {
			return v
		}
	default:
		fmt.Printf("not supported node: %s\n", reflect.TypeOf(node))
	}

	fmt.Printf("mmm...should not get here :(\n")
	return 0
}
