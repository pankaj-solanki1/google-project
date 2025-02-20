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
		if sum%2 == 0 {
			result = append(result, false)
			return result
		}
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
			sum += i
		}
	}
	return result
}

func main() {
	result := FizzBuzz()
	fmt.Println(result)
}

func TestFizzBuzz(t *testing.T) {
	t.Run("Multiples of 3, 5, and both", func(t *testing.T) {
		result := FizzBuzz()
		if result[2] != "Fizz" {
			t.Errorf("Expected Fizz at index 2, got %v", result[2])
		}
		if result[4] != "Buzz" {
			t.Errorf("Expected Buzz at index 4, got %v", result[4])
		}
		if result[14] != "FizzBuzz" {
			t.Errorf("Expected FizzBuzz at index 14, got %v", result[14])
		}
	})

	t.Run("Stops execution when sum is even", func(t *testing.T) {
		result := FizzBuzz()
		sum := 0
		expectedLength := 0
		for i := 1; i <= 100; i++ {
			if sum%2 == 0 {
				expectedLength = i - 1
				break
			}
			if i%3 == 0 && i%5 == 0 {
				//Do nothing, FizzBuzz case
			} else if i%3 == 0 {
				//Do nothing, Fizz case
			} else if i%5 == 0 {
				//Do nothing, Buzz case
			} else {
				sum += i
			}

			if i == 100 {
				expectedLength = 100
			}
		}

		if sum%2 == 0 {
			if len(result) != expectedLength {
				t.Errorf("Expected length %d when sum is even, got %d", expectedLength, len(result))
			}

			if len(result) > 0 {
				if result[len(result)-1] != false {
					t.Errorf("Expected false at the end of the slice when sum is even, got %v", result[len(result)-1])
				}
			}
		} else {
			if len(result) != 100 {
				t.Errorf("Expected length 100 when sum is always odd, got %d", len(result))
			}
		}
	})

	t.Run("First element is 1", func(t *testing.T) {
		result := FizzBuzz()
		if len(result) > 0 {
			firstElement, ok := result[0].(string)
			if !ok {
				t.Fatalf("Expected first element to be a string, got %T", result[0])
			}
			if firstElement != "1" {
				t.Errorf("Expected first element to be '1', got %s", firstElement)
			}
		}
	})
}
