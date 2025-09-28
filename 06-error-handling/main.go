package main

import (
	"errors"
	"fmt"
	"os"
)

// Basic error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// Function that returns custom error
func validateUser(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "name", Message: "cannot be empty"}
	}
	if age < 0 {
		return ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return ValidationError{Field: "age", Message: "cannot be greater than 150"}
	}
	return nil
}

// Error wrapping with fmt.Errorf
func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Simulate some processing
	return fmt.Errorf("failed to process file %s: %w", filename, errors.New("simulated processing error"))
}

// Unwrapping errors
func handleFileError(err error) {
	fmt.Printf("Original error: %v\n", err)

	// Check if it's a wrapped error
	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
		fmt.Printf("Unwrapped error: %v\n", wrappedErr)
	}

	// Check for specific error types
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	}
}

// Custom error with additional context
type NetworkError struct {
	Op      string
	URL     string
	Err     error
	Timeout bool
}

func (e NetworkError) Error() string {
	if e.Timeout {
		return fmt.Sprintf("network %s %s: timeout", e.Op, e.URL)
	}
	return fmt.Sprintf("network %s %s: %v", e.Op, e.URL, e.Err)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}

// Function that returns network error
func fetchURL(url string) error {
	// Simulate network timeout
	return NetworkError{
		Op:      "GET",
		URL:     url,
		Err:     errors.New("connection refused"),
		Timeout: true,
	}
}

// Sentinel errors (predefined error values)
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// Function using sentinel errors
func findUser(id string) error {
	if id == "" {
		return ErrInvalidInput
	}
	if id == "notfound" {
		return ErrNotFound
	}
	if id == "private" {
		return ErrUnauthorized
	}
	fmt.Printf("Found user: %s\n", id)
	return nil
}

// Error handling with defer and recover (not recommended for normal errors)
func riskyOperation() (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("operation panicked: %v", r)
		}
	}()

	// Simulate a panic
	if true {
		panic("something went wrong")
	}

	return 42, nil
}

// Multiple error handling patterns
func processMultipleItems(items []string) error {
	var errs []error

	for i, item := range items {
		if err := validateItem(item); err != nil {
			errs = append(errs, fmt.Errorf("item %d: %w", i, err))
		}
	}

	if len(errs) > 0 {
		// Return combined error
		return fmt.Errorf("multiple validation errors: %v", errs)
	}

	return nil
}

func validateItem(item string) error {
	if len(item) < 3 {
		return errors.New("item too short")
	}
	return nil
}

// Best practice: return errors early
func goodErrorHandling(data string) (result string, err error) {
	if data == "" {
		return "", errors.New("data cannot be empty")
	}

	// Process data
	result = fmt.Sprintf("processed: %s", data)
	return result, nil
}

// Bad practice: nested error handling
func badErrorHandling(data string) (result string, err error) {
	if data != "" {
		// Deep nesting makes code hard to read
		if len(data) > 5 {
			result = fmt.Sprintf("processed: %s", data)
		} else {
			err = errors.New("data too short")
		}
	} else {
		err = errors.New("data cannot be empty")
	}
	return
}

func main() {
	fmt.Println("=== Basic Error Handling ===")

	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n=== Custom Error Types ===")

	// Custom error type
	err = validateUser("", 25)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		// Type assertion to access custom fields
		if valErr, ok := err.(ValidationError); ok {
			fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	err = validateUser("Alice", -5)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	fmt.Println("\n=== Error Wrapping ===")

	// Error wrapping
	err = processFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("Process file error: %v\n", err)
		handleFileError(err)
	}

	fmt.Println("\n=== Network Error Example ===")

	// Network error
	err = fetchURL("http://example.com")
	if err != nil {
		fmt.Printf("Network error: %v\n", err)
		// Check if it's a timeout
		if netErr, ok := err.(NetworkError); ok && netErr.Timeout {
			fmt.Println("This was a timeout error")
		}
	}

	fmt.Println("\n=== Sentinel Errors ===")

	// Sentinel errors
	testCases := []string{"", "notfound", "private", "validuser"}
	for _, id := range testCases {
		err := findUser(id)
		if err != nil {
			fmt.Printf("findUser(%s) error: %v\n", id, err)
			// Check for specific sentinel errors
			if errors.Is(err, ErrNotFound) {
				fmt.Println("  -> User not found")
			} else if errors.Is(err, ErrUnauthorized) {
				fmt.Println("  -> Access denied")
			}
		}
	}

	fmt.Println("\n=== Panic and Recover (Not Recommended) ===")

	// Panic and recover (use sparingly)
	resultInt, err := riskyOperation()
	if err != nil {
		fmt.Printf("Recovered from panic: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", resultInt)
	}

	fmt.Println("\n=== Multiple Error Handling ===")

	// Multiple errors
	items := []string{"a", "validitem", "b", "another"}
	err = processMultipleItems(items)
	if err != nil {
		fmt.Printf("Multiple errors: %v\n", err)
	}

	fmt.Println("\n=== Error Handling Best Practices ===")

	// Good practice
	resultStr, err := goodErrorHandling("")
	if err != nil {
		fmt.Printf("Good practice error: %v\n", err)
	}

	resultStr, err = goodErrorHandling("valid data")
	if err != nil {
		fmt.Printf("Good practice error: %v\n", err)
	} else {
		fmt.Printf("Good practice result: %s\n", resultStr)
	}

	// Bad practice (for comparison)
	resultStr, err = badErrorHandling("")
	if err != nil {
		fmt.Printf("Bad practice error: %v\n", err)
	}

	fmt.Println("\n=== Error Type Checking ===")

	// Different ways to check error types
	err = ValidationError{Field: "email", Message: "invalid format"}

	// Type assertion
	if valErr, ok := err.(ValidationError); ok {
		fmt.Printf("Type assertion: Field=%s, Message=%s\n", valErr.Field, valErr.Message)
	}

	// errors.Is for sentinel errors
	sentErr := ErrNotFound
	if errors.Is(sentErr, ErrNotFound) {
		fmt.Println("errors.Is: This is a NotFound error")
	}

	// errors.As for custom error types
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("errors.As: Field=%s, Message=%s\n", valErr.Field, valErr.Message)
	}
}
