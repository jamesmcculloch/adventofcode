package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestBlackTiles(t *testing.T) {
	tiles, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}
	_, blackTileCount := blackTiles(tiles)

	assert.Equal(t, 10, blackTileCount)
}

func TestEvolveTiles(t *testing.T) {
	tiles, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}
	flippedTiles, _ := blackTiles(tiles)

	assert.Equal(t, 2208, evolveTiles(flippedTiles, 100))
}
