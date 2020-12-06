package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadGroups(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	blocks := [][]string{}
	currentBlock := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			blocks = append(blocks, currentBlock)
			currentBlock = []string{}
			continue
		}

		currentBlock = append(currentBlock, line)
	}
	blocks = append(blocks, currentBlock)
	return blocks, nil
}

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
	groups, err := loadGroups("input")
	if err != nil {
		panic(err)
	}

	part1 := totalDistinctAnswerForGroups(groups)
	part2 := totalSharedAnswerForGroups(groups)

	fmt.Printf("part 1: %d\n", part1)
	fmt.Printf("part 2: %d\n", part2)
}
