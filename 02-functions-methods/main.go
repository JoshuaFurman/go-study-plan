package main

import (
	"fmt"
	"math"
)

// Basic function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Function with named return values
func rectangleProperties(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Variadic function - accepts variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Person struct to demonstrate methods
type Person struct {
	Name string
	Age  int
}

// Value receiver method
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Pointer receiver method
func (p *Person) Birthday() {
	p.Age++
}

// Rectangle struct for method demonstration
type Rectangle struct {
	Width, Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Interface definition
type Shape interface {
	Area() float64
}

// Circle struct implementing Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Function demonstrating defer
func demonstrateDefer() {
	fmt.Println("Start")
	defer fmt.Println("Deferred 1") // Executes last
	defer fmt.Println("Deferred 2") // Executes second to last
	fmt.Println("End")
}

// Function demonstrating panic and recover
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

// Anonymous function example
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	// Basic function call
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Named return values
	area, perimeter := rectangleProperties(4, 5)
	fmt.Printf("Rectangle 4x5: Area=%.1f, Perimeter=%.1f\n", area, perimeter)

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1+2+3+4+5 = %d\n", total)

	// Recursive function
	fact := factorial(5)
	fmt.Printf("5! = %d\n", fact)

	// Methods
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.String())

	person.Birthday()
	fmt.Println("After birthday:", person.String())

	// Rectangle methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle area: %.1f\n", rect.Area())

	rect.Scale(2)
	fmt.Printf("After scaling by 2: Width=%.1f, Height=%.1f, Area=%.1f\n", rect.Width, rect.Height, rect.Area())

	// Interface usage
	var shape Shape = Circle{Radius: 3}
	fmt.Printf("Circle area: %.2f\n", shape.Area())

	// Defer demonstration
	demonstrateDefer()

	// Panic and recover
	result, err = safeDivision(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	result, err = safeDivision(10, 0)
	if err != nil {
		fmt.Printf("Recovered from panic: %v\n", err)
	}

	// Anonymous function
	addResult := applyOperation(3, 4, func(a, b int) int { return a + b })
	fmt.Printf("3 + 4 = %d (using anonymous function)\n", addResult)

	multiplyResult := applyOperation(3, 4, func(a, b int) int { return a * b })
	fmt.Printf("3 * 4 = %d (using anonymous function)\n", multiplyResult)
}
