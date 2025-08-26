# Dread Language Development Milestones

This document tracks completed milestones and defines upcoming development targets for the Dread programming language.

## 🎯 Current Status: **MVP Complete**

## ✅ **Milestone 1: Minimum Viable Product (MVP)** - *Completed August 2025*

**Goal**: Create a working compiler that can compile and run a basic "Hello, World!" program.

### Completed Features:
- [x] **Lexical Analysis**: Complete tokenization of basic Dread syntax
- [x] **Parser**: AST generation for Entry functions, assignments, and calls
- [x] **Code Generation**: x86-64 assembly output with Linux system calls
- [x] **Entry Functions**: Support for `Entry main() (Int)` program entry points
- [x] **Variables**: Duck-typed variable assignment
- [x] **Built-ins**: `Print()` and `Return()` functions
- [x] **Comments**: Single-line (`//`) and multi-line (`/* */`) support
- [x] **String Literals**: Single-quoted strings
- [x] **Compiler Driver**: Complete compilation pipeline with assembler/linker integration
- [x] **Documentation**: Comprehensive docs (README, ARCHITECTURE, SPECIFICATION, etc.)

### Key Constraint:
- **One Entry per executable**: Each program must have exactly one `Entry main()` function

---

## 🚧 **Milestone 2: Function Support** - *Next Target*

**Goal**: Add support for regular `Function` declarations alongside `Entry` functions.

### Phase 2.1: Function Declaration Support
- [ ] **Lexer**: Add `FUNCTION` token type
- [ ] **Keywords**: Add `"Function"` to keywords map
- [ ] **Parser**: Handle both `ENTRY` and `FUNCTION` in parseStatement()
- [ ] **AST**: Extend FunctionStatement to distinguish Entry vs Function
- [ ] **Code Generation**: Generate assembly for multiple functions
- [ ] **Validation**: Ensure exactly one Entry function per program

### Phase 2.2: Function Calling Mechanism
- [ ] **Function Calls**: Implement calling regular functions (not just built-ins)
- [ ] **Call Stack**: Proper stack frame management
- [ ] **Return Values**: Handle function return values
- [ ] **Function Parameters**: Add support for function parameters

### Example Target Syntax:
```dread
Function greet() (Int)
{
    Print('Hello from greet!')
    Return(0)
}

Entry main() (Int)
{
    greet()  // Call the function
    Return(0)
}
```

### Constraints:
- **One Entry per executable**: Still exactly one `Entry` function required
- **Functions cannot call Entry**: Entry functions are program entry points only
- **No recursion initially**: Keep function calls simple

---

## 🔮 **Milestone 3: Expressions and Arithmetic** - *Future*

**Goal**: Add mathematical expressions and operators.

### Planned Features:
- [ ] **Arithmetic Operators**: `+`, `-`, `*`, `/`
- [ ] **Comparison Operators**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- [ ] **Operator Precedence**: Proper expression parsing
- [ ] **Integer Variables**: Full integer arithmetic support
- [ ] **Parentheses**: Expression grouping

### Example Target Syntax:
```dread
Entry main() (Int)
{
    result = 5 + 3 * 2  // Should be 11
    Print(result)
    Return(0)
}
```

---

## 🔮 **Milestone 4: Control Flow** - *Future*

**Goal**: Add conditional statements and loops.

### Planned Features:
- [ ] **If/Else Statements**: Conditional execution
- [ ] **Boolean Type**: `True`/`False` literals
- [ ] **Boolean Operators**: `And`, `Or`, `Not`
- [ ] **While Loops**: Basic iteration
- [ ] **For Loops**: Counter-based iteration

### Example Target Syntax:
```dread
Entry main() (Int)
{
    x = 5
    If (x > 3)
    {
        Print('x is greater than 3')
    }
    Else
    {
        Print('x is not greater than 3')
    }
    Return(0)
}
```

---

## 🔮 **Milestone 5: Advanced Types** - *Future*

**Goal**: Expand the type system beyond strings and integers.

### Planned Features:
- [ ] **Float Type**: Floating-point numbers
- [ ] **Boolean Type**: Proper boolean values
- [ ] **Arrays**: Basic array support
- [ ] **Type Checking**: Static type validation
- [ ] **Type Conversion**: Explicit and implicit conversions

---

## 🔮 **Milestone 6: Standard Library** - *Future*

**Goal**: Build a comprehensive standard library.

### Planned Features:
- [ ] **String Functions**: Length, substring, concatenation
- [ ] **Math Functions**: Sin, cos, sqrt, etc.
- [ ] **File I/O**: Read/write files
- [ ] **Input Functions**: Read from stdin

---

## 📊 **Development Metrics**

### Current Codebase Size:
- **Lines of Go Code**: ~800 lines
- **Supported Language Features**: 8 core features
- **Test Coverage**: Manual testing (automated tests planned)

### Next Milestone Progress:
- **Function Support**: 0% complete
- **Estimated Effort**: 2-3 development sessions
- **Risk Level**: Low (extends existing patterns)

---

## 🎯 **Development Principles**

1. **Incremental Progress**: Each milestone builds on the previous
2. **Documentation First**: Update specs before implementing
3. **Test-Driven**: Create test cases for new features
4. **Backward Compatibility**: Don't break existing programs
5. **Simplicity**: Keep the language easy to understand

---

## 📝 **Notes**

- **Current Priority**: Focus on Milestone 2 (Function Support)
- **Testing Strategy**: Manual testing with example programs
- **Documentation**: Update all docs when milestones complete
- **Architecture**: Maintain clean separation between lexer/parser/codegen

---

*This milestone document should be updated as we complete each phase and plan new features.*
