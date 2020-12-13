package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEarliestBus(t *testing.T) {
	testCases := []struct {
		timestampNote       string
		busIDsNote          string
		expectedBusID       int
		expectedWaitingTime int
	}{
		{
			timestampNote:       "939",
			busIDsNote:          "7,13,x,x,59,x,31,19",
			expectedBusID:       59,
			expectedWaitingTime: 5,
		},
	}

	for _, testCase := range testCases {
		earliestDeparture, err := getEarliestDeparture(testCase.timestampNote)
		if err != nil {
			panic(err)
		}
		busIDs, err := getBusIDs(testCase.busIDsNote)
		if err != nil {
			panic(err)
		}

		earliestBus, waitingTime := findEarliestBus(earliestDeparture, busIDs)

		assert.Equal(t, testCase.expectedBusID, earliestBus)
		assert.Equal(t, testCase.expectedWaitingTime, waitingTime)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		busIDsNote   string
		expectedTime int
	}{
		{
			busIDsNote:   "7,13,x,x,59,x,31,19",
			expectedTime: 1068781,
		},
	}

	for _, testCase := range testCases {
		busIDs, err := getBusIDs(testCase.busIDsNote)
		if err != nil {
			panic(err)
		}

		busIDsAndOffsets, err := getBusIDsAndOffsets(testCase.busIDsNote)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, testCase.expectedTime, part2(busIDsAndOffsets, busIDs))
	}
}
