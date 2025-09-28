package main

import "fmt"

// Person struct for demonstrating pointers
type Person struct {
	Name string
	Age  int
}

// Function that takes a value (copy)
func modifyValue(p Person) {
	p.Name = "Modified"
	p.Age = 99
	fmt.Printf("Inside modifyValue: %+v\n", p)
}

// Function that takes a pointer
func modifyPointer(p *Person) {
	p.Name = "Modified"
	p.Age = 99
	fmt.Printf("Inside modifyPointer: %+v\n", p)
}

// Function that returns a pointer
func newPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p // Return pointer to local variable
}

// Slice demonstration
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Printf("Inside modifySlice: %v\n", s)
}

// Map demonstration
func modifyMap(m map[string]int) {
	m["modified"] = 999
	fmt.Printf("Inside modifyMap: %v\n", m)
}

// Demonstrating pointer arithmetic (not allowed in Go)
func demonstrateNoPointerArithmetic() {
	// This would be invalid in Go:
	// p := &x
	// p++  // Not allowed

	fmt.Println("Go does not support pointer arithmetic like C/C++")
}

// Nil pointer demonstration
func handleNilPointer(p *Person) {
	if p == nil {
		fmt.Println("Received nil pointer")
		return
	}
	fmt.Printf("Person: %+v\n", *p)
}

// Multiple levels of indirection
func doublePointer(pp **Person) {
	if pp == nil || *pp == nil {
		fmt.Println("Nil pointer dereference would occur")
		return
	}
	(*pp).Name = "Double Modified"
	fmt.Printf("Inside doublePointer: %+v\n", **pp)
}

// Memory allocation with new()
func demonstrateNew() {
	// new() allocates zeroed memory and returns pointer
	p := new(Person)
	fmt.Printf("new(Person): %+v\n", *p)

	// Equivalent to:
	var p2 *Person = new(Person)
	fmt.Printf("var p2 *Person = new(Person): %+v\n", *p2)
}

// Memory allocation with make() vs new()
func demonstrateMakeVsNew() {
	// make() is for slices, maps, channels
	slice := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// new() is for any type
	slicePtr := new([]int)
	fmt.Printf("new([]int): %v\n", *slicePtr) // nil slice

	// To initialize:
	*slicePtr = make([]int, 3)
	fmt.Printf("After make: %v\n", *slicePtr)
}

// Struct embedding and pointers
type Address struct {
	Street string
	City   string
}

type Employee struct {
	Name    string
	Address *Address // Pointer to embedded struct
}

// Method with pointer receiver
func (e *Employee) SetAddress(street, city string) {
	if e.Address == nil {
		e.Address = &Address{}
	}
	e.Address.Street = street
	e.Address.City = city
}

// Demonstrating reference types
func referenceTypes() {
	// Slices are reference types
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // Both point to same underlying array

	slice2[0] = 999
	fmt.Printf("slice1 after modifying slice2: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	// Maps are reference types
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map1 // Both point to same map

	map2["c"] = 3
	fmt.Printf("map1 after modifying map2: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)

	// Channels are reference types (not shown here)
}

// Demonstrating value types
func valueTypes() {
	// Structs are value types
	person1 := Person{Name: "Alice", Age: 25}
	person2 := person1 // Copy of the struct

	person2.Name = "Bob"
	fmt.Printf("person1 after modifying person2: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)

	// Basic types are value types
	x := 42
	y := x // Copy of the value
	y = 99
	fmt.Printf("x after modifying y: %d\n", x)
	fmt.Printf("y: %d\n", y)
}

// Pointer to interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func makeSpeak(s Speaker) {
	fmt.Printf("%s says: %s\n", s.(Dog).Name, s.Speak())
}

func main() {
	// Basic pointer operations
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Original person: %+v\n", person)

	// Pass by value
	modifyValue(person)
	fmt.Printf("After modifyValue: %+v\n", person)

	// Pass by reference
	modifyPointer(&person)
	fmt.Printf("After modifyPointer: %+v\n", person)

	// Returning pointers
	newP := newPerson("Bob", 30)
	fmt.Printf("New person: %+v\n", *newP)

	// Slice modification (slices are reference types)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("After modifySlice: %v\n", numbers)

	// Map modification (maps are reference types)
	scores := map[string]int{"alice": 100, "bob": 95}
	fmt.Printf("Original map: %v\n", scores)
	modifyMap(scores)
	fmt.Printf("After modifyMap: %v\n", scores)

	// Nil pointer handling
	var nilPerson *Person
	handleNilPointer(nilPerson)

	validPerson := &Person{Name: "Charlie", Age: 35}
	handleNilPointer(validPerson)

	// Double pointers
	pp := &validPerson
	doublePointer(pp)
	fmt.Printf("After doublePointer: %+v\n", *validPerson)

	// new() function
	demonstrateNew()

	// make() vs new()
	demonstrateMakeVsNew()

	// Struct with pointer field
	emp := Employee{Name: "David"}
	emp.SetAddress("123 Main St", "Anytown")
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Address: %+v\n", *emp.Address)

	// Reference vs value types
	fmt.Println("\n--- Reference Types ---")
	referenceTypes()

	fmt.Println("\n--- Value Types ---")
	valueTypes()

	// Pointer to interface
	dog := Dog{Name: "Buddy"}
	var speaker Speaker = dog
	makeSpeak(speaker)

	// Pointer syntax variations
	var ptr *int
	num := 42

	// Different ways to get pointer
	ptr = &num // Address-of operator
	fmt.Printf("ptr points to: %d\n", *ptr)

	// Pointer to pointer
	ptrPtr := &ptr
	fmt.Printf("ptrPtr points to value: %d\n", **ptrPtr)
}
