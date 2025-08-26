# Dread Developer Guide

This guide helps developers understand, modify, and extend the Dread programming language compiler.

## Getting Started

### Development Environment Setup

1. **Prerequisites**:
   ```bash
   # Go 1.21 or later
   go version

   # GNU binutils (assembler and linker)
   which as ld

   # Git for version control
   git --version
   ```

2. **Clone and Build**:
   ```bash
   git clone <repository-url>
   cd dreadlang
   go build -o dreadc ./cmd/dreadc
   ```

3. **Test the Installation**:
   ```bash
   ./dreadc examples/hello.dread test_program
   ./test_program
   # Should output: Hello, World!
   rm test_program
   ```

## Code Organization

### Module Structure

```
internal/
├── lexer/      # Tokenization (characters → tokens)
├── parser/     # Syntax analysis (tokens → AST)
└── codegen/    # Code generation (AST → assembly)

cmd/
└── dreadc/     # Compiler driver (main application)

examples/       # Example Dread programs
```

### Key Files

- `internal/lexer/lexer.go`: Lexical analyzer implementation
- `internal/parser/parser.go`: Parser and AST definitions
- `internal/codegen/codegen.go`: Assembly code generator
- `cmd/dreadc/main.go`: Main compiler driver

## Adding New Features

### 1. Adding a New Token Type

**Example**: Adding support for a `+` operator

1. **Define the token** in `lexer.go`:
   ```go
   const (
       // ... existing tokens
       PLUS // +
   )
   ```

2. **Add to lexer switch statement**:
   ```go
   switch l.ch {
   // ... existing cases
   case '+':
       tok = Token{Type: PLUS, Literal: string(l.ch), Line: l.line, Column: l.column}
   }
   ```

3. **Add to String() method**:
   ```go
   func (t TokenType) String() string {
       switch t {
       // ... existing cases
       case PLUS:
           return "PLUS"
       }
   }
   ```

### 2. Adding a New AST Node

**Example**: Adding an addition expression

1. **Define the node type** in `parser.go`:
   ```go
   type InfixExpression struct {
       Left     Expression
       Operator string
       Right    Expression
   }

   func (ie *InfixExpression) expressionNode() {}
   func (ie *InfixExpression) String() string {
       return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
   }
   ```

2. **Add parsing logic**:
   ```go
   func (p *Parser) parseInfixExpression(left Expression) Expression {
       expression := &InfixExpression{
           Left:     left,
           Operator: p.curToken.Literal,
       }

       precedence := p.curPrecedence()
       p.nextToken()
       expression.Right = p.parseExpression(precedence)

       return expression
   }
   ```

### 3. Adding Code Generation Support

**Example**: Generating code for addition

1. **Add to code generator** in `codegen.go`:
   ```go
   func (cg *CodeGenerator) generateExpression(expr parser.Expression) {
       switch e := expr.(type) {
       // ... existing cases
       case *parser.InfixExpression:
           cg.generateInfixExpression(e)
       }
   }

   func (cg *CodeGenerator) generateInfixExpression(ie *parser.InfixExpression) {
       // Generate code for left operand
       cg.generateExpression(ie.Left)
       cg.output.WriteString("    push rax\n") // Save left result

       // Generate code for right operand
       cg.generateExpression(ie.Right)
       cg.output.WriteString("    pop rbx\n")  // Restore left result

       switch ie.Operator {
       case "+":
           cg.output.WriteString("    add rax, rbx\n")
       }
   }
   ```

## Testing and Debugging

### Manual Testing

1. **Create test programs** in the `examples/` directory:
   ```dread
   Entry main() (Int)
   {
       result = 5 + 3
       Print(result)
       Return(0)
   }
   ```

2. **Compile and test**:
   ```bash
   ./dreadc examples/test_addition.dread test_add
   ./test_add
   ```

### Debugging Techniques

1. **Token Stream Debugging**:
   ```go
   // In main.go, after creating lexer
   for {
       tok := l.NextToken()
       if tok.Type == lexer.EOF { break }
       fmt.Printf("Token: %s, Literal: %q\n", tok.Type, tok.Literal)
   }
   ```

2. **AST Debugging**:
   ```go
   // In main.go, after parsing
   fmt.Printf("AST: %s\n", program.String())
   ```

3. **Assembly Debugging**:
   ```go
   // In main.go, comment out assembly file removal
   // os.Remove(asmFile)

   // Then inspect the generated assembly
   cat output.s
   ```

4. **Runtime Debugging**:
   ```bash
   # Use objdump to inspect the executable
   objdump -d executable_name

   # Use strace to see system calls (if available)
   strace ./executable_name
   ```

### Unit Testing

Create unit tests for each component:

```go
// internal/lexer/lexer_test.go
func TestTokenizeBasic(t *testing.T) {
    input := "Entry main() (Int)"
    lexer := New(input)

    tests := []struct {
        expectedType    TokenType
        expectedLiteral string
    }{
        {ENTRY, "Entry"},
        {IDENT, "main"},
        {LPAREN, "("},
        // ... more test cases
    }

    for _, tt := range tests {
        tok := lexer.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("expected %q, got %q", tt.expectedType, tok.Type)
        }
    }
}
```

## Common Development Patterns

### Error Handling

1. **Lexer Errors**:
   ```go
   if l.ch == 0 && expectedMore {
       return Token{Type: ILLEGAL, Literal: "unexpected EOF"}
   }
   ```

2. **Parser Errors**:
   ```go
   func (p *Parser) peekError(t TokenType) {
       msg := fmt.Sprintf("expected next token to be %s, got %s instead",
           t, p.peekToken.Type)
       p.errors = append(p.errors, msg)
   }
   ```

3. **Codegen Errors**:
   ```go
   if !cg.isValidExpression(expr) {
       return fmt.Errorf("cannot generate code for expression: %T", expr)
   }
   ```

### Adding Built-in Functions

1. **Add to lexer keywords**:
   ```go
   var keywords = map[string]TokenType{
       // ... existing keywords
       "Println": PRINTLN,
   }
   ```

2. **Handle in parser**:
   ```go
   case lexer.PRINTLN:
       return p.parseCallStatement()
   ```

3. **Generate code**:
   ```go
   case "Println":
       cg.generatePrint(stmt.Arguments[0])
       cg.output.WriteString("    # Add newline\n")
       cg.output.WriteString("    mov rax, 1\n")
       cg.output.WriteString("    mov rdi, 1\n")
       cg.output.WriteString("    lea rsi, [newline]\n")
       cg.output.WriteString("    mov rdx, 1\n")
       cg.output.WriteString("    syscall\n")
   ```

## Performance Considerations

### Compilation Speed

- **Avoid unnecessary allocations** in the lexer hot path
- **Use string builders** for assembly generation
- **Cache frequently accessed data** (like string constants)

### Generated Code Quality

- **Use appropriate registers** for different data types
- **Minimize memory accesses** when possible
- **Generate efficient system call sequences**

### Memory Usage

- **Reuse token and AST node objects** when possible
- **Clean up temporary files** after compilation
- **Avoid keeping entire source in memory** for large files (future enhancement)

## Extension Points

### Adding New Data Types

1. Add token types for type keywords
2. Extend the type system in the parser
3. Add type checking in semantic analysis
4. Generate appropriate assembly for each type

### Adding Control Flow

1. Add tokens for keywords (`If`, `Else`, `While`, etc.)
2. Create AST nodes for control structures
3. Implement parsing logic with proper precedence
4. Generate conditional jumps and labels in assembly

### Adding Function Parameters

1. Extend function declaration parsing
2. Add parameter lists to AST nodes
3. Implement calling convention in code generation
4. Add stack frame management

## Troubleshooting

### Common Issues

1. **"Parse error" messages**:
   - Check token definitions in lexer
   - Verify parser expects correct token sequence
   - Add debug output to see token stream

2. **"Assembly/linking failed"**:
   - Check generated assembly syntax
   - Verify system tools are available
   - Inspect intermediate files

3. **Program crashes or produces no output**:
   - Check system call arguments
   - Verify string constants are properly defined
   - Use debugging tools like `objdump`

### Getting Help

1. Check the `TODO.md` for known limitations
2. Review `ARCHITECTURE.md` for design details
3. Look at existing examples in `examples/`
4. Add debug output to trace execution

## Contributing Guidelines

1. **Follow Go conventions**: Use `gofmt`, proper naming, documentation
2. **Test thoroughly**: Add test cases for new features
3. **Update documentation**: Keep README and architecture docs current
4. **Small incremental changes**: Break large features into smaller PRs
5. **Maintain backward compatibility**: Don't break existing examples

---

This guide should help you get started with Dread compiler development. Remember that this is an evolving project, so patterns and practices may change as the language grows!
