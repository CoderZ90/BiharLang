package evaluator

import (
	"bihari/ast"
	"bihari/object"
)

// Eval takes an AST node and evaluates it into a runtime object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}

	}

	return nil
}
