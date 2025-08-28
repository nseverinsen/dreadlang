package main

import (
	"dreadlang/internal/codegen"
	"dreadlang/internal/lexer"
	"dreadlang/internal/parser"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dread-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Shows the generated assembly for a Dread source file\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	l := lexer.New(string(source))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Fprintf(os.Stderr, "Parse errors:\n")
		for _, err := range p.Errors() {
			fmt.Fprintf(os.Stderr, "  %s\n", err)
		}
		os.Exit(1)
	}

	cg := codegen.New()
	assembly := cg.Generate(program)
	fmt.Print(assembly)
}
