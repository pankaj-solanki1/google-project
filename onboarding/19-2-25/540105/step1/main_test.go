package main

import (
	"fmt"
	"testing"
)

// Factorial computes the factorial of n recursively
func Factorial(n int) int {
	if n > 20 {
		return 0
	}
	if n%2 == 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	n := 5
	fmt.Println(Factorial(n)) // Expected Output: 120

	n = 6
	fmt.Println(Factorial(n)) // Expected Output: -1

	n = 21
	fmt.Println(Factorial(n)) // Expected Output: 0

	n = 4
	fmt.Println(Factorial(n)) // Expected Output: -1
}

func TestFactorial(t *testing.T) {
	// Test case 1: Positive odd number
	result := Factorial(5)
	if result != 120 {
		t.Errorf("Factorial(5) = %d; expected 120", result)
	}

	// Test case 2: Even number
	result = Factorial(6)
	if result != -1 {
		t.Errorf("Factorial(6) = %d; expected -1", result)
	}

	// Test case 3: Number greater than 20
	result = Factorial(21)
	if result != 0 {
		t.Errorf("Factorial(21) = %d; expected 0", result)
	}

	// Test case 4: Another even number
	result = Factorial(4)
	if result != -1 {
		t.Errorf("Factorial(4) = %d; expected -1", result)
	}

	// Test case 5: Zero
	result = Factorial(0)
	if result != 1 {
		t.Errorf("Factorial(0) = %d; expected 1", result)
	}

	// Test case 6: One
	result = Factorial(1)
	if result != 1 {
		t.Errorf("Factorial(1) = %d; expected 1", result)
	}

	// Test case 7: Larger odd number within limit
	result = Factorial(9)
	if result != 362880 {
		t.Errorf("Factorial(9) = %d; expected 362880", result)
	}

	// Test case 8: Edge case, n = 20
	result = Factorial(20)
	if result != -1 {
		t.Errorf("Factorial(20) = %d; expected -1", result)
	}
}
