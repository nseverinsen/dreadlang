# Dread Programming Language - Development TODO

## Overview
Dread is a general-purpose programming language with a Go-based compiler that outputs executable machine code. This TODO outlines the development roadmap for implementing the language.

## Phase 1: Core Language Infrastructure

### 1.1 Lexical Analysis (Tokenizer/Lexer)
- [x] Implement tokenizer for basic tokens:
  - [x] Keywords (Entry, Print, Return, Int, etc.)
  - [x] Identifiers (variable names, function names)
  - [x] Literals (strings, integers, floats)
  - [x] Operators (+, -, *, /, =, etc.)
  - [x] Delimiters ({, }, (, ), ;, etc.)
  - [x] Comments (// single-line, /* */ multi-line)
- [x] Handle whitespace and newlines
- [ ] Error reporting for invalid tokens

### 1.2 Syntax Analysis (Parser)
- [x] Define grammar specification for Dread
- [x] Implement parser for:
  - [x] Function declarations (Entry functions)
  - [x] Variable declarations and assignments
  - [x] Function calls (Print, Return)
  - [ ] Expressions and operators
  - [x] Block statements
  - [x] Comments and docstrings
- [x] Build Abstract Syntax Tree (AST)
- [ ] Syntax error reporting with line numbers

### 1.3 Semantic Analysis
- [ ] Symbol table implementation
- [ ] Type checking system
- [ ] Scope management
- [ ] Function signature validation
- [ ] Variable initialization checking
- [ ] Semantic error reporting

## Phase 2: Type System

### 2.1 Basic Types
- [x] Integer types (Int)
- [x] String types with proper escaping
- [ ] Boolean types
- [ ] Float/Double types
- [ ] Character types

### 2.2 Type System Features
- [ ] Duck typing implementation (as shown in example)
- [ ] Type inference
- [ ] Type conversion rules
- [ ] Type compatibility checking

## Phase 3: Code Generation

### 3.1 Intermediate Representation (IR)
- [ ] Design IR format
- [ ] AST to IR translation
- [ ] IR optimization passes

### 3.2 Machine Code Generation
- [x] Target architecture selection (x86-64, ARM, etc.)
- [x] Assembly code generation
- [x] Object file creation
- [x] Linking with system libraries

### 3.3 Built-in Functions
- [x] Print function implementation
- [ ] Standard I/O operations
- [ ] Memory management functions
- [ ] System call interfaces

## Phase 4: Standard Library

### 4.1 Core Library
- [ ] String manipulation functions
- [ ] Mathematical operations
- [ ] Collection types (arrays, lists)
- [ ] File I/O operations

### 4.2 Advanced Features
- [ ] Memory management (garbage collection or manual)
- [ ] Concurrency primitives
- [ ] Networking capabilities
- [ ] Regular expressions

## Phase 5: Development Tools

### 5.1 Compiler Infrastructure
- [x] Command-line interface for compiler
- [ ] Build system integration
- [ ] Debugging information generation
- [ ] Optimization levels

### 5.2 Developer Experience
- [ ] Error message improvement
- [ ] Warning system
- [ ] Language server protocol (LSP) for IDE support
- [ ] Syntax highlighting definitions
- [ ] Documentation generator

## Phase 6: Language Features

### 6.1 Control Flow
- [ ] If/else statements
- [ ] Loop constructs (for, while)
- [ ] Switch/case statements
- [ ] Break and continue statements

### 6.2 Functions and Procedures
- [ ] Function parameters and arguments
- [ ] Return value handling
- [ ] Function overloading
- [ ] Anonymous functions/lambdas
- [ ] Closures

### 6.3 Data Structures
- [ ] Arrays and slices
- [ ] Structures/records
- [ ] Enumerations
- [ ] Unions
- [ ] Maps/dictionaries

### 6.4 Object-Oriented Features (Optional)
- [ ] Classes and objects
- [ ] Inheritance
- [ ] Polymorphism
- [ ] Interfaces

## Phase 7: Advanced Features

### 7.1 Memory Management
- [ ] Stack allocation
- [ ] Heap allocation
- [ ] Garbage collection (if applicable)
- [ ] Memory safety features

### 7.2 Concurrency
- [ ] Threading support
- [ ] Async/await patterns
- [ ] Channel communication
- [ ] Synchronization primitives

### 7.3 Metaprogramming
- [ ] Macros
- [ ] Compile-time evaluation
- [ ] Reflection capabilities
- [ ] Code generation

## Phase 8: Testing and Quality Assurance

### 8.1 Test Suite
- [ ] Unit tests for compiler components
- [ ] Integration tests for language features
- [ ] Performance benchmarks
- [ ] Regression test suite

### 8.2 Documentation
- [ ] Language specification
- [ ] User manual
- [ ] API documentation
- [ ] Tutorial and examples
- [ ] Best practices guide

## Phase 9: Package Management

### 9.1 Module System
- [ ] Import/export mechanisms
- [ ] Module resolution
- [ ] Dependency management
- [ ] Version compatibility

### 9.2 Package Manager
- [ ] Package registry
- [ ] Dependency resolution
- [ ] Build automation
- [ ] Distribution tools

## Implementation Notes

### Current Status (Based on README)
- ✅ Basic syntax design (Entry functions, Print, Return)
- ✅ Comment syntax (// and /* */)
- ✅ String literals with single quotes
- ✅ Duck typing concept
- ✅ Uppercase keyword convention

### MVP Accomplished (August 2025)
- ✅ **Working lexer** - Tokenizes all basic language constructs
- ✅ **Functional parser** - Builds AST for hello world program
- ✅ **Code generator** - Outputs x86-64 assembly
- ✅ **Complete compiler** - Produces working ELF executables
- ✅ **Hello World** - Successfully compiles and runs example program
- ✅ **System integration** - Uses system assembler and linker
- ✅ **Project structure** - Organized Go modules and packages

### Key Design Decisions to Make
- [ ] Memory management strategy (GC vs manual)
- [ ] Compilation target (native, VM, transpilation)
- [ ] Standard library scope and design
- [ ] Error handling mechanisms
- [ ] Concurrency model
- [ ] Package/module system design

### Development Environment Setup
- [x] Set up Go development environment
- [x] Create project structure
- [ ] Set up testing framework
- [ ] Configure CI/CD pipeline
- [ ] Set up documentation system

## Getting Started
1. ✅ Begin with Phase 1.1 (Lexical Analysis)
2. ✅ Create a simple tokenizer that can handle the hello world example
3. 🔄 Gradually expand to handle more language constructs
4. 🔄 Test each component thoroughly before moving to the next phase

## Next Priority Items
Based on the MVP completion, the immediate next steps should be:

1. **Phase 1.3 - Semantic Analysis** (Foundation)
   - Symbol table implementation
   - Basic type checking
   - Scope management

2. **Phase 6.1 - Control Flow** (Language Expansion)
   - If/else statements
   - Loop constructs (for, while)

3. **Phase 2.2 - Type System** (Robustness)
   - Type inference
   - Type conversion rules

4. **Error Handling & Testing** (Quality)
   - Syntax error reporting with line numbers
   - Unit tests for compiler components

---

*This TODO is a living document and should be updated as the language evolves and new requirements emerge.*
