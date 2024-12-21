//go:build wasm
package main

import (
	"fmt"
	"gks/monkey_intp/repl"
	"syscall/js"
)

func calculate(prog string) string {
	return repl.EvaluateProgramFromString(prog)
}

func evalProgram(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(calculate(args[0].String()))
}

func registerCallbacks() {
	js.Global().Set("evaluateProgram", js.FuncOf(evalProgram))
}


func main(){	
	fmt.Println("printing this from golang")
	c := make(chan struct{}, 0)
	fmt.Println("Go wasm initialized")
	registerCallbacks()
	<-c
}