package main

import (
	"fmt"
	"math"
)

// Basic struct definition
type Person struct {
	Name string
	Age  int
	City string
}

// Struct with tags (for JSON, etc.)
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

// Struct with embedded fields
type Address struct {
	Street  string
	City    string
	Country string
}

type Contact struct {
	Email string
	Phone string
}

type User struct {
	Name    string
	Age     int
	Address // Embedded struct
	Contact // Embedded struct
}

// Method with value receiver
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, City: %s}", p.Name, p.Age, p.City)
}

// Method with pointer receiver
func (p *Person) Birthday() {
	p.Age++
}

// Method that returns a modified copy
func (p Person) MoveTo(newCity string) Person {
	p.City = newCity
	return p
}

// Interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle implements Shape
type Triangle struct {
	Base, Height, Side1, Side2 float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Side1 + t.Side2
}

// Empty interface - can hold any type
type Any interface{}

// Function that accepts any type
func printType(value Any) {
	fmt.Printf("Type: %T, Value: %v\n", value, value)
}

// Interface with multiple methods
type Writer interface {
	Write(data []byte) (int, error)
}

type Reader interface {
	Read(data []byte) (int, error)
}

// ReadWriter combines multiple interfaces
type ReadWriter interface {
	Reader
	Writer
}

// Simple implementation
type Buffer struct {
	data []byte
}

func (b *Buffer) Write(data []byte) (int, error) {
	b.data = append(b.data, data...)
	return len(data), nil
}

func (b *Buffer) Read(data []byte) (int, error) {
	if len(b.data) == 0 {
		return 0, fmt.Errorf("buffer is empty")
	}

	n := copy(data, b.data)
	b.data = b.data[n:]
	return n, nil
}

// Type assertion function
func getArea(shape Any) float64 {
	if s, ok := shape.(Shape); ok {
		return s.Area()
	}
	return 0
}

// Type switch
func describeShape(shape Any) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle with radius %.2f\n", s.Radius)
	case Rectangle:
		fmt.Printf("Rectangle %.2fx%.2f\n", s.Width, s.Height)
	case Triangle:
		fmt.Printf("Triangle with base %.2f and height %.2f\n", s.Base, s.Height)
	default:
		fmt.Println("Unknown shape")
	}
}

// Interface{} usage (pre-generics)
func printSlice(slice []interface{}) {
	for i, v := range slice {
		fmt.Printf("Index %d: %v (type: %T)\n", i, v, v)
	}
}

func main() {
	// Basic struct usage
	person := Person{Name: "Alice", Age: 25, City: "New York"}
	fmt.Println("Person:", person)
	fmt.Println("String method:", person.String())

	// Pointer receiver method
	person.Birthday()
	fmt.Println("After birthday:", person.String())

	// Value receiver method (returns copy)
	newPerson := person.MoveTo("San Francisco")
	fmt.Println("Original person:", person.String())
	fmt.Println("Moved person:", newPerson.String())

	// Struct with tags
	employee := Employee{
		ID:       1,
		Name:     "Bob",
		Position: "Developer",
		Salary:   75000.50,
	}
	fmt.Printf("Employee: %+v\n", employee)

	// Embedded structs
	user := User{
		Name: "Charlie",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
		Contact: Contact{
			Email: "charlie@example.com",
			Phone: "555-0123",
		},
	}

	fmt.Printf("User: %+v\n", user)
	fmt.Println("User city:", user.City)   // Promoted field
	fmt.Println("User email:", user.Email) // Promoted field

	// Interface usage
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
		Triangle{Base: 6, Height: 4, Side1: 5, Side2: 5},
	}

	fmt.Println("\nShape calculations:")
	for _, shape := range shapes {
		fmt.Printf("%T: Area=%.2f, Perimeter=%.2f\n",
			shape, shape.Area(), shape.Perimeter())
	}

	// Empty interface
	fmt.Println("\nEmpty interface examples:")
	printType(42)
	printType("hello")
	printType(person)
	printType(shapes[0])

	// ReadWriter interface
	buffer := &Buffer{}
	data := []byte("Hello, World!")

	n, err := buffer.Write(data)
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes\n", n)
	}

	readData := make([]byte, len(data))
	n, err = buffer.Read(readData)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(readData))
	}

	// Type assertion
	circle := Circle{Radius: 3}
	area := getArea(circle)
	fmt.Printf("Circle area via assertion: %.2f\n", area)

	// Type switch
	fmt.Println("\nShape descriptions:")
	describeShape(circle)
	describeShape(Rectangle{Width: 4, Height: 6})
	describeShape(Triangle{Base: 3, Height: 4, Side1: 5, Side2: 5})
	describeShape("not a shape")

	// Interface slice
	fmt.Println("\nInterface slice:")
	mixedSlice := []interface{}{42, "hello", person, circle}
	printSlice(mixedSlice)
}
