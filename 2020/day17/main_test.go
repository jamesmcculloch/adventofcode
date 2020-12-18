package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		inputFile           string
		cycles              int
		expectedActiveCount int
		fourDimensional     bool
	}{
		{
			inputFile:           "input.sample",
			cycles:              1,
			expectedActiveCount: 11,
			fourDimensional:     false,
		},
		{
			inputFile:           "input.sample",
			cycles:              2,
			expectedActiveCount: 21,
			fourDimensional:     false,
		},
		{
			inputFile:           "input.sample",
			cycles:              3,
			expectedActiveCount: 38,
			fourDimensional:     false,
		},
		{
			inputFile:           "input.sample",
			cycles:              6,
			expectedActiveCount: 112,
			fourDimensional:     false,
		},
		{
			inputFile:           "input.sample",
			cycles:              1,
			expectedActiveCount: 29,
			fourDimensional:     true,
		},
		{
			inputFile:           "input.sample",
			cycles:              2,
			expectedActiveCount: 60,
			fourDimensional:     true,
		},
		{
			inputFile:           "input.sample",
			cycles:              6,
			expectedActiveCount: 848,
			fourDimensional:     true,
		},
	}

	for _, testCase := range testCases {
		initialState, err := utils.LoadStringsFromFile(testCase.inputFile)
		if err != nil {
			panic(err)
		}

		p := new(initialState, testCase.cycles, testCase.fourDimensional)

		p.bootProcess()

		assert.Equal(t, testCase.expectedActiveCount, p.active())
	}
}
