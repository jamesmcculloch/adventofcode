package main

import (
	"fmt"
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

	assert.Equal(t, 4, len(edgeTiles))
	assert.Equal(t, 20899048083289, productOfEdgeTileIDs(edgeTiles))
}

func TestWaterRoughness(t *testing.T) {
	tiles, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	imageTiles := load(tiles)

	mapping := edgeToTileMapping(imageTiles)

	findNumberOfMatchingEdges(imageTiles, mapping)

	edgeTiles := findEdgeTiles(imageTiles)

	image := assembleTiles(edgeTiles, imageTiles)
	seaMonster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	picture := image.formPicture()
	fmt.Printf("picture: %v\n", picture)
	assert.Equal(t, 273, waterRoughness(picture, seaMonster))
}
