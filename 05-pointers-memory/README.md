# 05 - Pointers and Memory Management

This example demonstrates Go's pointer system, memory management, and the distinction between value and reference types.

## Key Concepts

### Pointer Basics

#### Declaring Pointers
```go
var ptr *int        // nil pointer to int
var p *Person       // nil pointer to Person struct
```

#### Address-of Operator
```go
x := 42
ptr := &x          // ptr points to x
fmt.Println(*ptr)  // dereference: prints 42
```

#### Dereference Operator
```go
*ptr = 99          // modify value pointed to by ptr
```

### Value vs Reference Semantics

#### Value Types (copied when assigned)
- Basic types: `int`, `float64`, `bool`, `string`
- Structs: `type Person struct { Name string; Age int }`
- Arrays: `[5]int`

#### Reference Types (share underlying data)
- Slices: `[]int`
- Maps: `map[string]int`
- Channels: `chan int`
- Pointers: `*int`
- Interfaces: `io.Reader`
- Functions: `func(int) int`

### Memory Allocation

#### new() Function
```go
p := new(Person)   // Allocates zeroed memory, returns *Person
```
- Allocates memory for the type
- Returns pointer to zeroed memory
- Used for any type

#### make() Function
```go
slice := make([]int, 3, 10)    // len=3, cap=10
m := make(map[string]int)      // empty map
ch := make(chan int, 5)        // buffered channel
```
- Only for slices, maps, and channels
- Initializes the internal data structures
- Returns initialized (not zeroed) value

### Pointer Operations

#### Passing by Value vs Reference
```go
func modifyValue(p Person) {    // Receives copy
    p.Name = "Modified"         // Only modifies copy
}

func modifyPointer(p *Person) { // Receives pointer
    p.Name = "Modified"         // Modifies original
}
```

#### Returning Pointers
```go
func newPerson(name string) *Person {
    p := Person{Name: name}
    return &p  // Safe: returns pointer to heap-allocated value
}
```

### Common Patterns

#### Nil Pointer Checks
```go
func process(p *Person) {
    if p == nil {
        return // handle nil case
    }
    // safe to dereference p
}
```

#### Pointer to Interface
```go
var r io.Reader = &MyReader{}
```
- Interfaces can hold pointers or values
- Method sets determine what you can do

### Memory Management

Go has automatic memory management (garbage collection):
- No manual `free()` or `delete`
- No pointer arithmetic (unlike C/C++)
- No memory leaks in normal usage
- GC runs concurrently

### Method Receivers

#### Value Receiver
```go
func (p Person) String() string {
    return fmt.Sprintf("...")
}
```
- Works on copy of the receiver
- Cannot modify the original receiver

#### Pointer Receiver
```go
func (p *Person) Birthday() {
    p.Age++  // Modifies original
}
```
- Works on pointer to receiver
- Can modify the original receiver
- More efficient for large structs

### Common Pitfalls

#### 1. Nil Pointer Dereference
```go
var p *Person
p.Name = "Alice"  // Panic: runtime error
```

#### 2. Modifying Copies
```go
func (p Person) SetName(name string) {
    p.Name = name  // Only modifies copy
}
```

#### 3. Pointer to Local Variable
```go
func bad() *int {
    x := 42
    return &x  // Dangerous: x goes out of scope
}
```

#### 4. Slices and Maps are Reference Types
```go
slice1 := []int{1, 2, 3}
slice2 := slice1        // Both point to same array
slice2[0] = 999         // Modifies both slices
```

## Running the Example

```bash
go run main.go
```

## Expected Output

```
Original person: {Name:Alice Age:25}
Inside modifyValue: {Name:Modified Age:99}
After modifyValue: {Name:Alice Age:25}
Inside modifyPointer: &{Name:Modified Age:99}
After modifyPointer: {Name:Modified Age:99}
New person: {Name:Bob Age:30}
Original slice: [1 2 3 4 5]
Inside modifySlice: [999 2 3 4 5]
After modifySlice: [999 2 3 4 5]
Original map: map[alice:100 bob:95]
Inside modifyMap: map[alice:100 bob:95 modified:999]
After modifyMap: map[alice:100 bob:95 modified:999]
Received nil pointer
Person: {Name:Charlie Age:35}
Inside doublePointer: {Name:Double Modified Age:35}
After doublePointer: {Name:Double Modified Age:35}
new(Person): {Name: Age:0}
var p2 *Person = new(Person): {Name: Age:0}
make([]int, 3, 10): [0 0 0], len=3, cap=10
new([]int): []
After make: [0 0 0]
Employee: {Name:David Address:0xc0000...}
Address: {Street:123 Main St City:Anytown}

--- Reference Types ---
slice1 after modifying slice2: [999 2 3]
slice2: [999 2 3]
map1 after modifying map2: map[a:1 b:2 c:3]
map2: map[a:1 b:2 c:3]

--- Value Types ---
person1 after modifying person2: {Name:Alice Age:25}
person2: {Name:Bob Age:25}
x after modifying y: 42
y: 99
Buddy says: Woof!
ptr points to: 42
ptrPtr points to value: 42
```

## Best Practices

1. **Use pointers when:**
   - Method needs to modify the receiver
   - Struct is large (avoid copying overhead)
   - Need to represent optional values (nil)

2. **Use values when:**
   - Struct is small
   - Method doesn't need to modify receiver
   - Want immutability

3. **Always check for nil** before dereferencing pointers

4. **Use make()** for slices, maps, and channels

5. **Use new()** for structs and basic types when you need a pointer

## Interview Questions

- What's the difference between `new()` and `make()`?
- When should you use pointer receivers vs value receivers?
- How do slices and maps differ from arrays in terms of memory?
- What happens when you pass a slice to a function?

## Next Steps

- Practice implementing data structures with pointers
- Experiment with different receiver types on methods
- Try implementing linked lists or trees using pointers
- Practice memory-efficient coding patterns