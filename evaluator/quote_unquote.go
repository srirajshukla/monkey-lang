package evaluator

import "gks/monkey_intp/object"
import "gks/monkey_intp/ast"

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
