# Dread Language Development Milestones

This document tracks completed milestones and defines upcoming development targets for the Dread programming language.

## 🎯 Current Status: **Function Support Complete**

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

## ✅ **Milestone 2: Function Support** - *Completed August 2025* 🎉

**Goal**: Add support for regular `Function` declarations alongside `Entry` functions.

### Completed Features:

#### Phase 2.1: Function Declaration Support [✅ COMPLETE]
- ✅ **Lexer**: Add `FUNCTION` token type
- ✅ **Keywords**: Add `"Function"` to keywords map
- ✅ **Lexer**: Add `STRING_TYPE`, `VOID_TYPE` tokens
- ✅ **Lexer**: Add `COMMA` token for parameter lists
- ✅ **Parser**: Handle both `ENTRY` and `FUNCTION` in parseStatement()
- ✅ **Parser**: Support multiple return type syntaxes (`() Type`, `() (Type)`, `() {}`)
- ✅ **Parser**: Parameter parsing for `Type name` syntax
- ✅ **AST**: Extended FunctionStatement with IsEntry, Parameters, CallExpression
- ✅ **Code Generation**: Generate assembly for multiple functions
- ✅ **Validation**: Ensure exactly one Entry function per program

#### Phase 2.2: Function Calling Mechanism [✅ COMPLETE]
- ✅ **Function Calls**: Implement calling regular functions (no parameters)
- ✅ **Function Calls**: Implement calling functions WITH parameters
- ✅ **Call Stack**: Proper stack frame management
- ✅ **Entry vs Function Returns**: Entry functions exit program, regular functions return to caller
- ✅ **Function Call Expressions**: Support `var = function()` syntax
- ✅ **Return Values**: Functions return values and are captured in variables
- ✅ **Function Parameters**: Parameter passing implemented via x86-64 calling convention
- ✅ **Parameter Printing**: Functions can print parameters correctly
- ✅ **Return Value Printing**: Return values can be printed correctly

### 🏆 **Complete Implementation Status**
- **✅ Working**: Multiple function declarations (Entry + Function)
- **✅ Working**: Function calls without parameters
- **✅ Working**: Function calls with parameters
- **✅ Working**: Proper Entry vs Function distinction
- **✅ Working**: Function call assignments with return value capture
- **✅ Working**: Complex syntax variations (`Function name() Type`, etc.)
- **✅ Working**: Return value capture and storage in variables
- **✅ Working**: Parameter passing via x86-64 calling convention (rdi + rsi)
- **✅ Working**: Return value passing via x86-64 registers (rax + r8)

### Successfully Compiled & Running Examples:
```dread
// Function with no arguments that returns nothing
Function fun_noarg_noret() {
    Print('No args! No rets!\n')
}

// Function with no arguments that returns a String
Function fun_noarg_ret() String {
    Return('No args! Rets!\n')
}

// Function with one String argument that returns nothing
Function fun_arg_noret(String input_str) Void {
    Print(input_str)
}

// Function with one String argument that returns a String
Function fun_arg_ret(String input_str) String {
    Print(input_str)
    Return('Args! Rets!\n')
}

Entry main() {
    fun_noarg_noret()

    first_return_value = fun_noarg_ret()
    Print(first_return_value)

    fun_arg_noret('Args! No rets!\n')

    second_return_value = fun_arg_ret('Args! Rets! Input!\n')
    Print(second_return_value)
}
```

**Output**:
```
No args! No rets!
No args! Rets!
Args! No rets!
Args! Rets! Input!
Args! Rets!
```

### Technical Implementation Details:
- **Parameter Passing**: Uses x86-64 calling convention with `rdi` (address) and `rsi` (length) for first parameter
- **Return Values**: Uses `rax` (string address) and `r8` (string length) for return values
- **Stack Management**: Proper function prologue/epilogue with `rbp` and `rsp`
- **String Handling**: Correct length calculation and memory boundary management
- **Assembly Generation**: Clean separation between Entry and Function code paths

### Constraints:
- **One Entry per executable**: Still exactly one `Entry` function required
- **Single Parameter Support**: Currently supports one string parameter per function
- **String Return Types**: Functions can return String or Void
- **No Recursion**: Functions cannot call themselves

**🚀 Achievement Unlocked**: The Dread language now has complete function support with parameters and return values!
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
- **Lines of Go Code**: ~1200 lines (expanded significantly for function support)
- **Supported Language Features**: 12 core features (MVP + Functions)
- **Test Coverage**: Manual testing with comprehensive examples
- **Assembly Generation**: Full x86-64 with Linux system calls

### Completed Milestones:
- **Milestone 1 (MVP)**: 100% complete ✅
- **Milestone 2 (Functions)**: 100% complete ✅

### Next Milestone Progress:
- **Expressions and Arithmetic**: 0% complete
- **Estimated Effort**: 3-4 development sessions
- **Risk Level**: Medium (new parsing patterns required)

---

## 🎯 **Development Principles**

1. **Incremental Progress**: Each milestone builds on the previous
2. **Documentation First**: Update specs before implementing
3. **Test-Driven**: Create test cases for new features
4. **Backward Compatibility**: Don't break existing programs
5. **Simplicity**: Keep the language easy to understand

---

## 📝 **Notes**

- **Current Priority**: Begin Milestone 3 (Expressions and Arithmetic)
- **Recent Achievement**: Complete function support with parameters and return values ✅
- **Testing Strategy**: Manual testing with comprehensive function examples
- **Documentation**: All milestone documentation updated August 2025
- **Architecture**: Clean separation maintained across lexer/parser/codegen
- **Function Architecture**: Robust x86-64 calling convention implementation

---

*Last Updated: August 27, 2025 - Milestone 2 Complete*
