package main

import (
	"fmt"
	"gks/monkey_intp/repl"
	"os"
)

func main() {
	fmt.Printf("Hello! This is the Monkey programming language REPL!\n")
	fmt.Printf("Feel free to type in commands\n")

	args := os.Args[1:]
	fmt.Println("args = ", args)

	if len(args) < 1 {
		fmt.Printf("Wrong args. monkey run filename | monkey repl")
		os.Exit(1)
	}

	if args[0] == "repl" {
		repl.Start(os.Stdin, os.Stdout)
	} else if args[0] == "run" {
		if len(args) < 2 {
			fmt.Println("need a file name after run. Usage: monkey run filename")
			os.Exit(1)
		}

		file, err := os.ReadFile(args[1])

		if err != nil {
			fmt.Printf("Error opening the file %s. Error: %v\n", args[1], err)
			os.Exit(1)
		}

		evaluatedValue := repl.EvaluateProgramFromString(string(file))
		fmt.Println(evaluatedValue)

	}

}
