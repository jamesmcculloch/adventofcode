package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		startingNumbers          []int
		turns                    int
		expectedLastNumberSpoken int
	}{
		{
			startingNumbers:          []int{1, 3, 2},
			turns:                    2020,
			expectedLastNumberSpoken: 1,
		},
		{
			startingNumbers:          []int{2, 1, 3},
			turns:                    2020,
			expectedLastNumberSpoken: 10,
		},
		{
			startingNumbers:          []int{1, 2, 3},
			turns:                    2020,
			expectedLastNumberSpoken: 27,
		},
		{
			startingNumbers:          []int{2, 3, 1},
			turns:                    2020,
			expectedLastNumberSpoken: 78,
		},
		{
			startingNumbers:          []int{3, 2, 1},
			turns:                    2020,
			expectedLastNumberSpoken: 438,
		},
		{
			startingNumbers:          []int{3, 1, 2},
			turns:                    2020,
			expectedLastNumberSpoken: 1836,
		},
		{
			startingNumbers:          []int{0, 3, 6},
			turns:                    2020,
			expectedLastNumberSpoken: 436,
		},
		{
			startingNumbers:          []int{0, 3, 6},
			turns:                    10,
			expectedLastNumberSpoken: 0,
		},
	}

	for _, testCase := range testCases {
		testObject := &memoryGame{
			startingNumbers: testCase.startingNumbers,
			memory:          make(map[int][]int),
		}

		lastNumberSpoken := testObject.play(testCase.turns)

		assert.Equal(t, testCase.expectedLastNumberSpoken, lastNumberSpoken)
	}
}
