# Practice Problems - Functions and Methods

These practice problems reinforce function and method concepts. Try implementing the solutions using proper Go idioms.

## Problem 1: String Utilities

**Description:** Implement a `StringUtils` struct with methods for common string operations:
- `Reverse()` - returns the reversed string
- `CountVowels()` - returns the count of vowels (a, e, i, o, u, case insensitive)
- `IsPalindrome()` - returns true if the string is a palindrome

**Examples:**
- "hello".Reverse() → "olleh"
- "Hello World".CountVowels() → 3
- "racecar".IsPalindrome() → true

**Requirements:**
- Methods should work on the StringUtils struct
- Handle empty strings appropriately
- Case insensitive for vowels and palindromes

### Solution

```go
package main

import (
	"fmt"
	"strings"
)

type StringUtils struct {
	text string
}

func NewStringUtils(text string) *StringUtils {
	return &StringUtils{text: text}
}

func (su *StringUtils) Reverse() string {
	runes := []rune(su.text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (su *StringUtils) CountVowels() int {
	count := 0
	lower := strings.ToLower(su.text)
	vowels := "aeiou"

	for _, char := range lower {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

func (su *StringUtils) IsPalindrome() bool {
	lower := strings.ToLower(su.text)
	clean := strings.ReplaceAll(lower, " ", "")

	runes := []rune(clean)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func main() {
	// Test the StringUtils
	utils := NewStringUtils("Hello World")

	fmt.Printf("Original: %s\n", utils.text)
	fmt.Printf("Reversed: %s\n", utils.Reverse())
	fmt.Printf("Vowel count: %d\n", utils.CountVowels())
	fmt.Printf("Is palindrome: %t\n", utils.IsPalindrome())

	// Test palindrome
	palindrome := NewStringUtils("racecar")
	fmt.Printf("'%s' is palindrome: %t\n", palindrome.text, palindrome.IsPalindrome())
}
```

**Explanation:**
1. Create a StringUtils struct with a text field
2. Implement methods using pointer receivers to modify/access the struct
3. Use runes for proper Unicode handling in Reverse()
4. Convert to lowercase for case-insensitive operations
5. Clean string (remove spaces) for palindrome checking

## Problem 2: Safe Calculator

**Description:** Create a Calculator struct with methods for basic arithmetic operations. Each method should return the result and an error. Handle division by zero and overflow cases.

**Requirements:**
- Methods: Add, Subtract, Multiply, Divide
- Return (result float64, error)
- Handle division by zero
- Use defer for logging operations

**Examples:**
- calc.Add(5, 3) → 8.0, nil
- calc.Divide(10, 0) → 0.0, error("division by zero")

### Solution

```go
package main

import (
	"errors"
	"fmt"
	"math"
)

type Calculator struct {
	operations int
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) logOperation(op string) {
	defer func() {
		c.operations++
		fmt.Printf("Operation completed. Total operations: %d\n", c.operations)
	}()
	fmt.Printf("Performing %s operation...\n", op)
}

func (c *Calculator) Add(a, b float64) (float64, error) {
	defer c.logOperation("addition")
	return a + b, nil
}

func (c *Calculator) Subtract(a, b float64) (float64, error) {
	defer c.logOperation("subtraction")
	return a - b, nil
}

func (c *Calculator) Multiply(a, b float64) (float64, error) {
	defer c.logOperation("multiplication")
	result := a * b

	// Check for overflow (simplified check)
	if math.IsInf(result, 0) {
		return 0, errors.New("overflow occurred")
	}

	return result, nil
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	defer c.logOperation("division")

	if b == 0 {
		return 0, errors.New("division by zero")
	}

	result := a / b

	// Check for overflow
	if math.IsInf(result, 0) {
		return 0, errors.New("overflow occurred")
	}

	return result, nil
}

func main() {
	calc := NewCalculator()

	// Test operations
	result, err := calc.Add(10, 5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 + 5 = %.2f\n", result)
	}

	result, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = calc.Multiply(1e308, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("1e308 * 2 = %.2f\n", result)
	}
}
```

**Explanation:**
1. Calculator struct tracks operation count
2. Each method returns (result, error) following Go conventions
3. Use defer to log operations after they complete
4. Check for division by zero and overflow conditions
5. Use math.IsInf to detect overflow