package service

import (
	"fmt"
	"go/ast"
	"go/parser"
	"reflect"
	"strconv"
	"strings"
)

func eval(expr string) (float64, error) {
	//fs := token.NewFileSet()
	tr, err := parser.ParseExpr(expr)
	if err != nil {
		return 0, err
	}

	//ast.Print(fs, tr)
	result, err := traverseTree(tr)
	fmt.Printf("eval: %s = %f\n", expr, result)
	return result, err
}

func traverseTree(node ast.Node) (float64, error) {
	//fmt.Printf("node is: %s\n", reflect.TypeOf(node))
	switch n := node.(type) {
	case *ast.ParenExpr:
		return traverseTree(n.X)
	case *ast.BinaryExpr:
		X, err := traverseTree(n.X)
		if err != nil {
			return 0, err
		}
		Y, err := traverseTree(n.Y)
		if err != nil {
			return 0, err
		}
		switch n.Op.String() {
		case "+":
			return (X + Y), nil // TODO: call uService Add(traverseTree(n.X), traverseTree(n.Y))
		case "-":
			return (X + Y), nil
		case "*":
			return (X + Y), nil
		case "/":
			return (X + Y), nil
		default:
			return 0, fmt.Errorf("operator not supported: %s", n.Op)
		}
	case *ast.UnaryExpr:
		return 0, fmt.Errorf("unary not supported: %v", n)
	case *ast.BasicLit:
		//fmt.Printf("lit: %s\n", n.Value)
		if v, err := strconv.ParseFloat(strings.TrimSpace(n.Value), 64); err != nil {
			return 0, fmt.Errorf("unable to convert number to float64: %s", err)
		} else {
			return v, nil
		}

	default:
		return 0, fmt.Errorf("not supported node: %s", reflect.TypeOf(node))
	}

	return 0, fmt.Errorf("mmm...should not get here :(")
}
