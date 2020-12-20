package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type imageTile struct {
	id    int
	image []string

	edges       []string
	edgeMatched map[string]bool
}

func (t *imageTile) setEdges() {
	t.edges = append(t.edges, t.image[0])

	t.edges = append(t.edges, t.image[len(t.image)-1])

	left := ""
	for i := 0; i < len(t.image); i++ {
		left += string(t.image[i][0])
	}
	t.edges = append(t.edges, left)

	right := ""
	for i := 0; i < len(t.image); i++ {
		right += string(t.image[i][len(t.image[0])-1])
	}
	t.edges = append(t.edges, right)

	t.edgeMatched = make(map[string]bool)
	for _, edge := range t.edges {
		t.edgeMatched[edge] = false
	}
}

func (t *imageTile) print() {
	fmt.Printf("Tile %d:\n", t.id)
	for _, row := range t.image {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()
}

func load(tiles [][]string) []*imageTile {
	imageTiles := make([]*imageTile, len(tiles))
	for index, tile := range tiles {
		newTile := &imageTile{}

		newTile.id = getID(tile[0])

		newTile.image = getImage(tile[1:])

		newTile.setEdges()

		imageTiles[index] = newTile
	}

	return imageTiles
}

func getID(tileHeader string) int {
	tileHeader = strings.TrimPrefix(tileHeader, "Tile ")
	tileHeader = strings.TrimSuffix(tileHeader, ":")

	id, err := strconv.Atoi(tileHeader)
	if err != nil {
		panic(err)
	}
	return id
}

func getImage(pixels []string) []string {
	image := make([]string, len(pixels))
	for index, row := range pixels {
		row = strings.TrimSpace(row)
		if row == "" {
			panic(pixels)
		}
		image[index] = row
	}
	return image
}

func edgeToTileMapping(tiles []*imageTile) map[string][]*imageTile {
	mapping := map[string][]*imageTile{}
	for _, tile := range tiles {
		addEdgesToMapping(mapping, tile)
	}
	return mapping
}

func addEdgesToMapping(mapping map[string][]*imageTile, tile *imageTile) {

	for _, edge := range tile.edges {
		addEdgeToMapping(mapping, edge, tile)
	}
}

func addEdgeToMapping(mapping map[string][]*imageTile, edge string, tile *imageTile) {
	if tiles, ok := mapping[edge]; !ok {
		mapping[edge] = []*imageTile{tile}
	} else {
		if !!!isIn(tiles, tile) {
			mapping[edge] = append(tiles, tile)
		}
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isIn(tiles []*imageTile, targetTile *imageTile) bool {
	for _, tile := range tiles {
		if tile == targetTile {
			return true
		}
	}
	return false
}

func findNumberOfMatchingEdges(tiles []*imageTile, edgeMapping map[string][]*imageTile) {
	for _, tile := range tiles {
		updateTilesWithMatchingEdges(tile, edgeMapping)
	}
}

func updateTilesWithMatchingEdges(tile *imageTile, edgeMapping map[string][]*imageTile) {
	for _, edge := range tile.edges {
		matched := updateTilesForMatchingEdge(tile, edge, edgeMapping, false)
		if !matched {
			updateTilesForMatchingEdge(tile, edge, edgeMapping, true)
		}
	}
}

func updateTilesForMatchingEdge(tile *imageTile, edge string, edgeMapping map[string][]*imageTile, reversed bool) bool {
	targetEdge := edge
	if reversed {
		targetEdge = reverse(edge)
	}
	if tilesWithEdge, ok := edgeMapping[targetEdge]; ok {
		for _, tileWithEdge := range tilesWithEdge {
			if tileWithEdge != tile {
				tile.edgeMatched[edge] = true
				return true
			}
		}
	}
	return false
}

func findEdgeTiles(tiles []*imageTile) []*imageTile {
	edgeTiles := []*imageTile{}

	for _, tile := range tiles {
		matchedCount := 0
		for _, edge := range tile.edges {
			if tile.edgeMatched[edge] {
				matchedCount++
			}
		}
		if matchedCount == 2 {
			edgeTiles = append(edgeTiles, tile)
		}
	}

	return edgeTiles
}

func productOfEdgeTileIDs(edges []*imageTile) int {
	product := 1
	for _, tile := range edges {
		product *= tile.id
	}
	return product
}

func main() {
	tiles, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	imageTiles := load(tiles)

	mapping := edgeToTileMapping(imageTiles)

	findNumberOfMatchingEdges(imageTiles, mapping)

	edgeTiles := findEdgeTiles(imageTiles)

	fmt.Printf("part 1: %d\n", productOfEdgeTileIDs(edgeTiles))
}
