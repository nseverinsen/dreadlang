package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"dreadlang/internal/codegen"
	"dreadlang/internal/lexer"
	"dreadlang/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <source.dread> [output]\n", os.Args[0])
		os.Exit(1)
	}

	sourceFile := os.Args[1]

	// Determine output file name
	outputFile := "a.out"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	// Read source file
	source, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Compile
	if err := compile(string(source), outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Compilation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled %s to %s\n", sourceFile, outputFile)
}

func compile(source string, outputFile string) error {
	// Lexical analysis
	l := lexer.New(source)

	// Syntax analysis
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			fmt.Fprintf(os.Stderr, "Parse error: %s\n", err)
		}
		return fmt.Errorf("parsing failed")
	}

	// Code generation
	cg := codegen.New()
	assembly := cg.Generate(program)

	// Write assembly to temporary file
	asmFile := outputFile + ".s"
	if err := ioutil.WriteFile(asmFile, []byte(assembly), 0644); err != nil {
		return fmt.Errorf("failed to write assembly: %v", err)
	}

	// Assemble and link using system tools
	if err := assembleAndLink(asmFile, outputFile); err != nil {
		return fmt.Errorf("assembly/linking failed: %v", err)
	}

	// Clean up assembly file
	os.Remove(asmFile)

	return nil
}

func assembleAndLink(asmFile, outputFile string) error {
	objFile := strings.TrimSuffix(asmFile, ".s") + ".o"

	// Assemble
	cmd := exec.Command("as", "--64", "-o", objFile, asmFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("assembler error: %v\nOutput: %s", err, output)
	}

	// Link
	cmd = exec.Command("ld", "-o", outputFile, objFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("linker error: %v\nOutput: %s", err, output)
	}

	// Clean up object file
	os.Remove(objFile)

	return nil
}
