package main

import "fmt"

func main() {
	// Variables and types
	var name string = "Go"
	age := 15 // short variable declaration
	isAwesome := true

	fmt.Printf("Hello, %s! You are %d years old and it's %t that you're awesome.\n", name, age, isAwesome)

	// Constants
	const pi = 3.14159
	const (
		daysInWeek    = 7
		hoursInDay    = 24
		minutesInHour = 60
	)

	fmt.Printf("Pi is approximately %.2f\n", pi)
	fmt.Printf("There are %d minutes in a day\n", hoursInDay*minutesInHour)

	// Basic types
	var (
		intVar    int     = 42
		floatVar  float64 = 3.14
		stringVar string  = "hello"
		boolVar   bool    = false
	)

	fmt.Printf("int: %d, float: %.2f, string: %s, bool: %t\n", intVar, floatVar, stringVar, boolVar)

	// Type conversions
	var x int = 42
	var y float64 = float64(x)
	var z string = fmt.Sprintf("%d", x)

	fmt.Printf("x: %d, y: %.0f, z: %s\n", x, y, z)

	// Operators
	a, b := 10, 3
	fmt.Printf("Arithmetic: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Comparison: %d > %d is %t\n", a, b, a > b)
	fmt.Printf("Logical: true && false = %t\n", true && false)

	// Control flow - if statement
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// Control flow - switch statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Control flow - for loop
	fmt.Println("Counting to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range over slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers in slice:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}
}
