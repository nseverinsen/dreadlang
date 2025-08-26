# Dread Compiler Architecture

This document provides a detailed technical overview of how the Dread programming language compiler works.

## Overview

The Dread compiler is implemented in Go and follows a traditional multi-pass compilation architecture:

```
Source Code (.dread) → Lexer → Parser → CodeGen → Assembler → Linker → Executable
```

## Phase 1: Lexical Analysis

**File**: `internal/lexer/lexer.go`

The lexer performs tokenization of the source code, converting the character stream into a sequence of meaningful tokens.

### Token Types

```go
type TokenType int

const (
    // Special tokens
    ILLEGAL TokenType = iota
    EOF

    // Identifiers and literals
    IDENT  // variable names
    STRING // 'hello world'
    INT    // 123

    // Keywords
    ENTRY  // Entry
    PRINT  // Print
    RETURN // Return
    INT_TYPE // Int

    // Delimiters
    LPAREN // (
    RPAREN // )
    LBRACE // {
    RBRACE // }

    // Operators
    ASSIGN // =
)
```

### Key Components

1. **Character Reading**: The lexer reads characters one by one, maintaining position and line/column information for error reporting.

2. **Comment Handling**: Supports both single-line (`//`) and multi-line (`/* */`) comments, which are skipped during tokenization.

3. **String Parsing**: Handles single-quoted strings with basic escape sequence support.

4. **Keyword Recognition**: Uses a lookup table to distinguish keywords from identifiers.

### Example Token Stream

For the input:
```dread
Entry main() (Int) { hello = 'World' }
```

The lexer produces:
```
ENTRY("Entry") → IDENT("main") → LPAREN → RPAREN → LPAREN → INT_TYPE("Int") → RPAREN → LBRACE → IDENT("hello") → ASSIGN → STRING("World") → RBRACE → EOF
```

## Phase 2: Syntax Analysis

**File**: `internal/parser/parser.go`

The parser takes the token stream and builds an Abstract Syntax Tree (AST) representing the program structure.

### AST Node Types

```go
// Base interfaces
type Node interface {
    String() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}
```

### Key AST Nodes

1. **Program**: Root node containing all top-level statements
2. **FunctionStatement**: Represents `Entry` function declarations
3. **BlockStatement**: Code blocks within `{}`
4. **AssignStatement**: Variable assignments (`x = value`)
5. **CallStatement**: Function calls (`Print(x)`, `Return(0)`)
6. **StringLiteral**: String values
7. **Identifier**: Variable references

### Parsing Strategy

The parser uses recursive descent parsing:

1. **parseStatement()**: Handles top-level constructs (functions)
2. **parseBlockStatement()**: Processes code within `{}`
3. **parseInnerStatement()**: Handles assignments and function calls
4. **parseExpression()**: Processes values and identifiers

### Example AST

For the hello world program:
```
Program
└── FunctionStatement(name="main", returnType="Int")
    └── BlockStatement
        ├── AssignStatement(name="hello_string", value=StringLiteral("Hello, World!\n"))
        ├── CallStatement(function="Print", args=[Identifier("hello_string")])
        └── CallStatement(function="Return", args=[StringLiteral("0")])
```

## Phase 3: Code Generation

**File**: `internal/codegen/codegen.go`

The code generator traverses the AST and produces x86-64 assembly code for Linux.

### Assembly Structure

Generated assembly follows this structure:

```assembly
.intel_syntax noprefix    # Use Intel syntax
.global _start           # Entry point

.section .data           # Static data
str_0: .ascii "Hello"
str_0_len = . - str_0

.section .text           # Executable code
_start:
    # Program code here
```

### Code Generation Process

1. **String Collection**: First pass collects all string literals and assigns labels
2. **Data Section**: Generates string constants with length calculations
3. **Text Section**: Generates executable code for the main function

### System Call Interface

The compiler uses Linux system calls directly:

- **sys_write (1)**: For Print() function
  ```assembly
  mov rax, 1           # sys_write
  mov rdi, 1           # stdout
  lea rsi, [str_label] # string address
  mov rdx, str_len     # string length
  syscall
  ```

- **sys_exit (60)**: For Return() function
  ```assembly
  mov rax, 60          # sys_exit
  mov rdi, exit_code   # exit status
  syscall
  ```

### Variable Management

Currently, the code generator uses a simple variable mapping approach:
- Variables are mapped to string constant labels
- No stack allocation yet (planned for future versions)

## Phase 4: Assembly and Linking

**File**: `cmd/dreadc/main.go`

The compiler driver orchestrates the entire compilation process and invokes system tools.

### Compilation Pipeline

1. **Source Reading**: Read the `.dread` source file
2. **Lexical Analysis**: Create lexer and tokenize
3. **Syntax Analysis**: Create parser and build AST
4. **Code Generation**: Generate assembly code
5. **Assembly**: Invoke `as --64` to create object file
6. **Linking**: Invoke `ld` to create executable
7. **Cleanup**: Remove intermediate files

### Command Line Interface

```bash
./dreadc <source.dread> [output_name]
```

The compiler:
- Reads the source file
- Compiles to assembly (temporary `.s` file)
- Assembles to object code (temporary `.o` file)
- Links to final executable
- Cleans up temporary files

### Error Handling

The compiler includes basic error handling:
- File I/O errors
- Parse errors with location information
- Assembly/linking errors from system tools

## Memory Model

### Current Implementation

- **Stack**: Managed by the OS for function calls
- **Data**: String constants in data section
- **Text**: Executable code in text section
- **Variables**: Currently mapped to static string references

### Future Enhancements

- Dynamic memory allocation
- Proper stack frame management
- Heap allocation for complex data structures

## Debugging and Development

### Assembly Output

To inspect generated assembly, comment out the cleanup in `main.go`:
```go
// os.Remove(asmFile) // Keep assembly for debugging
```

### Token Debugging

Create a simple token printer:
```go
lexer := lexer.New(source)
for {
    tok := lexer.NextToken()
    if tok.Type == lexer.EOF { break }
    fmt.Printf("Token: %s, Literal: %q\n", tok.Type, tok.Literal)
}
```

### AST Debugging

The AST nodes implement `String()` methods for debugging:
```go
program := parser.ParseProgram()
fmt.Println("AST:", program.String())
```

## Performance Characteristics

### Compilation Speed
- Single-pass lexing and parsing
- Direct assembly generation (no intermediate representation yet)
- Fast for small programs (< 1ms for hello world)

### Generated Code
- Minimal runtime overhead
- Direct system calls (no C library dependency)
- Small executable size (~9KB for hello world)

## Limitations and Future Work

### Current Limitations
1. No arithmetic expressions
2. No control flow statements
3. Single-file compilation only
4. No function parameters
5. Limited type system

### Planned Improvements
1. Expression parsing and evaluation
2. If/else and loop constructs
3. Multi-file compilation and linking
4. Function parameters and local variables
5. Advanced type checking and inference

See `TODO.md` for detailed development roadmap.

---

This architecture provides a solid foundation for a programming language compiler while remaining simple enough for educational purposes and further development.
