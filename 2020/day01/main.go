package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func getNumbersFromInput(inputFilePath string) ([]int, error) {
	bytes, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return []int{}, err
	}
	lines := strings.Split(string(bytes), "\r\n")
	numbers := make([]int, len(lines))
	for index, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			return []int{}, err
		}
		numbers[index] = number
	}
	return numbers, nil
}

func twoSum(numbers []int, target int) (int, int, bool) {
	for i := 0; i < len(numbers); i++ {
		firstCandidate := numbers[i]
		if firstCandidate >= target {
			return 0, 0, false
		}
		targetSecondCandidate := target - firstCandidate
		found := searchInt(numbers, targetSecondCandidate)
		if found {
			return firstCandidate, targetSecondCandidate, true
		}
	}
	return 0, 0, false
}

func searchInt(numbers []int, target int) bool {
	index := sort.SearchInts(numbers, target)
	if index < len(numbers) && numbers[index] == target {
		return true
	}
	return false
}

func twoSumWithMap(numbers []int, target int) (int, int, bool) {
	candidates := make(map[int]bool)
	for _, number := range numbers {
		candidates[number] = true
	}

	for _, number := range numbers {
		remainder := target - number
		ok := candidates[remainder]
		if ok {
			return number, remainder, true
		}
	}
	return 0, 0, false
}

func threeSum(numbers []int, target int) (int, int, int, bool) {
	for i := 0; i < len(numbers); i++ {
		firstCandidate := numbers[i]
		if firstCandidate >= target {
			return 0, 0, 0, false
		}
		subTarget := target - firstCandidate
		if subTarget < 0 || subTarget < firstCandidate {
			return 0, 0, 0, false
		}
		second, third, solutionFound := twoSum(numbers[i:], subTarget)
		if solutionFound {
			return firstCandidate, second, third, true
		}
	}
	return 0, 0, 0, false
}

func threeSumWithMap(numbers []int, target int) (int, int, int, bool) {
	candidates := make(map[int]bool)
	for _, number := range numbers {
		candidates[number] = true
	}

	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			first := numbers[i]
			second := numbers[j]
			third := target - first - second
			ok := candidates[third]
			if ok {
				return first, second, third, true
			}
		}
	}
	return 0, 0, 0, false
}

func main() {
	numbers, err := getNumbersFromInput("input")
	if err != nil {
		fmt.Println("Error loading numbers from input file", err)
	}
	sort.Ints(numbers)

	target := 2020
	fmt.Printf("Searching for 2 numbers that add to: %d\n\r", target)
	first, second, solutionFound := twoSum(numbers, target)
	if solutionFound {
		fmt.Printf("found pair that add to 2020: %d, %d\n\r", first, second)
		solution := first * second
		fmt.Printf("solution 1: %d\n\r", solution)
	}

	fmt.Printf("Searching for 2 numbers that add to: %d using a map\n\r", target)
	first, second, solutionFound = twoSumWithMap(numbers, target)
	if solutionFound {
		fmt.Printf("found pair that add to 2020: %d, %d\n\r", first, second)
		solution := first * second
		fmt.Printf("solution 1: %d\n\r", solution)
	}

	fmt.Printf("Searching for 3 numbers that add to: %d\n\r", target)
	first, second, third, solutionFound := threeSum(numbers, target)
	if solutionFound {
		fmt.Printf("found 3 that add to 2020: %d, %d, %d\n\r", first, second, third)
		solution := first * second * third
		fmt.Printf("solution 2: %d\n\r", solution)
	}

	fmt.Printf("Searching for 3 numbers that add to: %d with map\n\r", target)
	first, second, third, solutionFound = threeSumWithMap(numbers, target)
	if solutionFound {
		fmt.Printf("found 3 that add to 2020: %d, %d, %d\n\r", first, second, third)
		solution := first * second * third
		fmt.Printf("solution 2: %d\n\r", solution)
	}
}
