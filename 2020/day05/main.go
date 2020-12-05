package main

import (
	"fmt"
	"sort"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func getSeatIDsFromBoardingPasses(boardingPasses []string) []int {
	seatIDS := make([]int, len(boardingPasses))

	for index, boardingPass := range boardingPasses {
		seatID := getSeatIDFromBoardingPass(boardingPass)
		seatIDS[index] = seatID
	}
	return seatIDS
}

func getSeatIDFromBoardingPass(boardingPass string) int {
	row, column := getRowAndColumnFromBoardingPass(boardingPass)
	return getSeatIDFromRowAndColumn(row, column)
}

func getRowAndColumnFromBoardingPass(boardingPass string) (int, int) {
	row := findRow(boardingPass[:len(boardingPass)-3])
	column := findColumn(boardingPass[len(boardingPass)-3:])
	return row, column
}

func findRow(rowIdentifier string) int {
	return findPlaneSection("F", "B", 0, 127, rowIdentifier)
}

func findColumn(columnIdentifier string) int {
	return findPlaneSection("L", "R", 0, 7, columnIdentifier)
}

func findPlaneSection(lower string, upper string, minValue int, maxValue int, identifier string) int {
	for _, value := range identifier {
		middle := minValue + ((maxValue - minValue) / 2)
		if string(value) == lower {
			maxValue = middle
			continue
		}
		if string(value) == upper {
			minValue = middle + 1
			continue
		}
	}
	return maxValue
}

func getSeatIDFromRowAndColumn(row, column int) int {
	return (row * 8) + column
}

func findMySeatID(seatIDs []int) int {
	sort.Ints(seatIDs)

	firstSeat := seatIDs[0]
	secondSeat := seatIDs[1]
	if firstSeat+1 != secondSeat {
		if !isInFrontRow(firstSeat+1) && !isInBackRow(firstSeat+1) {
			return firstSeat + 1
		}
	}
	for i := 2; i < len(seatIDs); i++ {
		firstSeat = secondSeat
		secondSeat = seatIDs[i]
		if firstSeat+1 != secondSeat {
			if !isInFrontRow(firstSeat+1) && !isInBackRow(firstSeat+1) {
				return firstSeat + 1
			}
		}
	}
	return 0
}

func isInFrontRow(seatID int) bool {
	if seatID <= 7 {
		return true
	}
	return false
}

func isInBackRow(seatID int) bool {
	if seatID >= 1016 {
		return true
	}
	return false
}

func main() {
	boardingPasses, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	seatIDs := getSeatIDsFromBoardingPasses(boardingPasses)

	fmt.Printf("part 1: %d\n", utils.FindMaxInt(seatIDs))
	fmt.Printf("part 2: %d\n", findMySeatID(seatIDs))
}
