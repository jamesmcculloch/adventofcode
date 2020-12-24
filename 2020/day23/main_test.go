package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	crabCupsConfig, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	game := new(crabCupsConfig[0], false)

	game.play(100)

	assert.Equal(t, "67384529", game.resultingOrder())
}

func TestPart2(t *testing.T) {
	crabCupsConfig, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	game := new(crabCupsConfig[0], true)

	game.play(10000000)

	assert.Equal(t, 149245887792, game.part2())
}
