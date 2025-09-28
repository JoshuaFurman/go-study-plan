# 04 - Structs and Interfaces

This example demonstrates Go's struct types, methods, and interfaces - the building blocks of object-oriented programming in Go.

## Key Concepts

### Structs

#### Basic Struct Definition
```go
type Person struct {
    Name string
    Age  int
    City string
}
```
- Collection of fields with types
- Fields are accessed with dot notation: `person.Name`

#### Struct Literals
```go
person := Person{Name: "Alice", Age: 25, City: "New York"}
// or
person := Person{"Alice", 25, "New York"}  // positional
```

#### Struct Tags
```go
type Employee struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Position string `json:"position"`
    Salary   float64 `json:"salary"`
}
```
- Metadata attached to struct fields
- Used by encoding packages (JSON, XML, etc.)
- Backtick-quoted strings

#### Embedded Structs
```go
type User struct {
    Name    string
    Age     int
    Address     // Embedded (no field name)
    Contact     // Embedded (no field name)
}
```
- Fields from embedded structs are "promoted"
- Can access as `user.City` instead of `user.Address.City`
- Useful for composition over inheritance

### Methods

#### Value vs Pointer Receivers
```go
// Value receiver - works on copy
func (p Person) String() string {
    return fmt.Sprintf("...")
}

// Pointer receiver - modifies original
func (p *Person) Birthday() {
    p.Age++
}
```
- Value receivers work on copies (can't modify original)
- Pointer receivers work on original (can modify)
- Use pointer receivers when method modifies receiver
- Use value receivers for immutable operations

### Interfaces

#### Interface Definition
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```
- Define method signatures (no implementation)
- Types implement interfaces by implementing their methods
- No explicit "implements" keyword needed

#### Empty Interface
```go
type Any interface{}
```
- Can hold any type
- Equivalent to `interface{}`
- Used when type is unknown or variable

#### Interface Composition
```go
type ReadWriter interface {
    Reader
    Writer
}
```
- Interfaces can embed other interfaces
- Composed interfaces inherit all methods

### Type Assertions and Type Switches

#### Type Assertion
```go
if s, ok := shape.(Shape); ok {
    area := s.Area()
}
```
- Check if interface value holds specific type
- Returns value and boolean indicating success
- Panics if assertion fails without `ok`

#### Type Switch
```go
switch s := shape.(type) {
case Circle:
    // s is of type Circle
case Rectangle:
    // s is of type Rectangle
default:
    // unknown type
}
```
- Switch on the concrete type of interface value
- `s` takes the concrete type in each case

### Implementation Details

#### Duck Typing
Go uses "duck typing" for interfaces:
> "If it walks like a duck and quacks like a duck, it must be a duck"

Types implement interfaces implicitly by having the right methods.

#### Interface Values
Interface values are represented as:
- Concrete type (what it actually is)
- Concrete value (the actual data)

#### Nil Interfaces
- Interface value is nil if both type and value are nil
- Method calls on nil interfaces panic
- Always check for nil before calling methods

## Running the Example

```bash
go run main.go
```

## Expected Output

```
Person: {Alice 25 New York}
String method: Person{Name: Alice, Age: 25, City: New York}
After birthday: Person{Name: Alice, Age: 26, City: New York}
Original person: Person{Name: Alice, Age: 26, City: New York}
Moved person: Person{Name: Alice, Age: 26, City: San Francisco}
Employee: {ID:1 Name:Bob Position:Developer Salary:75000.5}
User: {Name:Charlie Age:30 Address:{Street:123 Main St City:Boston Country:USA} Contact:{Email:charlie@example.com Phone:555-0123}}
User city: Boston
User email: charlie@example.com

Shape calculations:
main.Circle: Area=78.54, Perimeter=31.42
main.Rectangle: Area=50.00, Perimeter=30.00
main.Triangle: Area=12.00, Perimeter=16.00

Empty interface examples:
Type: int, Value: 42
Type: string, Value: hello
Type: main.Person, Value: {Alice 26 New York}
Type: main.Circle, Value: {3}

Wrote 12 bytes
Read 12 bytes: Hello, World!
Circle area via assertion: 28.27

Shape descriptions:
Circle with radius 3.00
Rectangle 4.00x6.00
Triangle with base 3.00 and height 4.00
Unknown shape

Interface slice:
Index 0: 42 (type: int)
Index 1: hello (type: string)
Index 2: {Alice 26 New York} (type: main.Person)
Index 3: {3} (type: main.Circle)
```

## Design Principles

1. **Composition over Inheritance**: Use embedding instead of inheritance
2. **Small Interfaces**: Keep interfaces focused on single responsibilities
3. **Accept Interfaces, Return Structs**: Functions should accept interfaces but return concrete types
4. **Implicit Implementation**: No need to declare interface implementation

## Common Patterns

1. **Interface Segregation**: Small, focused interfaces
2. **Dependency Injection**: Pass interfaces to constructors
3. **Type Assertions**: Use with `ok` to check success
4. **Method Sets**: Different method sets for value vs pointer receivers

## Best Practices

- Use pointer receivers when methods modify the receiver
- Keep interfaces small and focused
- Use struct embedding for composition
- Handle nil interface values safely
- Use type assertions with the comma ok idiom

## Next Steps

- Implement custom interfaces for your types
- Practice type assertions and type switches
- Experiment with interface composition
- Try implementing common design patterns using interfaces