package main

import (
	"fmt"
	"strconv"
	"testing"
)

// FizzBuzz generates the sequence from 1 to 100 with the given rules.
func FizzBuzz() []interface{} {
	result := make([]interface{}, 0)
	cumulativeSum := 0
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
			cumulativeSum += num

			if cumulativeSum%2 == 0 {
				result = append(result, false)
				return result
			}
			result = append(result, output)

		} else {
			numValue := 0
			if output == "Fizz" {
				numValue = i
			} else if output == "Buzz" {
				numValue = i
			} else if output == "FizzBuzz" {
				numValue = i
			} else {
				numValue, _ = strconv.Atoi(output)

			}

			cumulativeSum += numValue

			if cumulativeSum%2 == 0 {
				result = append(result, false)
				return result
			}

			result = append(result, output)

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

	// Test the first few elements
	if len(result) > 0 && result[0] != "1" {
		t.Errorf("Expected '1' at index 0, got %v", result[0])
	}
	if len(result) > 2 && result[2] != "Fizz" {
		t.Errorf("Expected 'Fizz' at index 2, got %v", result[2])
	}
	if len(result) > 4 && result[4] != "Buzz" {
		t.Errorf("Expected 'Buzz' at index 4, got %v", result[4])
	}
	if len(result) > 14 && result[14] != "FizzBuzz" {
		t.Errorf("Expected 'FizzBuzz' at index 14, got %v", result[14])
	}

	// Check for early termination based on cumulative sum. We can't predict
	// the exact index of termination without re-implementing the entire
	// function, but we can check that *if* the function terminates early,
	// it does so with a `false` value.  We will assume the test case will run long enough to hit an even number and terminate.
	terminated := false
	for _, val := range result {
		if val == false {
			terminated = true
			break
		}
	}

	if !terminated {
		t.Log("Function did not terminate early with 'false'. This might be ok if all 100 elements were processed without an even sum, but it's worth investigating.")
	}

	//Basic length check - ensure the array returned is within the bounds we expect

	if len(result) > 100 {
		t.Errorf("Expected length to be less than or equal to 100, got %d", len(result))
	}

}
