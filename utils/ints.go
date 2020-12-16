package utils

import (
	"strconv"
	"strings"
)

// FindMaxInt finds the maximum int in a unsorted slice of ints
func FindMaxInt(ints []int) int {
	var currentMax int

	for _, candidate := range ints {
		if candidate > currentMax {
			currentMax = candidate
		}
	}
	return currentMax
}

// GetInts returns a slice of ints parsed from the input string using the supplied separater
func GetInts(input string, separater string) ([]int, error) {
	values := strings.Split(input, separater)

	numbers := make([]int, len(values))
	for index, value := range values {
		number, err := strconv.Atoi(value)
		if err != nil {
			return []int{}, err
		}
		numbers[index] = number
	}
	return numbers, nil
}
