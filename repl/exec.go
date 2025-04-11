package repl

import (
	"gks/monkey_intp/evaluator"
	"gks/monkey_intp/lexer"
	"gks/monkey_intp/object"
	"gks/monkey_intp/parser"
)

func EvaluateProgramFromString(prog string) string {
	lex := lexer.NewLexer(prog)
	parser := parser.NewParser(lex)
	program := parser.ParseProgram()

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated.Type() == object.ERROR_OBJ {
		return "Error: " + evaluated.Inspect()
	}

	return evaluated.Inspect()
}
