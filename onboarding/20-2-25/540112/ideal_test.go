package main

import (
	"fmt"
	"strconv"
	"testing"
)

// FizzBuzz generates the sequence from 1 to 100 with the given rules.
func FizzBuzz() []interface{} {
	result := make([]interface{}, 0, 100)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func main() {
	result := FizzBuzz()
	fmt.Println(result)
}

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz()
	expected := []interface{}{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}

	// Check first few elements
	for i, val := range expected {
		if i >= len(result) {
			t.Errorf("Expected %v at index %d, but got nothing", val, i)
			continue
		}
		if result[i] != val {
			t.Errorf("Expected %v at index %d, got %v", val, i, result[i])
		}
	}

	// Ensure the length is exactly 100
	if len(result) != 100 {
		t.Errorf("Expected length 100, got %d", len(result))
	}
}
