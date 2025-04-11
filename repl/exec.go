package repl

import (
	"fmt"
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
	macroEnv := object.NewEnvironment()

	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	fmt.Println(expanded.String())

	evaluated := evaluator.Eval(expanded, env)

	if evaluated.Type() == object.ERROR_OBJ {
		return "Error: " + evaluated.Inspect()
	}

	return evaluated.Inspect()
}
