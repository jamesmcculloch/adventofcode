package main

import (
	"sort"
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	joltages, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}
	sort.Ints(joltages)

	assert.Equal(t, 220, part1(joltages))
}

func TestValidConnections(t *testing.T) {
	testCases := []struct {
		joltages            []int
		expectedConnections int
	}{
		{
			joltages: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			expectedConnections: 8,
		},
		{
			joltages: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			expectedConnections: 19208,
		},
	}

	for _, testCase := range testCases {
		sort.Ints(testCase.joltages)
		connections := totalValidConnections(testCase.joltages)

		assert.Equal(t, testCase.expectedConnections, connections)
	}
}
