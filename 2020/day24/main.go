package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

var offsets = map[string]offset{
	"ne": {1, 2},
	"e":  {2, 0},
	"se": {1, -2},
	"sw": {-1, -2},
	"w":  {-2, 0},
	"nw": {-1, 2},
}

type offset struct {
	x int
	y int
}

func blackTiles(tilesToFlip []string) (map[string]int, int) {
	tiles := map[string]int{}

	for _, tile := range tilesToFlip {
		flipTile(tile, tiles)
	}

	return tiles, countBlackTiles(tiles)
}

func flipTile(tile string, tiles map[string]int) {
	tileLocation := locateTile(tile)

	if _, ok := tiles[tileLocation]; !ok {
		tiles[tileLocation] = 1
	} else {
		tiles[tileLocation]++
	}
}

func locateTile(tile string) string {
	x := 0
	y := 0
	index := 0

	for index < len(tile) {
		switch rune(tile[index]) {
		case 'e':
			x += offsets["e"].x
			y += offsets["e"].y
			index++
			continue
		case 'w':
			x += offsets["w"].x
			y += offsets["w"].y
			index++
			continue
		case 'n':
			nextIndex := index + 1
			switch rune(tile[nextIndex]) {
			case 'e':
				x += offsets["ne"].x
				y += offsets["ne"].y
				index += 2
				continue
			case 'w':
				x += offsets["nw"].x
				y += offsets["nw"].y
				index += 2
				continue
			default:
				x += offsets["n"].x
				y += offsets["n"].y
				index++
				continue
			}
		case 's':
			nextIndex := index + 1
			switch rune(tile[nextIndex]) {
			case 'e':
				x += offsets["se"].x
				y += offsets["se"].y
				index += 2
				continue
			case 'w':
				x += offsets["sw"].x
				y += offsets["sw"].y
				index += 2
				continue
			default:
				x += offsets["s"].x
				y += offsets["s"].y
				index++
				continue
			}
		}
	}

	return fmt.Sprintf("%d,%d", x, y)
}

func countBlackTiles(tiles map[string]int) int {
	blackTiles := 0

	for _, flipCount := range tiles {
		if flipCount%2 != 0 {
			blackTiles++
		}
	}

	return blackTiles
}

func evolveTiles(tiles map[string]int, days int) int {
	minX, maxX, minY, maxY := findPatternEdges(tiles)

	for day := 1; day <= days; day++ {
		tilesToUpdate := []string{}
		for x := minX - 2; x <= maxX+2; x++ {
			for y := minY - 2; y <= maxY+2; y++ {
				tile := fmt.Sprintf("%d,%d", x, y)
				neighbours := adjacentTiles(x, y)
				blackNeighbourCount := blackNeighbours(neighbours, tiles)

				if isBlack(tile, tiles) {
					if blackNeighbourCount == 0 || blackNeighbourCount > 2 {
						tilesToUpdate = append(tilesToUpdate, tile)
					}
				} else {
					if blackNeighbourCount == 2 {
						tilesToUpdate = append(tilesToUpdate, tile)
					}
				}
			}
		}
		updateTiles(tilesToUpdate, tiles)
		minX, maxX, minY, maxY = findPatternEdges(tiles)
	}

	return countBlackTiles(tiles)
}

func adjacentTiles(x, y int) []string {
	tiles := []string{}

	for _, offset := range offsets {
		neighbourX := x + offset.x
		neighbourY := y + offset.y

		tiles = append(tiles, fmt.Sprintf("%d,%d", neighbourX, neighbourY))
	}

	return tiles
}

func findPatternEdges(tiles map[string]int) (int, int, int, int) {
	minX := utils.MaxInt
	maxX := utils.MinInt
	minY := utils.MaxInt
	maxY := utils.MinInt
	for tile, count := range tiles {
		if count%2 != 0 {
			var x int
			var y int
			fmt.Sscanf(tile, "%d,%d", &x, &y)
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	return minX, maxX, minY, maxY
}

func blackNeighbours(neighbours []string, tiles map[string]int) int {
	count := 0

	for _, neighbour := range neighbours {
		if isBlack(neighbour, tiles) {
			count++
		}
	}

	return count
}

func isBlack(tile string, tiles map[string]int) bool {
	if count, ok := tiles[tile]; ok {
		if count%2 != 0 {
			return true
		}
	}
	return false
}

func updateTiles(tilesToUpdate []string, tiles map[string]int) {
	for _, tile := range tilesToUpdate {
		if _, ok := tiles[tile]; !ok {
			tiles[tile] = 1
		} else {
			tiles[tile]++
		}
	}
}

func main() {
	tiles, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	flippedTiles, blackTileCount := blackTiles(tiles)
	fmt.Printf("part 1: %d\n", blackTileCount)

	fmt.Printf("part 2: %d\n", evolveTiles(flippedTiles, 100))
}
