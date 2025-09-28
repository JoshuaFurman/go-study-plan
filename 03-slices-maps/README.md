# 03 - Slices and Maps

This example demonstrates Go's slice and map data structures, which are fundamental to Go programming.

## Key Concepts

### Arrays vs Slices

#### Arrays (Fixed Size)
```go
var numbers [5]int          // Array of 5 integers
primes := [5]int{2, 3, 5, 7, 11}  // Array literal
```
- Fixed size determined at compile time
- Size is part of the type: `[5]int` ≠ `[3]int`
- Rarely used directly; slices are preferred

#### Slices (Dynamic Arrays)
```go
var slice []int                    // Nil slice
slice = []int{1, 2, 3, 4, 5}      // Slice literal
slice = make([]int, 3, 10)        // len=3, cap=10
```
- Dynamic size, can grow and shrink
- Built on top of arrays
- Most commonly used data structure

### Slice Operations

#### Creating Slices from Arrays
```go
array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
slice1 := array[2:5]  // elements 2, 3, 4
slice2 := array[:3]   // elements 0, 1, 2
slice3 := array[7:]   // elements 7, 8, 9
```
- `slice[low:high]` creates a slice from array
- `slice[:high]` from start to high
- `slice[low:]` from low to end

#### Append Operation
```go
slice = append(slice, 6, 7, 8)        // Add elements
slice = append(slice, moreNumbers...) // Add another slice
```
- `append` returns a new slice
- May allocate new underlying array if capacity exceeded
- Use `...` to expand slice into individual elements

#### Length vs Capacity
```go
slice := make([]int, 3, 10)  // len=3, cap=10
fmt.Println(len(slice))      // 3
fmt.Println(cap(slice))      // 10
```
- `len()`: number of elements in slice
- `cap()`: capacity of underlying array
- Slice can grow up to capacity without reallocation

#### Common Slice Operations
```go
// Copy slice
destination := make([]int, len(source))
copy(destination, source)

// Remove element at index
slice = append(slice[:index], slice[index+1:]...)

// Insert element at index
slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
```

### Maps (Hash Tables)

#### Creating Maps
```go
var person map[string]string         // Nil map
person = make(map[string]string)     // Empty map
person = map[string]string{          // Map literal
    "name": "Alice",
    "age":  "25",
}
```
- Key-value pairs
- Keys must be comparable types
- Values can be any type

#### Map Operations
```go
// Add/update
person["city"] = "New York"

// Access
name := person["name"]

// Check existence
age, exists := person["age"]
if exists {
    // use age
}

// Delete
delete(person, "city")

// Iterate
for key, value := range person {
    fmt.Printf("%s: %s\n", key, value)
}
```

### Sorting

#### Built-in Sort Functions
```go
sort.Ints(slice)        // Sort []int
sort.Strings(slice)     // Sort []string
sort.Float64s(slice)    // Sort []float64
```

#### Custom Sorting with sort.Slice
```go
sort.Slice(people, func(i, j int) bool {
    return people[i].age < people[j].age
})
```
- Use `sort.Slice` for custom sorting logic
- Comparator function defines sort order

#### Reverse Sorting
```go
sort.Sort(sort.Reverse(sort.IntSlice(slice)))
```

### Multi-dimensional Data Structures

#### Slice of Slices
```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```
- Each element is itself a slice
- Useful for matrices, grids, etc.

#### Map of Maps
```go
users := map[string]map[string]string{
    "user1": {"name": "Alice", "role": "admin"},
    "user2": {"name": "Bob", "role": "user"},
}
```
- Nested maps for complex data structures

## Running the Example

```bash
go run main.go
```

## Expected Output

```
Array: [1 2 0 0 0]
Primes array: [2 3 5 7 11]
Empty slice: [] len: 0 cap: 0
Slice: [1 2 3 4 5] len: 5 cap: 5
slice1: [2 3 4]
slice2: [0 1 2]
slice3: [7 8 9]
After append: [1 2 3 4 5 6 7 8] len: 8 cap: 10
After append slice: [1 2 3 4 5 6 7 8 9 10 11] len: 11 cap: 22
Make slice: [0 0 0] len: 3 cap: 10
Copied slice: [1 2 3 4 5]
After removing index 3: [1 2 3 5 6 7 8 9 10]
After inserting 99 at index 2: [1 2 99 3 5 6 7 8 9 10]
Nil map: map[]
Person map: map[age:25 city:New York name:Alice]
Student map: map[active:true grades:[85 92 78] name:Bob]
Name: Alice
Age: 25
After deleting city: map[age:25 name:Alice]
Iterating over person map:
  age: 25
  name: Alice
Users: map[user1:map[name:Alice role:admin] user2:map[name:Bob role:user]]
Unsorted: [3 1 4 1 5 9 2 6]
Sorted ascending: [1 1 2 3 4 5 6 9]
Sorted descending: [9 6 5 4 3 2 1 1]
Unsorted words: [zebra apple banana cherry]
Sorted words: [apple banana cherry zebra]
People before sorting:
  Alice: 25
  Bob: 30
  Charlie: 20
People after sorting by age:
  Charlie: 20
  Alice: 25
  Bob: 30
Matrix:
  [1 2 3]
  [4 5 6]
  [7 8 9]
matrix[1][2] = 6
```

## Performance Considerations

1. **Slice Capacity**: Pre-allocate with `make()` when size is known
2. **Append Efficiency**: Appending to slice with available capacity is O(1)
3. **Map Access**: Average O(1) time complexity
4. **Sorting**: `sort.Slice` is efficient for custom sorting

## Common Pitfalls

1. **Nil vs Empty Slice**: `var s []int` vs `s := []int{}`
2. **Slice Bounds**: Accessing beyond `len(slice)` causes panic
3. **Map Key Types**: Only comparable types can be map keys
4. **Concurrent Map Access**: Maps are not thread-safe

## Next Steps

- Practice slice manipulation operations
- Implement common algorithms using slices
- Experiment with different map key/value types
- Try implementing custom sorting logic