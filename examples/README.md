# Dread Programming Examples

This directory contains example programs demonstrating the current capabilities of the Dread programming language.

## Available Examples

### hello.dread
**Description**: Complete hello world program with extensive comments
**Features**: Comments, variables, Print function, Return statement
**Compilation**: `./dreadc examples/hello.dread hello`

```dread
// This is a comment!

/*
 * This is a docstring for the entrypoint (Entry), which is the start of the executable program.
 * It has no arguments, for now. It returns integer by default
 */
Entry main() (Int)
{
    /* This is also a comment! Underneath is a duck typed declaration and initialization of a variable of string type*/
    hello_string = 'Hello, World!\n'

    Print(hello_string) // This is a built in function in dread for printing to stdout
    // All keywords in dread start with an uppercase first letter! Like Entry, Print, and also:
    Return(0) // Return! Which returns zero, like most programs do
}
```

### hello_simple.dread
**Description**: Minimal hello world program
**Features**: Basic syntax without extensive comments
**Compilation**: `./dreadc examples/hello_simple.dread hello_simple`

```dread
Entry main() (Int)
{
    hello_string = 'Hello, World!
'
    Print(hello_string)
    Return(0)
}
```

### entry_demo.dread
**Description**: Demonstrates Entry function constraints and current language features
**Features**: Entry function rules, comments explaining future Function support
**Compilation**: `./dreadc examples/entry_demo.dread entry_demo`

```dread
// Example demonstrating current Entry function constraints
Entry main() (Int)
{
    program_name = 'Dread Language Demo
'
    Print(program_name)
    Return(0)
}

// Future: Regular functions will be declared with 'Function' keyword
// Function helper() (Int) { ... }
```

## Running Examples

1. **Build the compiler** (if not already built):
   ```bash
   go build -o dreadc ./cmd/dreadc
   ```

2. **Compile an example**:
   ```bash
   ./dreadc examples/hello.dread my_program
   ```

3. **Run the compiled program**:
   ```bash
   ./my_program
   ```

4. **Clean up**:
   ```bash
   rm my_program
   ```

## Example Patterns

### Basic Program Structure

Every Dread program follows this pattern:

```dread
Entry main() (Int)
{
    // Your code here
    Return(0)
}
```

### Variable Assignment

```dread
Entry main() (Int)
{
    // String variable
    message = 'Hello, Dread!'

    // Number variable (used with Return)
    exit_code = 0

    Print(message)
    Return(exit_code)
}
```

### Multiple Operations

```dread
Entry main() (Int)
{
    // Multiple print statements
    greeting = 'Hello'
    name = 'World'

    Print(greeting)
    Print(name)

    Return(0)
}
```

### Comments

```dread
// Single-line comment explaining the next line
Entry main() (Int)
{
    /*
     * Multi-line comment
     * explaining a complex section
     */
    message = 'Commented code'
    Print(message) // Inline comment
    Return(0)
}
```

## Creating Your Own Examples

1. **Create a new .dread file**:
   ```bash
   touch examples/my_example.dread
   ```

2. **Follow the basic structure**:
   ```dread
   Entry main() (Int)
   {
       // Your code here
       Return(0)
   }
   ```

3. **Use available features**:
   - String variables with single quotes
   - Print function for output
   - Return function for exit codes
   - Comments for documentation

4. **Test your program**:
   ```bash
   ./dreadc examples/my_example.dread test
   ./test
   rm test
   ```

## Current Limitations

These examples demonstrate the current MVP capabilities. The following features are **not yet available**:

- ❌ Arithmetic operations (`+`, `-`, `*`, `/`)
- ❌ Conditional statements (`if`, `else`)
- ❌ Loops (`while`, `for`)
- ❌ Function parameters
- ❌ Multiple functions
- ❌ Boolean values and logic
- ❌ Arrays or complex data structures

See `TODO.md` for the development roadmap of upcoming features.

## Example Output

Running the hello world example:

```bash
$ ./dreadc examples/hello.dread hello
Successfully compiled examples/hello.dread to hello

$ ./hello
Hello, World!

$ echo $?
0
```

## Tips for Writing Dread Programs

1. **Always include Return()**: Programs should end with `Return(0)` for successful execution

2. **Use meaningful variable names**: Variables help document your code
   ```dread
   welcome_message = 'Welcome!'  // Good
   x = 'Welcome!'                // Less clear
   ```

3. **Add comments**: Explain what your program does
   ```dread
   // Calculate and display greeting
   greeting = 'Hello, World!'
   Print(greeting)
   ```

4. **Test incrementally**: Compile and test after each change

5. **Check the specification**: See `SPECIFICATION.md` for detailed language rules

## Troubleshooting

### Common Issues

1. **"Parse error" during compilation**:
   - Check that all strings use single quotes `'`
   - Ensure all statements end properly
   - Verify function declaration syntax

2. **"Assembly/linking failed"**:
   - Make sure `as` and `ld` are installed
   - Check file permissions
   - Try recompiling

3. **Program produces no output**:
   - Ensure you're calling `Print()`
   - Check that strings contain visible characters
   - Verify the program isn't exiting early

4. **"Successfully compiled" but executable doesn't work**:
   - Check file permissions: `chmod +x program_name`
   - Try running with `./program_name`
   - Verify the program calls `Return()`

### Getting Help

- Check `SPECIFICATION.md` for language syntax
- Review `ARCHITECTURE.md` for compiler details
- Look at `DEVELOPER.md` for debugging tips
- Examine working examples for patterns

---

Happy coding with Dread! 🚀
