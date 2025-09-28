# Practice Problems - Error Handling

These practice problems reinforce Go's error handling patterns. Focus on proper error propagation and handling edge cases.

## Problem 1: Safe Division with Multiple Error Types

**Description:** Implement a safe division function that handles division by zero and overflow conditions. Return different error types for different failure modes.

**Requirements:**
- Define custom error types for different error conditions
- Handle division by zero
- Handle overflow (when result would be too large/small)
- Return (result, error) pairs

**Examples:**
- divide(10, 2) → 5.0, nil
- divide(10, 0) → 0.0, DivisionByZeroError
- divide(1e308, 0.1) → 0.0, OverflowError

**Constraints:**
- Use custom error types that implement error interface
- Handle float64 overflow using math.IsInf

### Solution

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

// Custom error types
type DivisionByZeroError struct {
	Dividend float64
}

func (e DivisionByZeroError) Error() string {
	return fmt.Sprintf("division by zero: cannot divide %.2f by zero", e.Dividend)
}

type OverflowError struct {
	Dividend, Divisor float64
}

func (e OverflowError) Error() string {
	return fmt.Sprintf("overflow: %.2f / %.2f would cause overflow", e.Dividend, e.Divisor)
}

func safeDivide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, DivisionByZeroError{Dividend: dividend}
	}

	result := dividend / divisor

	// Check for overflow
	if math.IsInf(result, 0) {
		return 0, OverflowError{Dividend: dividend, Divisor: divisor}
	}

	return result, nil
}

func main() {
	// Test cases
	testCases := []struct {
		dividend, divisor float64
	}{
		{10, 2},
		{10, 0},
		{1e308, 0.1},
		{-1e308, -0.1},
		{5, 3},
	}

	for _, tc := range testCases {
		result, err := safeDivide(tc.dividend, tc.divisor)
		fmt.Printf("safeDivide(%.2f, %.2f) = ", tc.dividend, tc.divisor)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("%.2f\n", result)
		}
	}
}
```

**Explanation:**
1. Define custom error types that implement error interface
2. Check for division by zero first
3. Perform division and check for overflow using math.IsInf
4. Return appropriate error types for different failure modes
5. Use struct embedding for error context

## Problem 2: File Processor with Error Handling

**Description:** Create a file processor that reads numbers from a file, processes them, and writes results to another file. Handle all possible error conditions gracefully.

**Requirements:**
- Read numbers from input file (one per line)
- Calculate sum and average
- Write results to output file
- Handle file I/O errors, parsing errors, and empty files
- Use proper error wrapping and context

**Examples:**
- Input file: "1\n2\n3\n" → Output: "Sum: 6, Average: 2.00"
- Invalid input: "1\nabc\n3" → Error: parsing error on line 2

**Constraints:**
- Use os.OpenFile and bufio for file operations
- Handle all error cases (file not found, permission denied, etc.)
- Provide meaningful error messages

### Solution

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProcessingError struct {
	Operation string
	Filename  string
	Line       int
	Err       error
}

func (e ProcessingError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("%s error in %s at line %d: %v", e.Operation, e.Filename, e.Line, e.Err)
	}
	return fmt.Sprintf("%s error in %s: %v", e.Operation, e.Filename, e.Err)
}

func (e ProcessingError) Unwrap() error {
	return e.Err
}

func processFile(inputFile, outputFile string) error {
	// Open input file
	input, err := os.Open(inputFile)
	if err != nil {
		return ProcessingError{
			Operation: "open",
			Filename:  inputFile,
			Err:       err,
		}
	}
	defer input.Close()

	// Read and process numbers
	scanner := bufio.NewScanner(input)
	var numbers []float64
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue // Skip empty lines
		}

		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return ProcessingError{
				Operation: "parse",
				Filename:  inputFile,
				Line:       lineNum,
				Err:       err,
			}
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return ProcessingError{
			Operation: "read",
			Filename:  inputFile,
			Err:       err,
		}
	}

	if len(numbers) == 0 {
		return ProcessingError{
			Operation: "process",
			Filename:  inputFile,
			Err:       fmt.Errorf("no valid numbers found"),
		}
	}

	// Calculate statistics
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	average := sum / float64(len(numbers))

	// Write output file
	output, err := os.Create(outputFile)
	if err != nil {
		return ProcessingError{
			Operation: "create",
			Filename:  outputFile,
			Err:       err,
		}
	}
	defer output.Close()

	_, err = fmt.Fprintf(output, "Count: %d\nSum: %.2f\nAverage: %.2f\n",
		len(numbers), sum, average)
	if err != nil {
		return ProcessingError{
			Operation: "write",
			Filename:  outputFile,
			Err:       err,
		}
	}

	return nil
}

func main() {
	// Create test input file
	inputContent := "1.5\n2.0\n\n3.5\n4.0\n"
	err := os.WriteFile("input.txt", []byte(inputContent), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %v\n", err)
		return
	}
	defer os.Remove("input.txt")
	defer os.Remove("output.txt")

	// Process file
	err = processFile("input.txt", "output.txt")
	if err != nil {
		fmt.Printf("Processing failed: %v\n", err)
		return
	}

	// Read and display output
	output, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Printf("Error reading output: %v\n", err)
		return
	}

	fmt.Println("Processing successful!")
	fmt.Println("Output:")
	fmt.Println(string(output))
}
```

**Explanation:**
1. Define custom error type with context (operation, filename, line number)
2. Handle all file operations with proper error checking
3. Use defer for resource cleanup
4. Parse numbers with error handling for invalid input
5. Calculate statistics only if valid numbers found
6. Write results to output file with error handling
7. Provide detailed error messages for debugging