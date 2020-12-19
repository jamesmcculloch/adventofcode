package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	transmission, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	rules := loadRules(transmission[0])

	assert.Equal(t, 2, messagesThatMatchRules(transmission[1], rules[0]))
}

func TestPart2Sample(t *testing.T) {
	transmission, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input2.sample")
	if err != nil {
		panic(err)
	}

	rules := loadRules(transmission[0])

	assert.Equal(t, 12, messagesThatMatchRules(transmission[1], rules[0]))
}
