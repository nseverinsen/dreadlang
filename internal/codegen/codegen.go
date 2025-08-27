package codegen

import (
	"dreadlang/internal/parser"
	"fmt"
	"strings"
)

type CodeGenerator struct {
	output          strings.Builder
	stringConstants map[string]string
	stringCounter   int
}

func New() *CodeGenerator {
	return &CodeGenerator{
		stringConstants: make(map[string]string),
		stringCounter:   0,
	}
}

func (cg *CodeGenerator) Generate(program *parser.Program) string {
	cg.output.Reset()

	// Generate assembly header
	cg.writeHeader()

	// Generate string constants
	cg.writeDataSection(program)

	// Generate code section
	cg.writeTextSection(program)

	return cg.output.String()
}

func (cg *CodeGenerator) writeHeader() {
	cg.output.WriteString(".intel_syntax noprefix\n")
	cg.output.WriteString(".global _start\n\n")
}

func (cg *CodeGenerator) writeDataSection(program *parser.Program) {
	cg.output.WriteString(".section .data\n")

	// Collect all string literals
	cg.collectStrings(program)

	// Generate string constants
	for literal, label := range cg.stringConstants {
		// Convert escape sequences
		processed := cg.processString(literal)
		cg.output.WriteString(fmt.Sprintf("%s: .ascii \"%s\"\n", label, processed))
		cg.output.WriteString(fmt.Sprintf("%s_len = . - %s\n", label, label))
	}

	cg.output.WriteString("\n")
}

func (cg *CodeGenerator) writeTextSection(program *parser.Program) {
	cg.output.WriteString(".section .text\n")

	// Find and generate the Entry function first
	var entryFound bool
	for _, stmt := range program.Statements {
		if funcStmt, ok := stmt.(*parser.FunctionStatement); ok {
			if funcStmt.IsEntry {
				cg.output.WriteString("_start:\n")
				cg.generateFunction(funcStmt)
				entryFound = true
				break
			}
		}
	}

	if !entryFound {
		// Default entry point if no Entry function found
		cg.output.WriteString("_start:\n")
		cg.output.WriteString("    # No Entry function found\n")
		cg.output.WriteString("    mov rax, 60      # sys_exit\n")
		cg.output.WriteString("    mov rdi, 1       # exit status\n")
		cg.output.WriteString("    syscall\n")
	}

	// Generate all regular functions
	for _, stmt := range program.Statements {
		if funcStmt, ok := stmt.(*parser.FunctionStatement); ok {
			if !funcStmt.IsEntry {
				cg.generateFunction(funcStmt)
			}
		}
	}
}

func (cg *CodeGenerator) generateBlockStatement(block *parser.BlockStatement, isEntry bool) {
	variables := make(map[string]string) // variable name -> label/register

	for _, stmt := range block.Statements {
		switch s := stmt.(type) {
		case *parser.AssignStatement:
			cg.generateAssignStatement(s, variables)
		case *parser.CallStatement:
			cg.generateCallStatement(s, variables, isEntry)
		}
	}
}

func (cg *CodeGenerator) generateAssignStatement(stmt *parser.AssignStatement, variables map[string]string) {
	switch expr := stmt.Value.(type) {
	case *parser.StringLiteral:
		// Store reference to string constant
		label := cg.getStringLabel(expr.Value)
		variables[stmt.Name] = label
	case *parser.Identifier:
		// Copy variable reference
		if ref, exists := variables[expr.Value]; exists {
			variables[stmt.Name] = ref
		}
	case *parser.CallExpression:
		// Function call assignment - for now, just call the function
		// TODO: Implement proper return value handling
		cg.output.WriteString(fmt.Sprintf("    # %s = %s()\n", stmt.Name, expr.Function))
		if len(expr.Arguments) == 0 {
			cg.output.WriteString(fmt.Sprintf("    call %s\n", expr.Function))
		} else {
			cg.output.WriteString("    # TODO: Function calls with parameters not yet implemented\n")
		}
		// For now, we don't capture the return value
	}
}

func (cg *CodeGenerator) generateCallStatement(stmt *parser.CallStatement, variables map[string]string, isEntry bool) {
	switch stmt.Function {
	case "Print":
		if len(stmt.Arguments) > 0 {
			arg := stmt.Arguments[0]
			switch a := arg.(type) {
			case *parser.Identifier:
				if label, exists := variables[a.Value]; exists {
					cg.generatePrint(label)
				}
			case *parser.StringLiteral:
				label := cg.getStringLabel(a.Value)
				cg.generatePrint(label)
			}
		}
	case "Return":
		if len(stmt.Arguments) > 0 {
			switch a := stmt.Arguments[0].(type) {
			case *parser.StringLiteral:
				if isEntry {
					// Entry function: exit the program
					exitCode := a.Value
					cg.output.WriteString(fmt.Sprintf("    # Return(%s)\n", exitCode))
					cg.output.WriteString("    mov rax, 60      # sys_exit\n")
					cg.output.WriteString(fmt.Sprintf("    mov rdi, %s      # exit status\n", exitCode))
					cg.output.WriteString("    syscall\n")
				} else {
					// Regular function: return value (for now, just return to caller)
					cg.output.WriteString(fmt.Sprintf("    # Return(%s)\n", a.Value))
					cg.output.WriteString("    mov rsp, rbp\n")
					cg.output.WriteString("    pop rbp\n")
					cg.output.WriteString("    ret\n")
				}
			}
		}
	default:
		// User-defined function call
		cg.output.WriteString(fmt.Sprintf("    # Call %s\n", stmt.Function))

		// For now, we'll support simple function calls without parameters
		// TODO: Handle function parameters and return values
		if len(stmt.Arguments) == 0 {
			cg.output.WriteString(fmt.Sprintf("    call %s\n", stmt.Function))
		} else {
			cg.output.WriteString("    # TODO: Function calls with parameters not yet implemented\n")
		}
	}
}

func (cg *CodeGenerator) generatePrint(label string) {
	cg.output.WriteString(fmt.Sprintf("    # Print(%s)\n", label))
	cg.output.WriteString("    mov rax, 1       # sys_write\n")
	cg.output.WriteString("    mov rdi, 1       # stdout\n")
	cg.output.WriteString(fmt.Sprintf("    lea rsi, [%s]    # string address\n", label))
	cg.output.WriteString(fmt.Sprintf("    mov rdx, %s_len  # string length\n", label))
	cg.output.WriteString("    syscall\n")
}

func (cg *CodeGenerator) collectStrings(program *parser.Program) {
	for _, stmt := range program.Statements {
		cg.collectStringsFromStatement(stmt)
	}
}

func (cg *CodeGenerator) collectStringsFromStatement(stmt parser.Statement) {
	switch s := stmt.(type) {
	case *parser.FunctionStatement:
		cg.collectStringsFromStatement(s.Body)
	case *parser.BlockStatement:
		for _, innerStmt := range s.Statements {
			cg.collectStringsFromStatement(innerStmt)
		}
	case *parser.AssignStatement:
		cg.collectStringsFromExpression(s.Value)
	case *parser.CallStatement:
		for _, arg := range s.Arguments {
			cg.collectStringsFromExpression(arg)
		}
	}
}

func (cg *CodeGenerator) collectStringsFromExpression(expr parser.Expression) {
	switch e := expr.(type) {
	case *parser.StringLiteral:
		cg.getStringLabel(e.Value)
	}
}

func (cg *CodeGenerator) getStringLabel(literal string) string {
	if label, exists := cg.stringConstants[literal]; exists {
		return label
	}

	label := fmt.Sprintf("str_%d", cg.stringCounter)
	cg.stringConstants[literal] = label
	cg.stringCounter++
	return label
}

func (cg *CodeGenerator) processString(s string) string {
	// Handle basic escape sequences
	s = strings.ReplaceAll(s, "\\n", "\\n")
	s = strings.ReplaceAll(s, "\\t", "\\t")
	s = strings.ReplaceAll(s, "\\r", "\\r")
	s = strings.ReplaceAll(s, "\\\\", "\\\\")
	s = strings.ReplaceAll(s, "\\\"", "\\\"")
	return s
}

func (cg *CodeGenerator) generateFunction(funcStmt *parser.FunctionStatement) {
	if !funcStmt.IsEntry {
		// Generate function label
		cg.output.WriteString(fmt.Sprintf("%s:\n", funcStmt.Name))

		// Set up stack frame for regular functions
		cg.output.WriteString("    push rbp\n")
		cg.output.WriteString("    mov rbp, rsp\n")
	}

	// Generate function body
	cg.generateBlockStatement(funcStmt.Body, funcStmt.IsEntry)

	if !funcStmt.IsEntry {
		// Default return for regular functions
		cg.output.WriteString("    # Default function return\n")
		cg.output.WriteString("    mov rsp, rbp\n")
		cg.output.WriteString("    pop rbp\n")
		cg.output.WriteString("    ret\n")
	} else {
		// Default exit for Entry function
		cg.output.WriteString("    # Default exit\n")
		cg.output.WriteString("    mov rax, 60      # sys_exit\n")
		cg.output.WriteString("    mov rdi, 0       # exit status\n")
		cg.output.WriteString("    syscall\n")
	}
}
