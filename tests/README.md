# Test Suite

This directory contains test files for the Dread programming language compiler.

## Test Files

### Passing Tests (7/8)
- `entry.dread` - Basic Entry function
- `function.dread` - Function definitions and calls with string parameters
- `hello.dread` - Hello World program
- `integer.dread` - Basic integer printing
- `test_basic_integers.dread` - Multiple integer printing
- `test_int_vars.dread` - Integer variable assignment and printing
- `test_integers.dread` - Complex integer and string printing

### Known Limitations
- `test_int_functions.dread` - Function parameters with integers (segfaults due to limited parameter handling)

## Running Tests

Use the test runner to execute all tests:
```bash
go run cmd/test/main.go
```

## Adding New Tests

1. Create a new `.dread` file in this directory
2. The test runner will automatically pick it up
3. Tests should compile successfully and run without errors
4. Output is captured and shown for passing tests

## Test Organization

- Tests from `examples/valid/` are automatically included
- Manual test files use the `test_*.dread` naming convention
- Empty or broken test files should be removed to keep the suite clean
