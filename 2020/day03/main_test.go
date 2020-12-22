package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	puzzleMap, err := getMapPattern("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	numberOfTreeEncountered := puzzleMap.treeCountForRun(3, 1)
	assert.Equal(t, 7, numberOfTreeEncountered)
}

func TestPart2(t *testing.T) {
	puzzleMap, err := getMapPattern("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	part2Inputs := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	part2Result := 1
	for _, input := range part2Inputs {
		numberOfTreeEncountered := puzzleMap.treeCountForRun(input[0], input[1])
		part2Result *= numberOfTreeEncountered
	}
	assert.Equal(t, 336, part2Result)
}
