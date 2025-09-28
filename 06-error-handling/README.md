# 06 - Error Handling

This example demonstrates Go's error handling patterns, which are fundamental to writing robust Go programs. Go takes a different approach to error handling compared to many other languages.

## Key Concepts

### Basic Error Handling

Go uses explicit error returns rather than exceptions:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 0)
if err != nil {
    // handle error
}
```

### Custom Error Types

Create custom error types for more context:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}
```

### Error Wrapping

Go 1.13+ supports error wrapping for better error context:

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    // ...
}
```

### Unwrapping Errors

Access the underlying error:

```go
if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
    // handle wrapped error
}
```

### Sentinel Errors

Predefined error values for common cases:

```go
var (
    ErrInvalidInput = errors.New("invalid input")
    ErrNotFound     = errors.New("not found")
)

// Check with errors.Is()
if errors.Is(err, ErrNotFound) {
    // handle not found case
}
```

### Error Type Checking

#### Type Assertion (old way)
```go
if valErr, ok := err.(ValidationError); ok {
    // handle ValidationError
}
```

#### errors.As() (new way)
```go
var valErr ValidationError
if errors.As(err, &valErr) {
    // handle ValidationError
}
```

### Panic and Recover

Use sparingly - only for truly exceptional circumstances:

```go
func riskyOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("operation panicked: %v", r)
        }
    }()
    // ... code that might panic
}
```

## Error Handling Patterns

### 1. Early Return Pattern
```go
func goodFunction(data string) (result string, err error) {
    if data == "" {
        return "", errors.New("data cannot be empty")
    }

    // process data
    result = fmt.Sprintf("processed: %s", data)
    return result, nil
}
```

### 2. Error Wrapping for Context
```go
func readConfig(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("failed to read config file %s: %w", filename, err)
    }
    // ...
}
```

### 3. Multiple Error Collection
```go
func processItems(items []string) error {
    var errs []error
    for i, item := range items {
        if err := validateItem(item); err != nil {
            errs = append(errs, fmt.Errorf("item %d: %w", i, err))
        }
    }

    if len(errs) > 0 {
        return fmt.Errorf("multiple validation errors: %v", errs)
    }
    return nil
}
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
=== Basic Error Handling ===
10 / 2 = 5.00
Error: cannot divide by zero

=== Custom Error Types ===
Validation error: validation error on field 'name': cannot be empty
Field: name, Message: cannot be empty
Validation error: validation error on field 'age': cannot be negative

=== Error Wrapping ===
Process file error: failed to process file nonexistent.txt: failed to open file nonexistent.txt: open nonexistent.txt: no such file or directory
Original error: failed to process file nonexistent.txt: failed to open file nonexistent.txt: open nonexistent.txt: no such file or directory
Unwrapped error: failed to open file nonexistent.txt: open nonexistent.txt: no such file or directory

=== Network Error Example ===
Network error: network GET http://example.com: timeout

=== Sentinel Errors ===
findUser() error: invalid input
findUser(notfound) error: not found
  -> User not found
findUser(private) error: unauthorized
  -> Access denied
Found user: validuser

=== Panic and Recover (Not Recommended) ===
Recovered from panic: operation panicked: something went wrong

=== Multiple Error Handling ===
Multiple errors: multiple validation errors: [item 0: item too short item 2: item too short]

=== Error Handling Best Practices ===
Good practice error: data cannot be empty
Good practice result: processed: valid data
Bad practice error: data cannot be empty

=== Error Type Checking ===
Type assertion: Field=email, Message=invalid format
errors.Is: This is a NotFound error
errors.As: Field=email, Message=invalid format
```

## Best Practices

### ✅ Do's
- **Return errors explicitly** - Don't use panics for normal errors
- **Check errors immediately** - Handle errors as soon as they occur
- **Use error wrapping** - Add context with `fmt.Errorf` and `%w`
- **Use sentinel errors** - For common, expected error conditions
- **Use custom error types** - When you need additional context
- **Use `errors.Is()` and `errors.As()`** - For checking error types

### ❌ Don'ts
- **Don't ignore errors** - Always check and handle errors
- **Don't use panic/recover** for normal flow control
- **Don't create unnecessary custom error types** - Use wrapping instead
- **Don't hide errors** - Preserve error context when wrapping
- **Don't use `err == nil` for custom errors** - Use `errors.Is()`

### Error Handling Guidelines

1. **Fail Fast** - Return errors as soon as possible
2. **Single Responsibility** - Functions should do one thing and report errors clearly
3. **Consistent Patterns** - Use the same error handling patterns throughout your codebase
4. **Meaningful Messages** - Error messages should help with debugging
5. **Context Preservation** - Wrap errors to maintain the error chain

## Common Interview Questions

- How does Go's error handling differ from exceptions?
- When should you use panic vs return an error?
- What's the difference between `errors.Is()` and `errors.As()`?
- How do you handle multiple errors from different operations?

## Next Steps

- Practice wrapping errors in your own functions
- Create custom error types for domain-specific errors
- Implement proper error handling in existing code
- Study error handling patterns in popular Go libraries