package main

import (
	"fmt"
	"sort"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func getJoltageDifferenceCounts(adapters []int) (int, int) {
	deviceJoltage := adapters[len(adapters)-1] + 3
	adapters = append(adapters, deviceJoltage)
	oneJoltDiffs := 0
	threeJoltDiffs := 0
	chargingOutlet := adapters[0]
	currentTargetAdapter := chargingOutlet
	for i := 1; i < len(adapters); i++ {
		joltageDiff := adapters[i] - currentTargetAdapter
		switch joltageDiff {
		case 1:
			oneJoltDiffs++
		case 3:
			threeJoltDiffs++
		}
		currentTargetAdapter = adapters[i]
	}
	return oneJoltDiffs, threeJoltDiffs
}

func part1(adapters []int) int {
	oneJoltDiffs, threeJoltDiffs := getJoltageDifferenceCounts(adapters)

	return oneJoltDiffs * threeJoltDiffs
}

func totalValidConnections(adapters []int) int {
	total := 0
	memo := setUpMemo(adapters)

	maximumJoltageDifference := 3
	chargingOutlet := 0
	adapterToConnectTo := chargingOutlet
	for i := 0; i <= maximumJoltageDifference; i++ {
		joltageDiff := adapters[i] - adapterToConnectTo
		if joltageDiff <= maximumJoltageDifference {
			connections := validConnections(adapters[i:], memo)
			memo[i] = connections
			total += connections
		}
	}
	return total
}

func setUpMemo(joltages []int) []int {
	memo := make([]int, joltages[len(joltages)-1])
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return memo
}

func validConnections(adapters []int, memo []int) int {
	if len(adapters) == 1 {
		return 1
	}

	currentAdapter := adapters[0]
	if memo[currentAdapter] != -1 {
		return memo[currentAdapter]
	}

	maximumJoltageDifference := 3
	totalConnectionsFromCurrentAdapter := 0
	adapterToConnectTo := currentAdapter
	for i := 1; i <= maximumJoltageDifference; i++ {
		candidateIndex := i
		if candidateIndex > len(adapters)-1 {
			continue
		}

		joltageDiff := adapters[candidateIndex] - adapterToConnectTo
		if joltageDiff <= maximumJoltageDifference {
			totalConnectionsFromCurrentAdapter += validConnections(adapters[candidateIndex:], memo)
		}
	}

	memo[currentAdapter] = totalConnectionsFromCurrentAdapter
	return totalConnectionsFromCurrentAdapter
}

func main() {
	joltages, err := utils.LoadNumbersFromInput("input")
	if err != nil {
		panic(err)
	}
	sort.Ints(joltages)

	fmt.Printf("part 1: %d\n", part1(joltages))
	fmt.Printf("part 2: %d\n", totalValidConnections(joltages))
}
