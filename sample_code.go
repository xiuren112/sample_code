// =============================================
// Sample Go Code
// Author: https://github.com/xiuren112
// Date: 2026-01-05
// Description: Utility functions with basic error handling and examples
// =============================================

package main

import (
	"errors"
	"fmt"
	"strings"
)

// -----------------------------
// Function 1: Add two numbers
// -----------------------------
func AddNumbers(a, b float64) float64 {
	return a + b
}

// -----------------------------
// Function 2: Count words in a string
// -----------------------------
func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// -----------------------------
// Function 3: Safe division
// -----------------------------
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// -----------------------------
// Main function with examples
// -----------------------------
func main() {
	// Example 1: AddNumbers
	sum := AddNumbers(5, 7)
	fmt.Println("AddNumbers(5,7) =", sum)

	// Example 2: CountWords
	count := CountWords("Hello world from Go")
	fmt.Println("CountWords('Hello world from Go') =", count)

	// Example 3: SafeDivide
	result, err := SafeDivide(10, 2)
	if err != nil {
		fmt.Println("SafeDivide(10,2) error:", err)
	} else {
		fmt.Println("SafeDivide(10,2) =", result)
	}

	// Example 3b: division by zero
	result, err = SafeDivide(10, 0)
	if err != nil {
		fmt.Println("SafeDivide(10,0) error:", err)
	} else {
		fmt.Println("SafeDivide(10,0) =", result)
	}
}
