package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamesmcculloch/adventofcode/utils"
)

func TestTotalDistinctAnswerForGroups(t *testing.T) {
	groups, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	count := totalDistinctAnswerForGroups(groups)

	assert.Equal(t, 11, count)
}

func TestTotalSharedAnswerForGroups(t *testing.T) {
	groups, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	count := totalSharedAnswerForGroups(groups)

	assert.Equal(t, 6, count)
}
