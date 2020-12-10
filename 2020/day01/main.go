package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func part1(numbers []int, target int) int {
	first, second, solutionFound := findTwoNumbersThatSumToTarget(numbers, target)
	if solutionFound {
		return first * second

	}
	return -1
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

func part2(numbers []int, target int) int {
	first, second, third, solutionFound := findThreeNumbersThatSumToTarget(numbers, target)
	if solutionFound {
		return first * second * third
	}
	return -1
}

func findThreeNumbersThatSumToTarget(numbers []int, target int) (int, int, int, bool) {
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
	expenseReport, err := utils.LoadNumbersFromInput("input")
	if err != nil {
		panic(err)
	}

	target := 2020
	fmt.Printf("part 1: %d\n", part1(expenseReport, target))
	fmt.Printf("part 2: %d\n", part2(expenseReport, target))
}
