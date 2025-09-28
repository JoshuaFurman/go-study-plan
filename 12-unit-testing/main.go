package main

import (
	"fmt"
	"math"
)

// Calculator provides basic arithmetic operations
type Calculator struct{}

// Add returns the sum of two numbers
func (c Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
// Returns error if dividing by zero
func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// IsEven returns true if number is even
func (c Calculator) IsEven(n int) bool {
	return n%2 == 0
}

// Factorial returns the factorial of n (n!)
// Returns error for negative numbers
func (c Calculator) Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// IsPrime returns true if n is a prime number
func (c Calculator) IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Max returns the maximum of two numbers
func (c Calculator) Max(a, b float64) float64 {
	return math.Max(a, b)
}

// Min returns the minimum of two numbers
func (c Calculator) Min(a, b float64) float64 {
	return math.Min(a, b)
}

func main() {
	calc := Calculator{}

	fmt.Println("=== Calculator Demo ===")
	fmt.Printf("5 + 3 = %.0f\n", calc.Add(5, 3))
	fmt.Printf("10 - 4 = %.0f\n", calc.Subtract(10, 4))
	fmt.Printf("6 * 7 = %.0f\n", calc.Multiply(6, 7))

	if result, err := calc.Divide(15, 3); err == nil {
		fmt.Printf("15 / 3 = %.0f\n", result)
	}

	fmt.Printf("Is 4 even? %t\n", calc.IsEven(4))
	fmt.Printf("Is 7 even? %t\n", calc.IsEven(7))

	if fact, err := calc.Factorial(5); err == nil {
		fmt.Printf("5! = %d\n", fact)
	}

	fmt.Printf("Is 7 prime? %t\n", calc.IsPrime(7))
	fmt.Printf("Is 9 prime? %t\n", calc.IsPrime(9))

	fmt.Printf("Max(10, 20) = %.0f\n", calc.Max(10, 20))
	fmt.Printf("Min(10, 20) = %.0f\n", calc.Min(10, 20))

	fmt.Println("\nRun 'go test' to see the unit tests in action!")
}
