package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type waitingArea struct {
	seats                         [][]rune
	numberOfRows                  int
	numberOfColumns               int
	emptySeat                     rune
	occupiedSeat                  rune
	floor                         rune
	maxAdjacentOccupiedSeats      int
	searchImmediatelyAdjacentOnly bool
}

func new(seatingRows []string, emptySeat rune, ocupiedSeat rune, floor rune, maxAdjacentOccupiedSeats int, searchImmediatelyAdjacentOnly bool) *waitingArea {
	seats := make([][]rune, len(seatingRows))
	for rowNumber, row := range seatingRows {
		newRow := make([]rune, len(row))
		for columnNumber, seat := range row {
			newRow[columnNumber] = seat
		}
		seats[rowNumber] = newRow
	}
	return &waitingArea{
		seats:                         seats,
		emptySeat:                     emptySeat,
		occupiedSeat:                  ocupiedSeat,
		floor:                         floor,
		numberOfRows:                  len(seats),
		numberOfColumns:               len(seats[0]),
		maxAdjacentOccupiedSeats:      maxAdjacentOccupiedSeats,
		searchImmediatelyAdjacentOnly: searchImmediatelyAdjacentOnly,
	}
}

func (wa *waitingArea) print() {
	for _, row := range wa.seats {
		fmt.Printf("%s\n", string(row))
	}
}

func (wa *waitingArea) numberOfOccupiedSeats() int {
	return wa.numberOfSeatsOfType(wa.occupiedSeat)
}

func (wa *waitingArea) numberOfSeatsOfType(seatType rune) int {
	numberOfSeats := 0
	for _, row := range wa.seats {
		for _, seat := range row {
			if seat == seatType {
				numberOfSeats++
			}
		}
	}
	return numberOfSeats
}

func (wa *waitingArea) simulateSeatingEvolution() {
	for {
		seatsToUpdate := wa.cloneSeats()
		numberOfUpdatedSeats := wa.simulateSeatingRound(seatsToUpdate)
		if numberOfUpdatedSeats == 0 {
			return
		}
	}
}

func (wa *waitingArea) cloneSeats() [][]rune {
	clonedSeats := make([][]rune, len(wa.seats))
	for rowNumber, row := range wa.seats {
		clonedRow := make([]rune, len(row))
		copy(clonedRow, row)
		clonedSeats[rowNumber] = clonedRow
	}
	return clonedSeats
}

func (wa *waitingArea) simulateSeatingRound(seatsToUpdate [][]rune) int {
	numberOfSeatsThatChangedState := 0
	for rowNumber, row := range seatsToUpdate {
		for columnNumber := range row {
			seatUpdated := wa.updateSeat(seatsToUpdate, rowNumber, columnNumber)
			if seatUpdated {
				numberOfSeatsThatChangedState++
			}
		}
	}
	wa.seats = seatsToUpdate
	return numberOfSeatsThatChangedState
}

func (wa *waitingArea) updateSeat(seatsToUpdate [][]rune, row int, column int) bool {
	seat := wa.seats[row][column]
	if seat == wa.floor {
		return false
	}
	adjacentOccupiedSeatCount := 0
	for _, seat := range wa.adjacentSeats(row, column) {
		if seat == wa.occupiedSeat {
			adjacentOccupiedSeatCount++
		}
	}
	if seat == wa.emptySeat && adjacentOccupiedSeatCount == 0 {
		seatsToUpdate[row][column] = wa.occupiedSeat
		return true
	}
	if seat == wa.occupiedSeat && adjacentOccupiedSeatCount >= wa.maxAdjacentOccupiedSeats {
		seatsToUpdate[row][column] = wa.emptySeat
		return true
	}
	return false
}

func (wa *waitingArea) adjacentSeats(row int, column int) []rune {
	seats := []rune{}
	vectors := []vector{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, vector := range vectors {
		seats = wa.adjacentSeatsInDirection(seats, row, column, vector.xDirection, vector.yDirection)
	}

	return seats
}

type vector struct {
	xDirection int
	yDirection int
}

func (wa *waitingArea) adjacentSeatsInDirection(seats []rune, row int, column int, xDirection int, yDirection int) []rune {
	offset := 1
	for {
		candidateRow := row + offset*xDirection
		candidateColumn := column + offset*yDirection
		rowExists := candidateRow >= 0 && candidateRow <= wa.numberOfRows-1
		columnExists := candidateColumn >= 0 && candidateColumn <= wa.numberOfColumns-1
		if rowExists && columnExists {
			notFloor := wa.seats[candidateRow][candidateColumn] != wa.floor
			if notFloor {
				seats = append(seats, wa.seats[candidateRow][candidateColumn])
				return seats
			}
		} else {
			return seats
		}
		if wa.searchImmediatelyAdjacentOnly {
			return seats
		}
		offset++
	}
}

func main() {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	seatingRows := strings.Split(strings.TrimSpace(fileString), "\n")

	wa := new(seatingRows, rune("L"[0]), rune("#"[0]), rune("."[0]), 4, true)

	wa.simulateSeatingEvolution()

	fmt.Printf("part 1: %d\n", wa.numberOfOccupiedSeats())

	wa = new(seatingRows, rune("L"[0]), rune("#"[0]), rune("."[0]), 5, false)

	wa.simulateSeatingEvolution()

	fmt.Printf("part 2: %d\n", wa.numberOfOccupiedSeats())

}
