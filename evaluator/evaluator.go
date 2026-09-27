package evaluator

import (
	"bihari/ast"
	"bihari/object"
)

func Eval(node ast.Node, env *Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node, env)

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.Set(node.Name.Value, val)
		return val

	case *ast.PrintStatement:
		val := Eval(node.Expression, env)
		return val

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}

	case *ast.Identifier:
		return evalIdentifier(node, env)

	}

	return nil
}

func evalProgram(program *ast.Program, env *Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return nil
}
