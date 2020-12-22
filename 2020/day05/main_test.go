package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeatIDFromBoardingPass(t *testing.T) {
	testCases := []struct {
		boardingPass   string
		expectedSeatID int
	}{
		{
			boardingPass:   "BFFFBBFRRR",
			expectedSeatID: 567,
		},
		{
			boardingPass:   "FFFBBBFRRR",
			expectedSeatID: 119,
		},
		{
			boardingPass:   "BBFFBBFRLL",
			expectedSeatID: 820,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.boardingPass, func(t *testing.T) {
			testOuput := getSeatIDFromBoardingPass(testCase.boardingPass)

			assert.Equal(t, testCase.expectedSeatID, testOuput)
		})
	}
}
