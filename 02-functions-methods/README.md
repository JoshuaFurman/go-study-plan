# 02 - Functions and Methods

This example demonstrates Go's function syntax, methods, interfaces, defer, and panic/recover mechanisms.

## Key Concepts

### Functions

#### Basic Functions
```go
func add(a, b int) int {
    return a + b
}
```
- Functions are declared with `func` keyword
- Parameters can share types: `a, b int`
- Return type follows parameters

#### Multiple Return Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```
Go commonly returns `(result, error)` pairs for error handling.

#### Named Return Values
```go
func rectangleProperties(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}
```
- Return values can be named
- `return` without values uses named returns
- Improves readability for multiple returns

#### Variadic Functions
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```
- `...int` accepts variable number of arguments
- Treated as slice inside function

#### Recursive Functions
```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```
Go supports recursion, but be mindful of stack overflow.

### Methods

#### Value vs Pointer Receivers
```go
// Value receiver - works on copy
func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Pointer receiver - modifies original
func (p *Person) Birthday() {
    p.Age++
}
```
- Value receivers work on copies (can't modify original)
- Pointer receivers work on original (can modify)
- Use pointer receivers when method modifies the receiver

### Interfaces

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```
- Interfaces define method signatures
- Types implement interfaces by implementing their methods
- No explicit "implements" keyword needed

### Defer

```go
func demonstrateDefer() {
    fmt.Println("Start")
    defer fmt.Println("Deferred 1") // Executes last
    defer fmt.Println("Deferred 2") // Executes second to last
    fmt.Println("End")
}
```
- `defer` postpones function execution until surrounding function returns
- Deferred calls are executed in LIFO (Last In, First Out) order
- Commonly used for cleanup operations

### Panic and Recover

```go
func safeDivision(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }

    result = a / b
    return
}
```
- `panic` stops normal execution and begins panicking
- `recover` regains control of a panicking goroutine
- Only useful inside deferred functions
- Use sparingly; prefer error returns

### Anonymous Functions

```go
addResult := applyOperation(3, 4, func(a, b int) int { return a + b })
```
- Functions without names
- Can be assigned to variables or passed as arguments
- Useful for callbacks and closures

## Running the Example

```bash
go run main.go
```

## Expected Output

```
5 + 3 = 8
10 / 2 = 5.00
Error: cannot divide by zero
Rectangle 4x5: Area=20.0, Perimeter=18.0
Sum of 1+2+3+4+5 = 15
5! = 120
Person{Name: Alice, Age: 25}
After birthday: Person{Name: Alice, Age: 26}
Rectangle area: 50.0
After scaling by 2: Width=20.0, Height=10.0, Area=200.0
Circle area: 28.27
Start
End
Deferred 2
Deferred 1
10 / 2 = 5
Recovered from panic: panic recovered: division by zero
3 + 4 = 7 (using anonymous function)
3 * 4 = 12 (using anonymous function)
```

## Common Patterns

1. **Error Handling**: Return `(result, error)` pairs
2. **Resource Cleanup**: Use `defer` for file closes, mutex unlocks
3. **Interface Design**: Keep interfaces small (Single Responsibility)
4. **Method Receivers**: Use pointer receivers when modifying state

## Best Practices

- Use multiple return values for error handling
- Prefer error returns over panic for expected errors
- Use defer for cleanup operations
- Keep interfaces small and focused
- Use pointer receivers when methods modify the receiver

## Next Steps

- Practice implementing interfaces
- Experiment with defer for resource management
- Try implementing recursive algorithms
- Create methods for custom types