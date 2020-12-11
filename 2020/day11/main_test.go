package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateSeatingRound(t *testing.T) {
	testCases := []struct {
		seatRows                      []string
		expectedOccupiedSeats         int
		searchImmediatelyAdjacentOnly bool
		maxAdjacentOccupiedSeats      int
	}{
		{
			seatRows: []string{
				"L.LL.LL.LL",
				"LLLLLLL.LL",
				"L.L.L..L..",
				"LLLL.LL.LL",
				"L.LL.LL.LL",
				"L.LLLLL.LL",
				"..L.L.....",
				"LLLLLLLLLL",
				"L.LLLLLL.L",
				"L.LLLLL.LL",
			},
			expectedOccupiedSeats:         37,
			searchImmediatelyAdjacentOnly: true,
			maxAdjacentOccupiedSeats:      4,
		},
		{
			seatRows: []string{
				"L.LL.LL.LL",
				"LLLLLLL.LL",
				"L.L.L..L..",
				"LLLL.LL.LL",
				"L.LL.LL.LL",
				"L.LLLLL.LL",
				"..L.L.....",
				"LLLLLLLLLL",
				"L.LLLLLL.L",
				"L.LLLLL.LL",
			},
			expectedOccupiedSeats:         26,
			searchImmediatelyAdjacentOnly: false,
			maxAdjacentOccupiedSeats:      5,
		},
	}

	for _, testCase := range testCases {
		testObject := new(
			testCase.seatRows,
			rune("L"[0]),
			rune("#"[0]),
			rune("."[0]),
			testCase.maxAdjacentOccupiedSeats,
			testCase.searchImmediatelyAdjacentOnly,
		)

		testObject.simulateSeatingEvolution()

		assert.Equal(t, testCase.expectedOccupiedSeats, testObject.numberOfOccupiedSeats())
	}
}
