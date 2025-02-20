package main

import (
	"fmt"
	"strconv"
	"testing"
)

// FizzBuzz generates the sequence from 1 to 100 with the given rules.
func FizzBuzz() []interface{} {
	result := make([]interface{}, 0)
	sum := 0
	for i := 1; i <= 100; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			output = strconv.Itoa(i)
			num, _ := strconv.Atoi(output)
			sum += num
		} else {
			sum += i // Add i to the sum even when Fizz/Buzz/FizzBuzz is output
		}

		if sum%2 == 0 {
			result = append(result, false)
			return result
		}

		result = append(result, output)
	}
	return result
}

func main() {
	result := FizzBuzz()
	fmt.Println(result)
}

// --- Unit Tests ---

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz()

	// Test the first few elements
	if result[0] != "1" {
		t.Errorf("Expected '1', got '%v'", result[0])
	}
	if result[1] != "2" {
		t.Errorf("Expected '2', got '%v'", result[1])
	}
	if result[2] != "Fizz" {
		t.Errorf("Expected 'Fizz', got '%v'", result[2])
	}
	if result[4] != "Buzz" {
		t.Errorf("Expected 'Buzz', got '%v'", result[4])
	}
	if result[14] != "FizzBuzz" {
		t.Errorf("Expected 'FizzBuzz', got '%v'", result[14])
	}

	//Test for the early exit condition
	sum := 0
	earlyExitIndex := -1
	expectedValues := []interface{}{}
	for i := 1; i <= 100; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			output = strconv.Itoa(i)
			num, _ := strconv.Atoi(output)
			sum += num
		} else {
			sum += i
		}

		if sum%2 == 0 {
			earlyExitIndex = i
			expectedValues = append(expectedValues, false)
			break
		}

		expectedValues = append(expectedValues, output)
	}

	// Check the length and the last element for early termination
	if len(result) != len(expectedValues) {
		t.Errorf("Expected length %d, got %d", len(expectedValues), len(result))
	}

	if earlyExitIndex != -1 {
		if result[len(result)-1] != false {
			t.Errorf("Expected 'false' at the end due to early exit, got '%v'", result[len(result)-1])
		}
	}

	//Check all elements before early termination point
	for i := 0; i < len(result); i++ {
		if result[i] != expectedValues[i] && result[i] != false {
			t.Errorf("Mismatch at index %d: expected '%v', got '%v'", i, expectedValues[i], result[i])
		}
	}

}
