# Practice Problems - Unit Testing

These practice problems reinforce testing concepts. Focus on writing comprehensive tests, table-driven tests, and proper test organization.

## Problem 1: Testing a Calculator Package

**Description:** Write comprehensive unit tests for a calculator package that includes basic arithmetic operations and error handling.

**Requirements:**
- Test all arithmetic operations (add, subtract, multiply, divide)
- Test error cases (division by zero)
- Use table-driven tests for multiple test cases
- Test edge cases (overflow, negative numbers)
- Include benchmarks for performance testing

**Test Coverage Goals:**
- All functions tested
- Error conditions tested
- Edge cases covered
- Benchmarks included

**Constraints:**
- Use testing package conventions
- Follow Go testing best practices
- Include meaningful test names and failure messages

### Solution

```go
package calculator

import (
	"errors"
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

// Divide returns the quotient of two numbers and an error if division by zero
func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Power returns a raised to the power of b
func (c Calculator) Power(a, b float64) float64 {
	return math.Pow(a, b)
}
```

```go
package calculator

import (
	"math"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", 5, -3, 2},
		{"zero addition", 0, 5, 5},
		{"decimal numbers", 1.5, 2.5, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zero subtraction", 5, 0, 5},
		{"decimal numbers", 5.5, 2.5, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 3, 4, 12},
		{"negative numbers", -3, 4, -12},
		{"zero multiplication", 0, 5, 0},
		{"decimal numbers", 1.5, 2, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := Calculator{}

	t.Run("valid divisions", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     float64
			expected float64
		}{
			{"positive numbers", 10, 2, 5},
			{"decimal result", 5, 2, 2.5},
			{"negative divisor", 10, -2, -5},
			{"both negative", -10, -2, 5},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result, err := calc.Divide(tt.a, tt.b)
				if err != nil {
					t.Errorf("Divide(%v, %v) returned unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, result, tt.expected)
				}
			})
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10, 0)
		if err == nil {
			t.Error("Divide by zero should return an error")
		}
		if err.Error() != "division by zero" {
			t.Errorf("Expected 'division by zero' error, got: %v", err)
		}
	})
}

func TestCalculator_Power(t *testing.T) {
	calc := Calculator{}

	tests := []struct {
		name     string
		base     float64
		exp      float64
		expected float64
	}{
		{"square", 3, 2, 9},
		{"cube", 2, 3, 8},
		{"zero power", 5, 0, 1},
		{"power of zero", 0, 5, 0},
		{"fractional power", 4, 0.5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Power(tt.base, tt.exp)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Power(%v, %v) = %v, want %v", tt.base, tt.exp, result, tt.expected)
			}
		})
	}
}

// Benchmarks
func BenchmarkCalculator_Add(b *testing.B) {
	calc := Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(1.5, 2.5)
	}
}

func BenchmarkCalculator_Multiply(b *testing.B) {
	calc := Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Multiply(3, 4)
	}
}

func BenchmarkCalculator_Power(b *testing.B) {
	calc := Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Power(2, 10)
	}
}
```

**Explanation:**
1. Table-driven tests for multiple test cases
2. Subtests for organizing related test cases
3. Error testing for division by zero
4. Benchmarks for performance testing
5. Clear test names and failure messages

## Problem 2: Testing String Utilities

**Description:** Write tests for string utility functions including validation, transformation, and analysis operations.

**Requirements:**
- Test string validation functions
- Test string transformation functions
- Use subtests for organization
- Include edge cases and error conditions
- Test with various Unicode characters

**Test Coverage Goals:**
- Empty strings
- Unicode characters
- Case sensitivity
- Boundary conditions

### Solution

```go
package stringsutil

import (
	"strings"
	"unicode"
)

// IsPalindrome checks if a string is a palindrome (case insensitive, ignores spaces)
func IsPalindrome(s string) bool {
	// Clean the string: remove spaces and convert to lowercase
	clean := strings.ToLower(strings.ReplaceAll(s, " ", ""))

	runes := []rune(clean)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	if s == "" {
		return 0
	}

	words := strings.Fields(s)
	return len(words)
}

// CapitalizeWords capitalizes the first letter of each word
func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

// ReverseWords reverses the order of words in a string
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
```

```go
package stringsutil

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"simple palindrome", "radar", true},
		{"mixed case palindrome", "Racecar", true},
		{"palindrome with spaces", "A man a plan a canal Panama", true},
		{"not a palindrome", "hello", false},
		{"unicode palindrome", "été", true},
		{"unicode not palindrome", "café", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world", 2},
		{"words with punctuation", "hello, world!", 2},
		{"multiple spaces", "hello   world", 2},
		{"leading/trailing spaces", "  hello world  ", 2},
		{"unicode words", "café au lait", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalizeWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
		{"multiple words", "hello world", "Hello World"},
		{"mixed case", "hELLO wORLD", "HELLO WORLD"},
		{"unicode", "café au lait", "Café Au Lait"},
		{"already capitalized", "Hello World", "Hello World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CapitalizeWords(tt.input)
			if result != tt.expected {
				t.Errorf("CapitalizeWords(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single word", "hello", "hello"},
		{"two words", "hello world", "world hello"},
		{"three words", "the quick brown", "brown quick the"},
		{"multiple spaces", "hello   world", "world hello"},
		{"unicode", "café au lait", "lait au café"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWords(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseWords(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkIsPalindrome(b *testing.B) {
	testString := "A man a plan a canal Panama"
	for i := 0; i < b.N; i++ {
		IsPalindrome(testString)
	}
}

func BenchmarkCountWords(b *testing.B) {
	testString := "The quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		CountWords(testString)
	}
}

func BenchmarkCapitalizeWords(b *testing.B) {
	testString := "the quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		CapitalizeWords(testString)
	}
}
```

**Explanation:**
1. Comprehensive table-driven tests for all functions
2. Edge cases including empty strings and Unicode
3. Clear test organization with descriptive names
4. Benchmarks for performance testing
5. Error-free functions tested for correctness