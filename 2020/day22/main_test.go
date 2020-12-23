package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	game, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load test input", err)
	}

	player1, err := getDeckOfCards(game[0][1:])
	if err != nil {
		assert.Fail(t, "failed to load player 1", err)
	}
	player2, err := getDeckOfCards(game[1][1:])
	if err != nil {
		assert.Fail(t, "failed to load player 2", err)
	}

	winner := playCombat(player1, player2)

	assert.Equal(t, 306, determineWinnersScore(winner))
}

func TestPart2(t *testing.T) {
	game, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load test input", err)
	}

	player1, err := getDeckOfCards(game[0][1:])
	if err != nil {
		assert.Fail(t, "failed to load player 1", err)
	}
	player2, err := getDeckOfCards(game[1][1:])
	if err != nil {
		assert.Fail(t, "failed to load player 2", err)
	}

	winner, _ := playRecursiveCombat(player1, player2)

	assert.Equal(t, 291, determineWinnersScore(winner))
}
