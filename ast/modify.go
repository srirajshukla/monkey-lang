package ast

import (
	"fmt"
)

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			modified := Modify(statement, modifier)
			stmt, ok := modified.(Statement)
			if !ok {
				return &ErrorExpression{Message: fmt.Sprintf("expected Statement in Program, got %T", modified)}
			}
			node.Statements[i] = stmt
		}

	case *ExpressionStatement:
		modified := Modify(node.Expression, modifier)
		exp, ok := modified.(Expression)
		if !ok {
			return &ErrorExpression{Message: fmt.Sprintf("expected Expression in ExpressionStatement, got %T", modified)}
		}
		node.Expression = exp

	case *InfixExpression:
		left := Modify(node.Left, modifier)
		right := Modify(node.Right, modifier)
		lExp, ok1 := left.(Expression)
		rExp, ok2 := right.(Expression)
		if !ok1 || !ok2 {
			return &ErrorExpression{Message: "invalid Expression in InfixExpression"}
		}
		node.Left = lExp
		node.Right = rExp

	case *PrefixExpression:
		right := Modify(node.Right, modifier)
		rExp, ok := right.(Expression)
		if !ok {
			return &ErrorExpression{Message: "invalid Expression in PrefixExpression"}
		}
		node.Right = rExp

	case *IndexExpression:
		left := Modify(node.Left, modifier)
		index := Modify(node.Index, modifier)
		lExp, ok1 := left.(Expression)
		iExp, ok2 := index.(Expression)
		if !ok1 || !ok2 {
			return &ErrorExpression{Message: "invalid Expression in IndexExpression"}
		}
		node.Left = lExp
		node.Index = iExp

	case *IfExpression:
		cond := Modify(node.Condition, modifier)
		conseq := Modify(node.Consequence, modifier)
		cExp, ok1 := cond.(Expression)
		bs1, ok2 := conseq.(*BlockStatement)
		if !ok1 || !ok2 {
			return &ErrorExpression{Message: "invalid Condition or Consequence in IfExpression"}
		}
		node.Condition = cExp
		node.Consequence = bs1

		if node.Alternative != nil {
			alt := Modify(node.Alternative, modifier)
			bs2, ok := alt.(*BlockStatement)
			if !ok {
				return &ErrorExpression{Message: "invalid Alternative in IfExpression"}
			}
			node.Alternative = bs2
		}

	case *BlockStatement:
		for i := range node.Statements {
			modified := Modify(node.Statements[i], modifier)
			stmt, ok := modified.(Statement)
			if !ok {
				return &ErrorExpression{Message: fmt.Sprintf("expected Statement in BlockStatement, got %T", modified)}
			}
			node.Statements[i] = stmt
		}

	case *ReturnStatement:
		modified := Modify(node.ReturnValue, modifier)
		exp, ok := modified.(Expression)
		if !ok {
			return &ErrorExpression{Message: "invalid ReturnValue in ReturnStatement"}
		}
		node.ReturnValue = exp

	case *LetStatement:
		modified := Modify(node.Value, modifier)
		exp, ok := modified.(Expression)
		if !ok {
			return &ErrorExpression{Message: "invalid Value in LetStatement"}
		}
		node.Value = exp

	case *FunctionLiteral:
		for i := range node.Parameters {
			modified := Modify(node.Parameters[i], modifier)
			ident, ok := modified.(*Identifier)
			if !ok {
				return &ErrorExpression{Message: "invalid Parameter in FunctionLiteral"}
			}
			node.Parameters[i] = ident
		}
		modified := Modify(node.Body, modifier)
		body, ok := modified.(*BlockStatement)
		if !ok {
			return &ErrorExpression{Message: "invalid Body in FunctionLiteral"}
		}
		node.Body = body

	case *ArrayLiteral:
		for i := range node.Elements {
			modified := Modify(node.Elements[i], modifier)
			elem, ok := modified.(Expression)
			if !ok {
				return &ErrorExpression{Message: fmt.Sprintf("invalid Element in ArrayLiteral: got %T", modified)}
			}
			node.Elements[i] = elem
		}

	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKeyNode := Modify(key, modifier)
			newValNode := Modify(val, modifier)
			newKey, ok1 := newKeyNode.(Expression)
			newVal, ok2 := newValNode.(Expression)
			if !ok1 || !ok2 {
				return &ErrorExpression{Message: "invalid key or value in HashLiteral"}
			}
			newPairs[newKey] = newVal
		}
		node.Pairs = newPairs
	}

	return modifier(node)
}
