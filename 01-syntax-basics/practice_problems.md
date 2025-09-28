# Practice Problems - Syntax Basics

These practice problems reinforce the fundamental Go syntax concepts covered in this example. Try solving them on your own before checking the solutions.

## Problem 1: Temperature Converter

**Description:** Write a program that converts temperatures between Celsius and Fahrenheit. The program should:
- Take a temperature value and a unit ('C' for Celsius, 'F' for Fahrenheit)
- Convert to the other unit
- Display the result with 2 decimal places

**Formula:**
- Celsius to Fahrenheit: `F = (C × 9/5) + 32`
- Fahrenheit to Celsius: `C = (F - 32) × 5/9`

**Examples:**
- Input: 25°C → Output: 77.00°F
- Input: 77°F → Output: 25.00°C

**Constraints:**
- Temperature values will be between -100 and 100
- Unit will be either 'C' or 'F' (case sensitive)

### Solution

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <temperature> <unit>")
		fmt.Println("Example: go run main.go 25 C")
		return
	}

	tempStr := os.Args[1]
	unit := os.Args[2]

	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		fmt.Println("Error: Invalid temperature value")
		return
	}

	var result float64
	var resultUnit string

	if unit == "C" {
		// Celsius to Fahrenheit
		result = (temp * 9 / 5) + 32
		resultUnit = "F"
	} else if unit == "F" {
		// Fahrenheit to Celsius
		result = (temp - 32) * 5 / 9
		resultUnit = "C"
	} else {
		fmt.Println("Error: Unit must be 'C' or 'F'")
		return
	}

	fmt.Printf("%.2f°%s = %.2f°%s\n", temp, unit, result, resultUnit)
}
```

**Explanation:**
1. Parse command line arguments for temperature and unit
2. Convert string temperature to float64
3. Apply appropriate conversion formula based on unit
4. Display result with 2 decimal places

## Problem 2: Simple Calculator

**Description:** Create a basic calculator that performs addition, subtraction, multiplication, and division on two numbers.

**Requirements:**
- Take two numbers and an operator (+, -, *, /) as input
- Perform the calculation
- Handle division by zero
- Display the result

**Examples:**
- Input: 10 + 5 → Output: 10 + 5 = 15
- Input: 10 / 0 → Output: Error: Division by zero

**Constraints:**
- Numbers will be valid floats
- Operator will be one of: +, -, *, /

### Solution

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		fmt.Println("Example: go run main.go 10 + 5")
		return
	}

	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid number format")
		return
	}

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator. Use +, -, *, or /")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
```

**Explanation:**
1. Parse command line arguments for two numbers and operator
2. Convert strings to float64
3. Use switch statement to perform appropriate operation
4. Handle division by zero error
5. Display formatted result