package main

import (
	"dreadlang/internal/lexer"
	"dreadlang/internal/parser"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dread-file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Shows debug information for a Dread source file (tokens, AST, etc.)\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("=== DEBUGGING: %s ===\n\n", filename)

	// Show source
	fmt.Println("=== SOURCE ===")
	fmt.Print(string(source))
	fmt.Println()

	// Tokenize and show tokens
	fmt.Println("=== TOKENS ===")
	l := lexer.New(string(source))
	for {
		tok := l.NextToken()
		if tok.Type == lexer.EOF {
			fmt.Printf("Token: %s\n", tok.Type.String())
			break
		}
		fmt.Printf("Token: %s, Literal: %q\n", tok.Type.String(), tok.Literal)
	}
	fmt.Println()

	// Parse and show AST
	fmt.Println("=== AST ===")
	l2 := lexer.New(string(source))
	p := parser.New(l2)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println("Parse errors:")
		for _, err := range p.Errors() {
			fmt.Printf("  %s\n", err)
		}
		fmt.Println()
	}

	fmt.Printf("AST: %s\n", program.String())
}
