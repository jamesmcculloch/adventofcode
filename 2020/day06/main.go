package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func totalDistinctAnswerForGroups(groups [][]string) int {
	total := 0
	for _, group := range groups {
		total += totalDistinctAnswerForGroup(group)
	}
	return total
}

func totalDistinctAnswerForGroup(group []string) int {
	answers := make(map[string]int)
	total := 0
	for _, person := range group {
		for _, answer := range person {
			if _, ok := answers[string(answer)]; !ok {
				total++
				answers[string(answer)] = 1
			} else {
				answers[string(answer)]++
			}
		}
	}

	return total
}

func totalSharedAnswerForGroups(groups [][]string) int {
	total := 0
	for _, group := range groups {
		total += totalSharedAnswerForGroup(group)
	}
	return total
}

func totalSharedAnswerForGroup(group []string) int {
	numPeopleInGroup := len(group)
	answers := make(map[string]int)
	for _, person := range group {
		for _, answer := range person {
			if _, ok := answers[string(answer)]; !ok {
				answers[string(answer)] = 1
			} else {
				answers[string(answer)]++
			}
		}
	}
	total := 0
	for _, count := range answers {
		if count == numPeopleInGroup {
			total++
		}
	}
	return total
}

func main() {
	groups, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	part1 := totalDistinctAnswerForGroups(groups)
	part2 := totalSharedAnswerForGroups(groups)

	fmt.Printf("part 1: %d\n", part1)
	fmt.Printf("part 2: %d\n", part2)
}
