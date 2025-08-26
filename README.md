
# Dread Programming Language

Dread is a general-purpose programming language with a Go-based compiler that generates native x86-64 machine code. The language features a clean syntax with uppercase keywords and supports basic programming constructs like functions, variables, and built-in I/O operations.

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or later
- GNU Assembler (`as`) and Linker (`ld`)
- Linux x86-64 system

### Building the Compiler
```bash
git clone <repository-url>
cd dreadlang
go build -o dreadc ./cmd/dreadc
```

### Your First Dread Program
Create a file called `hello.dread`:

```dread
// This is a comment!

/*
 * This is a docstring for the entrypoint (Entry), which is the start of the executable program.
 * It has no arguments, for now. It returns integer by default
 */
Entry main() (Int)
{
    /* Variable assignment with duck typing */
    hello_string = 'Hello, World!
'

    Print(hello_string) // Print to stdout
    Return(0) // Exit with status 0
}
```

Compile and run:
```bash
./dreadc hello.dread hello
./hello
```

## 📝 Language Syntax

### Keywords
All keywords in Dread start with an uppercase letter:
- `Entry` - Entry point function declaration (special function)
- `Function` - Regular function declaration keyword
- `Print` - Built-in print function
- `Return` - Return statement
- `Int` - Integer type annotation

### Comments
```dread
// Single-line comment

/*
 * Multi-line comment
 * Can span multiple lines
 */
```

### Functions

**Entry Point Function (Program Entry)**:
```dread
Entry main() (ReturnType)
{
    // Entry point function body - where program execution begins
}
```

**Regular Functions (Future Feature)**:
```dread
Function functionName() (ReturnType)
{
    // Regular function body
}
```

*Note: Currently only Entry functions are implemented. Regular Function support is planned.*

### Variables
Variables use duck typing - no explicit type declaration needed:
```dread
variable_name = 'String value'
number_var = 42
```

### Built-in Functions
- `Print(value)` - Print to stdout
- `Return(code)` - Exit program with status code

## 🏗️ Architecture

The Dread compiler follows a traditional three-phase compilation pipeline:

### 1. Lexical Analysis (`internal/lexer/`)
The lexer tokenizes the source code into a stream of tokens:
- **Keywords**: `Entry`, `Print`, `Return`, `Int`
- **Identifiers**: Variable and function names
- **Literals**: Strings (`'text'`) and integers (`123`)
- **Operators**: Assignment (`=`)
- **Delimiters**: `()`, `{}`, etc.
- **Comments**: Both `//` and `/* */` styles

### 2. Syntax Analysis (`internal/parser/`)
The parser builds an Abstract Syntax Tree (AST) from tokens:
- **Program**: Root node containing all statements
- **Function Statements**: `Entry` function declarations
- **Assignment Statements**: Variable assignments
- **Call Statements**: Function calls like `Print()` and `Return()`
- **Block Statements**: Code blocks within `{}`

### 3. Code Generation (`internal/codegen/`)
The code generator produces x86-64 assembly:
- **Data Section**: String constants and their lengths
- **Text Section**: Executable code using Linux system calls
- **System Calls**: `sys_write` for printing, `sys_exit` for program termination

### 4. Assembly and Linking (`cmd/dreadc/`)
The compiler driver coordinates the entire process:
1. Read source file
2. Tokenize → Parse → Generate assembly
3. Invoke system assembler (`as`) to create object file
4. Invoke system linker (`ld`) to create executable

## 📁 Project Structure

```
dreadlang/
├── README.md                 # This file
├── TODO.md                   # Development roadmap
├── MILESTONES.md            # Development milestones and progress tracking
├── SPECIFICATION.md         # Language specification
├── ARCHITECTURE.md          # Technical architecture documentation
├── DEVELOPER.md             # Developer guide
├── go.mod                   # Go module definition
├── .gitignore              # Git ignore rules
├── cmd/
│   └── dreadc/
│       └── main.go          # Compiler main entry point
├── internal/
│   ├── lexer/
│   │   └── lexer.go         # Lexical analyzer
│   ├── parser/
│   │   └── parser.go        # Syntax analyzer and AST
│   └── codegen/
│       └── codegen.go       # x86-64 assembly generator
└── examples/
    ├── hello.dread          # Hello world with comments
    └── hello_simple.dread   # Minimal hello world
```

## 🔧 Compiler Usage

```bash
./dreadc <source_file.dread> [output_executable]
```

**Examples:**
```bash
# Compile to default output (a.out)
./dreadc examples/hello.dread

# Compile to specific executable name
./dreadc examples/hello.dread my_program

# Run the compiled program
./my_program
```

## 🧪 Current Capabilities

### ✅ Implemented Features
- **Lexical Analysis**: Complete tokenization of Dread syntax
- **Parser**: AST generation for basic language constructs
- **Code Generation**: x86-64 assembly output
- **String Literals**: Single-quoted strings with newline support
- **Comments**: Both single-line and multi-line
- **Function Declarations**: Entry point functions
- **Variable Assignment**: Duck-typed variables
- **Built-in I/O**: Print function for stdout
- **Program Control**: Return statements with exit codes

### 🚧 Current Limitations
- **Single Entry Point**: Exactly one `Entry main()` function per program
- **No regular functions**: Only Entry functions supported (see [MILESTONES.md](MILESTONES.md))
- **No arithmetic expressions**: Mathematical operations not yet implemented
- **No control flow**: No if/else statements or loops yet
- **Limited type system**: Only strings and integers
- **No function parameters**: Functions take no arguments
- **No error handling**: Basic error reporting only
- **Single-file compilation**: No module system yet

**Next Milestone**: Function support - see [MILESTONES.md](MILESTONES.md) for development roadmap.

## 🔍 Technical Details

### Memory Layout
- **Stack**: Function calls and local variables
- **Data Section**: String constants and static data
- **Text Section**: Executable code

### System Calls Used
- `sys_write` (1): For Print() function output
- `sys_exit` (60): For Return() program termination

### Assembly Generation
The compiler generates AT&T syntax assembly with Intel mnemonics:
```assembly
.intel_syntax noprefix
.global _start

.section .data
str_0: .ascii "Hello, World!\n"
str_0_len = . - str_0

.section .text
_start:
    mov rax, 1           # sys_write
    mov rdi, 1           # stdout
    lea rsi, [str_0]     # string address
    mov rdx, str_0_len   # string length
    syscall

    mov rax, 60          # sys_exit
    mov rdi, 0           # exit status
    syscall
```

## 🤝 Contributing

This is an active development project. See `TODO.md` for the development roadmap and current priorities.

### Development Setup
1. Install Go 1.21+
2. Clone the repository
3. Run `go build -o dreadc ./cmd/dreadc`
4. Test with `./dreadc examples/hello.dread test && ./test`

## 📋 License

[Specify your license here]

---

*Dread is under active development. The language specification and compiler implementation are subject to change.*
