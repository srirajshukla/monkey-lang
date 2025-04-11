package evaluator

import (
	"fmt"
	"gks/monkey_intp/ast"
	"gks/monkey_intp/object"
	"gks/monkey_intp/token"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = evalUnquoteCalls(node, env)
	return &object.Quote{Node: node}
}

func evalUnquoteCalls(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquotedCall(node) {
			return node
		}

		call, ok := node.(*ast.CallExpression)

		if !ok {
			return node
		}
		// check if unquote has only one argument
		if len(call.Arguments) != 1 {
			return node
		}

		unquoted := Eval(call.Arguments[0], env)
		return convertObjectToAstNode(unquoted)
	})
}

func isUnquotedCall(node ast.Node) bool {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return callExpression.Function.TokenLiteral() == "unquote"
}

func convertObjectToAstNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}

	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}

	case *object.Quote:
		return obj.Node

	case *object.String:
		t := token.Token{
			Type:    token.STRING,
			Literal: obj.Value,
		}
		return &ast.StringLiteral{Token: t, Value: obj.Value}

	default:
		return nil
	}
}
