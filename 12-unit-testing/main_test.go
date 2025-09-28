package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}

	result := calc.Add(2, 3)
	expected := 5.0

	if result != expected {
		t.Errorf("Add(2, 3) = %f; want %f", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	calc := Calculator{}

	result := calc.Subtract(10, 4)
	expected := 6.0

	if result != expected {
		t.Errorf("Subtract(10, 4) = %f; want %f", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	calc := Calculator{}

	result := calc.Multiply(6, 7)
	expected := 42.0

	if result != expected {
		t.Errorf("Multiply(6, 7) = %f; want %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	calc := Calculator{}

	// Test normal division
	result, err := calc.Divide(15, 3)
	if err != nil {
		t.Errorf("Divide(15, 3) returned error: %v", err)
	}
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(15, 3) = %f; want %f", result, expected)
	}

	// Test division by zero
	_, err = calc.Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}
}

func TestIsEven(t *testing.T) {
	calc := Calculator{}

	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, tc := range testCases {
		result := calc.IsEven(tc.input)
		if result != tc.expected {
			t.Errorf("IsEven(%d) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	calc := Calculator{}

	testCases := []struct {
		input    int
		expected int
		hasError bool
	}{
		{0, 1, false},
		{1, 1, false},
		{5, 120, false},
		{3, 6, false},
		{-1, 0, true},
	}

	for _, tc := range testCases {
		result, err := calc.Factorial(tc.input)

		if tc.hasError {
			if err == nil {
				t.Errorf("Factorial(%d) should return error", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("Factorial(%d) returned unexpected error: %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tc.input, result, tc.expected)
			}
		}
	}
}

func TestIsPrime(t *testing.T) {
	calc := Calculator{}

	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{9, false},
		{13, true},
		{15, false},
		{17, true},
		{1, false},
		{0, false},
		{-1, false},
	}

	for _, tc := range testCases {
		result := calc.IsPrime(tc.input)
		if result != tc.expected {
			t.Errorf("IsPrime(%d) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}

func TestMax(t *testing.T) {
	calc := Calculator{}

	testCases := []struct {
		a, b     float64
		expected float64
	}{
		{10, 20, 20},
		{20, 10, 20},
		{5.5, 5.5, 5.5},
		{-1, 1, 1},
		{-5, -10, -5},
	}

	for _, tc := range testCases {
		result := calc.Max(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Max(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestMin(t *testing.T) {
	calc := Calculator{}

	testCases := []struct {
		a, b     float64
		expected float64
	}{
		{10, 20, 10},
		{20, 10, 10},
		{5.5, 5.5, 5.5},
		{-1, 1, -1},
		{-5, -10, -10},
	}

	for _, tc := range testCases {
		result := calc.Min(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Min(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

// TestArithmeticOperations demonstrates subtests
func TestArithmeticOperations(t *testing.T) {
	calc := Calculator{}

	t.Run("Addition", func(t *testing.T) {
		testCases := []struct {
			name     string
			a, b     float64
			expected float64
		}{
			{"positive numbers", 2, 3, 5},
			{"negative numbers", -2, -3, -5},
			{"mixed numbers", 5, -3, 2},
			{"zero addition", 0, 5, 5},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := calc.Add(tc.a, tc.b)
				if result != tc.expected {
					t.Errorf("Add(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
				}
			})
		}
	})

	t.Run("Subtraction", func(t *testing.T) {
		testCases := []struct {
			name     string
			a, b     float64
			expected float64
		}{
			{"positive numbers", 10, 4, 6},
			{"negative numbers", -2, -3, 1},
			{"mixed numbers", 5, -3, 8},
			{"zero subtraction", 5, 0, 5},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := calc.Subtract(tc.a, tc.b)
				if result != tc.expected {
					t.Errorf("Subtract(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
				}
			})
		}
	})

	t.Run("Multiplication", func(t *testing.T) {
		testCases := []struct {
			name     string
			a, b     float64
			expected float64
		}{
			{"positive numbers", 6, 7, 42},
			{"negative numbers", -2, 3, -6},
			{"zero multiplication", 0, 5, 0},
			{"fractional numbers", 1.5, 2, 3},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := calc.Multiply(tc.a, tc.b)
				if result != tc.expected {
					t.Errorf("Multiply(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
				}
			})
		}
	})
}

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	calc := Calculator{}

	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

// BenchmarkFactorial benchmarks the Factorial function
func BenchmarkFactorial(b *testing.B) {
	calc := Calculator{}

	for i := 0; i < b.N; i++ {
		calc.Factorial(10)
	}
}

// BenchmarkIsPrime benchmarks the IsPrime function
func BenchmarkIsPrime(b *testing.B) {
	calc := Calculator{}

	for i := 0; i < b.N; i++ {
		calc.IsPrime(97) // Use a prime number
	}
}
