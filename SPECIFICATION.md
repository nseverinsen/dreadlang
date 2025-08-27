# Dread Language Specification

**Version**: 0.1.0 (MVP)
**Status**: Draft
**Last Updated**: August 2025

## Overview

Dread is a general-purpose programming language designed for simplicity and readability. The language features uppercase keywords, duck typing, and direct system integration.

## Lexical Structure

### Character Set

Dread source files are encoded in UTF-8. The language uses ASCII characters for syntax elements.

### Comments

Dread supports two types of comments:

1. **Single-line comments**: Start with `//` and continue to the end of the line
   ```dread
   // This is a single-line comment
   ```

2. **Multi-line comments**: Enclosed in `/*` and `*/`, can span multiple lines
   ```dread
   /*
    * This is a multi-line comment
    * that spans multiple lines
    */
   ```

### Identifiers

Identifiers name variables, functions, and other user-defined entities.

**Syntax**:
- Must start with a letter (`a-z`, `A-Z`) or underscore (`_`)
- Followed by letters, digits (`0-9`), or underscores
- Case-sensitive

**Examples**:
```dread
main
hello_world
variable1
_private
```

### Keywords

All keywords in Dread start with an uppercase letter:

| Keyword    | Purpose                         |
|------------|---------------------------------|
| `Entry`    | Entry point function declaration|
| `Function` | Regular function declaration    |
| `Print`    | Built-in print function         |
| `Return`   | Return statement                |
| `Int`      | Integer type annotation         |

**Reserved for future use**:
`If`, `Else`, `While`, `For`, `True`, `False`, `String`, `Bool`, `Float`, `Function`

### Literals

#### String Literals

String literals are enclosed in single quotes (`'`):

```dread
'Hello, World!'
'This is a string'
'String with\nnewline'  // Note: \n is literal, not escape sequence in current implementation
```

**Current limitations**:
- Only single quotes supported
- Limited escape sequence processing
- Newlines must be literal characters in the string

#### Integer Literals

Integer literals are sequences of digits:

```dread
0
42
123
```

**Current limitations**:
- Only decimal integers supported
- No negative number syntax yet
- No floating-point literals

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `=`      | Assignment  | `x = 5` |

**Future operators**: `+`, `-`, `*`, `/`, `==`, `!=`, `<`, `>`, etc.

### Delimiters

| Symbol | Purpose           |
|--------|-------------------|
| `(`    | Left parenthesis  |
| `)`    | Right parenthesis |
| `{`    | Left brace        |
| `}`    | Right brace       |

## Syntax

### Program Structure

A Dread program consists of function definitions with specific structural requirements.

#### Program Entry Point
Every Dread program **must** have exactly one `Entry` function named `main`:

```dread
Entry main() (Int)
{
    // Program body
    Return(0)
}
```

#### Program Structure Rules
1. **Exactly one Entry function**: Each executable must contain one `Entry main()` function
2. **Entry function naming**: The Entry function must be named `main`
3. **Entry function signature**: Must be `Entry main() (Int)`
4. **Multiple regular functions**: Programs may contain multiple `Function` declarations (future feature)
5. **Entry cannot be called**: Entry functions are entry points, not callable functions

#### Invalid Programs
```dread
// ❌ ERROR: Multiple Entry functions
Entry main() (Int) { Return(0) }
Entry start() (Int) { Return(1) }
```

### Functions

#### Entry Point Function Declaration

**Syntax** (Currently supported):
```
Entry <function_name>() (<return_type>)
{
    <statements>
}
```

#### Regular Function Declaration

**Syntax** (Future feature):
```
Function <function_name>() (<return_type>)
{
    <statements>
}
```

**Current limitations**:
- Only `Entry` functions supported (regular `Function` declarations planned)
- No function parameters
- Only `Int` return type
- Only one function per program

**Entry Function Constraints**:
- **Exactly one Entry per executable**: Each program must have one and only one `Entry` function
- **Entry function must be named `main`**: The entry point must be `Entry main()`
- **Entry functions cannot be called**: Entry functions are program entry points, not callable functions
- **Entry functions must return Int**: Exit code for the operating system

**Example**:
```dread
Entry main() (Int)
{
    Return(0)
}
```

### Variables

#### Variable Declaration and Assignment

Variables are declared and initialized using assignment:

```dread
variable_name = value
```

**Type inference**: Variable types are inferred from their assigned values (duck typing).

**Examples**:
```dread
message = 'Hello, World!'  // String type inferred
number = 42                // Integer type inferred
```

### Statements

#### Assignment Statement

**Syntax**: `<identifier> = <expression>`

**Example**:
```dread
greeting = 'Hello!'
```

#### Function Call Statement

**Syntax**: `<function_name>(<arguments>)`

**Built-in functions**:

1. **Print**: Output to standard output
   ```dread
   Print(expression)
   ```

2. **Return**: Exit program with status code
   ```dread
   Return(exit_code)
   ```

**Examples**:
```dread
Print('Hello, World!')
Print(variable_name)
Return(0)
```

### Expressions

#### Primary Expressions

1. **String literals**: `'text'`
2. **Integer literals**: `123`
3. **Identifiers**: `variable_name`

**Current limitations**:
- No arithmetic expressions
- No boolean expressions
- No function call expressions

## Type System

### Types

#### Built-in Types

1. **Int**: Integer numbers
   - Used for return values and exit codes
   - Currently limited to literals

2. **String**: Text data
   - Single-quoted literals
   - Duck-typed variables

#### Type Inference

Variables are duck-typed - their type is inferred from the assigned value:

```dread
text = 'Hello'     // text is String type
code = 0           // code is Int type
```

#### Type Annotations

Function return types must be explicitly annotated:

```dread
Entry main() (Int)  // Must return Int type
{
    Return(0)       // Int literal
}
```

## Built-in Functions

### Print

**Purpose**: Output text to standard output

**Syntax**: `Print(expression)`

**Parameters**:
- `expression`: String or variable to print

**Example**:
```dread
Print('Hello, World!')
message = 'Goodbye!'
Print(message)
```

### Return

**Purpose**: Exit the program with a status code

**Syntax**: `Return(exit_code)`

**Parameters**:
- `exit_code`: Integer exit status (0 = success)

**Example**:
```dread
Return(0)  // Successful exit
Return(1)  // Error exit
```

## Program Execution

### Entry Point

Every Dread program must have an `Entry` function named `main`:

```dread
Entry main() (Int)
{
    // Program logic here
    Return(0)
}
```

### Execution Flow

1. Program starts at `Entry main()`
2. Statements execute in order
3. Program exits when `Return()` is called
4. Exit code becomes the program's exit status

## Memory Model

### Current Implementation

- **Variables**: Mapped to static string constants
- **Strings**: Stored in data section
- **Integers**: Compile-time constants only

### Future Plans

- Stack-allocated local variables
- Dynamic memory allocation
- Proper variable scoping

## Examples

### Hello World

```dread
// Basic hello world program
Entry main() (Int)
{
    Print('Hello, World!')
    Return(0)
}
```

### Variables

```dread
// Using variables
Entry main() (Int)
{
    greeting = 'Hello, Dread!'
    Print(greeting)
    Return(0)
}
```

### Comments

```dread
// This is a single-line comment

/*
 * This is a multi-line comment
 * explaining the program
 */
Entry main() (Int)
{
    /* Variable declaration */
    message = 'Hello!'

    Print(message)  // Print the message
    Return(0)       // Exit successfully
}
```

## Error Handling

### Compile-time Errors

- **Syntax errors**: Invalid token sequences
- **Parse errors**: Malformed program structure
- **Assembly errors**: Generated assembly issues

### Runtime Behavior

- Programs that don't call `Return()` may have undefined behavior
- Invalid system calls will cause program termination

## Limitations and Future Work

### Current Limitations

1. **Single file compilation**: No module system
2. **No arithmetic**: No mathematical expressions
3. **No control flow**: No if/else, loops, etc.
4. **Limited types**: Only String and Int
5. **No functions**: Only Entry points
6. **No parameters**: Functions take no arguments

### Planned Features

See `TODO.md` for detailed roadmap:

1. **Arithmetic expressions**: `+`, `-`, `*`, `/`
2. **Boolean logic**: `and`, `or`, `not`
3. **Control flow**: `If`, `Else`, `While`, `For`
4. **Functions**: Parameters, local variables, multiple functions
5. **Advanced types**: Arrays, structures, floats
6. **Module system**: Import/export, packages

## Grammar (BNF)

```bnf
<program>     ::= <entry_function>+

<entry_function> ::= "Entry" <identifier> "(" ")" "(" <type> ")" <block>

// Future: Regular functions will use "Function" keyword
// <function>    ::= "Function" <identifier> "(" ")" "(" <type> ")" <block>

<block>       ::= "{" <statement>* "}"

<statement>   ::= <assignment> | <call>

<assignment>  ::= <identifier> "=" <expression>

<call>        ::= <identifier> "(" <expression>? ")"

<expression>  ::= <string> | <integer> | <identifier>

<type>        ::= "Int"

<identifier>  ::= <letter> (<letter> | <digit> | "_")*

<string>      ::= "'" <character>* "'"

<integer>     ::= <digit>+

<letter>      ::= "a" ... "z" | "A" ... "Z"

<digit>       ::= "0" ... "9"

<character>   ::= any Unicode character except "'"
```

---

**Note**: This specification describes the current MVP implementation. Features marked as "future" or "planned" are not yet implemented. See the project's TODO.md for development roadmap.
