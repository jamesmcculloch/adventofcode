package main

import (
	"fmt"
	"sort"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func firstInvalidNumber(xmas []int) (int, int) {
	for i := 25; i < len(xmas); i++ {
		valid := isNumberValid(xmas[i-25:i], xmas[i])
		if !valid {
			return i, xmas[i]
		}
	}
	return -1, -1
}

func isNumberValid(numbers []int, target int) bool {
	_, _, valid := findTwoNumbersThatSumToTarget(numbers, target)
	return valid
}

func findTwoNumbersThatSumToTarget(numbers []int, target int) (int, int, bool) {
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

func encryptionWeakness(xmas []int, firstInvalidNumber int) int {
	group := findContiguousGroupThatSumToTarget(xmas, firstInvalidNumber)
	if len(group) < 2 {
		return -1
	}
	return addSmallestAndLargest(group)
}

func findContiguousGroupThatSumToTarget(numbers []int, target int) []int {
	for index, candidateStart := range numbers {
		if candidateStart > target {
			continue
		}
		runningTotal := candidateStart
		for i := index + 1; i < len(numbers); i++ {
			runningTotal += numbers[i]
			if runningTotal == target {
				return numbers[index : i+1]
			}
			if runningTotal > target {
				break
			}
		}
	}
	return []int{}
}

func addSmallestAndLargest(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0] + numbers[len(numbers)-1]
}

func main() {
	xmas, err := utils.LoadNumbersFromInput("input")
	if err != nil {
		panic(err)
	}

	_, invalidNumber := firstInvalidNumber(xmas)
	fmt.Printf("part 1: %d\n", invalidNumber)
	fmt.Printf("part 2: %d\n", encryptionWeakness(xmas, invalidNumber))
}
