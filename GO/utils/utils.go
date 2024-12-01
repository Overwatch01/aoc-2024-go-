package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// CountOccurrences counts the number of times an item appears in a slice
func CountOccurrences[T comparable](slice []T, target T) int {
	count := 0
	for _, item := range slice {
		if item == target {
			count++
		}
	}
	return count
}

func ReadInput() ([]string, error) {
	input := make([]string, 0)
	file, err := os.Open("../../input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return input, nil
}
