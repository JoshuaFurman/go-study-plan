# 12. Unit Testing

This example demonstrates Go's built-in testing framework, including basic unit tests, table-driven tests, subtests, and benchmarks. Testing is crucial for maintaining code quality and preventing regressions.

## Key Concepts

### Basic Unit Tests
- Functions named `TestXxx` with `*testing.T` parameter
- Use `t.Errorf()` or `t.Fatalf()` for failures
- Tests run in isolation

### Table-Driven Tests
- Define test cases as slices of structs
- Iterate through test cases in a loop
- Each test case has input and expected output
- Makes tests more maintainable and comprehensive

### Subtests
- Use `t.Run()` to create nested test groups
- Organize related test cases hierarchically
- Run specific subtests with `go test -run TestName/SubtestName`

### Benchmarks
- Functions named `BenchmarkXxx` with `*testing.B`
- Use `b.N` to control iterations
- Measure performance of functions

## Test Structure

### Basic Test Function
```go
func TestAdd(t *testing.T) {
    calc := Calculator{}
    result := calc.Add(2, 3)
    expected := 5.0

    if result != expected {
        t.Errorf("Add(2, 3) = %f; want %f", result, expected)
    }
}
```

### Table-Driven Test
```go
func TestIsEven(t *testing.T) {
    testCases := []struct {
        input    int
        expected bool
    }{
        {2, true},
        {3, false},
        {0, true},
    }

    for _, tc := range testCases {
        result := calc.IsEven(tc.input)
        if result != tc.expected {
            t.Errorf("IsEven(%d) = %t; want %t", tc.input, result, tc.expected)
        }
    }
}
```

### Subtests
```go
func TestArithmeticOperations(t *testing.T) {
    t.Run("Addition", func(t *testing.T) {
        // Addition test cases
    })

    t.Run("Subtraction", func(t *testing.T) {
        // Subtraction test cases
    })
}
```

### Benchmark
```go
func BenchmarkAdd(b *testing.B) {
    calc := Calculator{}
    for i := 0; i < b.N; i++ {
        calc.Add(2, 3)
    }
}
```

## Running Tests

### Run all tests
```bash
go test
```

### Run specific test
```bash
go test -run TestAdd
```

### Run subtests
```bash
go test -run TestArithmeticOperations/Addition
```

### Run benchmarks
```bash
go test -bench=.
```

### Run with verbose output
```bash
go test -v
```

### Run with coverage
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Organization

### File Naming
- Test files end with `_test.go`
- Same package as code being tested (package main)
- Can create separate test packages for black-box testing

### Test Helpers
- Helper functions should take `*testing.T` as first parameter
- Use `t.Helper()` to mark helper functions
- Failures in helpers are attributed to calling test

### Parallel Tests
- Use `t.Parallel()` to run tests concurrently
- Useful for integration tests
- Must be careful with shared resources

## Best Practices

1. **Write tests first** (TDD - Test Driven Development)
2. **Test edge cases** and error conditions
3. **Use descriptive test names** that explain what they're testing
4. **Keep tests fast** and independent
5. **Use table-driven tests** for similar test cases
6. **Test error conditions** explicitly
7. **Use subtests** to organize complex test suites
8. **Measure performance** with benchmarks
9. **Aim for high coverage** but focus on meaningful tests

## Common Testing Patterns

### Testing Errors
```go
result, err := calc.Divide(10, 0)
if err == nil {
    t.Error("Expected error for division by zero")
}
```

### Testing Multiple Return Values
```go
fact, err := calc.Factorial(-1)
if err == nil {
    t.Error("Expected error for negative factorial")
}
```

### Setup/Teardown
```go
func TestMain(m *testing.M) {
    // Setup code
    code := m.Run()
    // Teardown code
    os.Exit(code)
}
```

## Running the Example

```bash
cd 12-unit-testing

# Run the demo
go run main.go

# Run all tests
go test

# Run tests with verbose output
go test -v

# Run benchmarks
go test -bench=.

# Run with coverage
go test -cover
```

This example provides a comprehensive foundation for testing Go code, essential for professional development and interview preparation.