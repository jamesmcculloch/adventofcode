package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tiles, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	imageTiles := load(tiles)

	mapping := edgeToTileMapping(imageTiles)

	findNumberOfMatchingEdges(imageTiles, mapping)

	edgeTiles := findEdgeTiles(imageTiles)
	for _, tile := range edgeTiles {
		tile.print()
	}

	assert.Equal(t, 4, len(edgeTiles))
	assert.Equal(t, 20899048083289, productOfEdgeTileIDs(edgeTiles))
}
