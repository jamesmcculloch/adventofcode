package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func getEarliestDeparture(timestampNote string) (int, error) {
	timestamp, err := strconv.Atoi(timestampNote)
	if err != nil {
		return -1, err
	}
	return timestamp, nil
}

const outOfService string = "x"

func getBusIDs(busIDsNote string) ([]int, error) {
	busIDs := []int{}
	ids := strings.Split(busIDsNote, ",")
	for _, id := range ids {
		if id == outOfService {
			continue
		}
		busID, err := strconv.Atoi(id)
		if err != nil {
			return []int{}, err
		}
		busIDs = append(busIDs, busID)
	}
	return busIDs, nil
}

func findEarliestBus(earliestDeparture int, busIDs []int) (int, int) {
	minWaitingTime := int(^uint(0) >> 1)
	earliestBus := -1

	for _, bus := range busIDs {
		waitingTime := waitingTimeForBus(bus, earliestDeparture)
		if waitingTime < minWaitingTime {
			minWaitingTime = waitingTime
			earliestBus = bus
		}
	}
	return earliestBus, minWaitingTime
}

func waitingTimeForBus(busID int, earliestDeparture int) int {
	if earliestDeparture%busID == 0 {
		return 0
	}
	return busID - (earliestDeparture % busID)
}

func part1(busID int, waitingTime int) int {
	return busID * waitingTime
}

func getBusIDsAndOffsets(busIDsNote string) (map[int]int, error) {
	busIDs := make(map[int]int)
	ids := strings.Split(busIDsNote, ",")
	offset := 0
	for _, id := range ids {
		if id != outOfService {
			busID, err := strconv.Atoi(id)
			if err != nil {
				return map[int]int{}, err
			}
			busIDs[busID] = offset
		}
		offset++
	}
	return busIDs, nil
}

func part2(busIDsAndOffsets map[int]int, busIDs []int) int {
	// find smallest t such that for all buses this equation holds (t + offset) % busID == 0
	timestamp := 0
	multiplier := 1

	for busID, offset := range busIDsAndOffsets {
		for (timestamp+offset)%busID != 0 { // solve each equation in turn

			timestamp += multiplier
		}
		multiplier *= busID // solve the next equation whilst making sure the multiplier factor that you add to the timestamp still solves the previous equations
	}

	return timestamp
}

func main() {
	notes, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	earliestDeparture, err := getEarliestDeparture(notes[0])
	if err != nil {
		panic(err)
	}

	busIDs, err := getBusIDs(notes[1])
	if err != nil {
		panic(err)
	}

	earliestBus, waitingTime := findEarliestBus(earliestDeparture, busIDs)

	fmt.Printf("part 1: %d\n", part1(earliestBus, waitingTime))

	busIDsAndOffsets, err := getBusIDsAndOffsets(notes[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 2: %d\n", part2(busIDsAndOffsets, busIDs))
}
