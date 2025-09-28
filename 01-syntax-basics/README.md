# 01 - Go Syntax Basics

This example covers the fundamental syntax elements of Go programming language.

## Key Concepts

### Variables and Declaration

Go has two ways to declare variables:

1. **var keyword**: `var name string = "Go"`
2. **Short declaration**: `age := 15` (type is inferred)

The `:=` operator can only be used inside functions, not at package level.

### Constants

Constants are declared with the `const` keyword and cannot be changed after declaration:

```go
const pi = 3.14159
const (
    daysInWeek = 7
    hoursInDay = 24
)
```

### Basic Types

Go has several built-in types:
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `string`
- `bool`

### Type Conversions

Go requires explicit type conversions (no implicit conversions):

```go
var x int = 42
var y float64 = float64(x)  // explicit conversion
```

### Operators

Go supports standard arithmetic, comparison, and logical operators:
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `==`, `!=`, `<`, `<=`, `>`, `>=`
- Logical: `&&`, `||`, `!`

### Control Flow

#### If Statements
```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

#### Switch Statements
Go's switch is more flexible than other languages:
```go
switch value {
case val1, val2:
    // code
case val3:
    // code
default:
    // code
}
```

#### For Loops
Go only has one loop construct:
```go
for i := 0; i < 10; i++ {
    // code
}
```

#### Range Loops
Used to iterate over arrays, slices, maps, channels:
```go
for index, value := range collection {
    // code
}
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
Hello, Go! You are 15 years old and it's true that you're awesome.
Pi is approximately 3.14
There are 1440 minutes in a day
int: 42, float: 3.14, string: hello, bool: false
x: 42, y: 42, z: 42
Arithmetic: 10 + 3 = 13
Comparison: 10 > 3 is true
Logical: true && false = false
You are a minor
It's a weekday
Counting to 5:
1 2 3 4 5
Numbers in slice:
Index 0: 1
Index 1: 2
Index 2: 3
Index 3: 4
Index 4: 5
```

## Common Mistakes

1. **Forgotten semicolons**: Go doesn't require semicolons, but they're automatically inserted
2. **Short declaration outside functions**: `:=` only works inside functions
3. **Type mismatches**: Go is strongly typed, no implicit conversions
4. **Unused variables**: Go compiler will error on unused variables

## Next Steps

- Practice declaring variables with different types
- Experiment with different control flow patterns
- Try converting between different types