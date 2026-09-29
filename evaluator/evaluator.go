package evaluator

import (
	"bihari/ast"
	"bihari/object"
	"fmt"
)

func Eval(node ast.Node, env *Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node, env)

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.Set(node.Name.Value, val)
		return nil

	case *ast.PrintStatement:
		val := Eval(node.Expression, env)
		if isError(val) {
			return val // propagate error
		}
		fmt.Println(val.Inspect())
		return nil

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)
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

	return newError("are bhai, '%s' naam ke variable hawa me ta na ba! Pehle `rakha` ta.", node.Value)
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	if isError(left) {
		return left
	}
	if isError(right) {
		return right
	}
	if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
		return evalIntegerInfixExpression(operator, left, right)
	}
	if left.Type() != right.Type() {
		return newError("bhai, %s aur %s ke jod na ho sakat baa! 💀", left.Type(), right.Type())
	}

	return newError("bhai, %s %s %s ka matlab ka baa? 💀", left.Type(), operator, right.Type())
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return nil
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return &object.Boolean{Value: true}
	}
	return &object.Boolean{Value: false}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func evalBlockStatement(block *ast.BlockStatement, env *Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression, env *Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return evalBlockStatement(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return evalBlockStatement(ie.Alternative, env)
	} else {
		return nil
	}
}

/*
isTruthy determines what counts as "true" in Bihari. sach is true. jhoot and kuchnahi (null) are false. Everything else is true!
*/
func isTruthy(obj object.Object) bool {
	switch obj {
	case nil:
		return false
	default:
		if b, ok := obj.(*object.Boolean); ok {
			return b.Value
		}
		if _, ok := obj.(*object.Null); ok {
			return false
		}
		return true
	}
}
