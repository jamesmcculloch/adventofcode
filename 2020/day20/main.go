package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type imageTile struct {
	id    int
	image []string

	edgeMatched            map[string]bool
	potentialAdjacentTiles []*imageTile
}

func (t *imageTile) column(i int) string {
	imageColumn := ""
	for _, s := range t.image {
		imageColumn += string(s[i])
	}
	return imageColumn
}

func (t *imageTile) edges() []string {
	return []string{t.image[0], t.column(len(t.image[0]) - 1), t.image[len(t.image)-1], t.column(0)}
}

func (t imageTile) orientations() []imageTile {
	orientations := make([]imageTile, 8)
	for r := 0; r < 8; r += 2 {
		for _, row := range t.image {
			orientations[r].image = append(orientations[r].image, reverse(row))
		}
		orientations[r].id = t.id
		orientations[r].potentialAdjacentTiles = t.potentialAdjacentTiles
		for i := range t.image {
			orientations[r+1].image = append(orientations[r+1].image, reverse(t.column(i)))
		}
		orientations[r+1].potentialAdjacentTiles = t.potentialAdjacentTiles
		orientations[r+1].id = t.id
		t = orientations[r+1]
	}
	return orientations
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
		newTile.edgeMatched = map[string]bool{}
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
	for _, edge := range tile.edges() {
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
	for _, edge := range tile.edges() {
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
				tile.potentialAdjacentTiles = append(tile.potentialAdjacentTiles, tileWithEdge)
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
		for _, edge := range tile.edges() {
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

func assembleTiles(edgeTiles []*imageTile, tiles []*imageTile) image {

	for _, edgeTile := range edgeTiles {
		for _, tile := range edgeTile.orientations() {
			image := emptyImage(tiles)
			x := 0
			y := 0
			image[x][y] = &tile
			if assembleTilesFromTile(image, x, y, &tile) {
				return image
			}
		}
	}

	return image{}
}

func assembleTilesFromTile(image image, x, y int, tileForNeighbours *imageTile) bool {
	if x == len(image)-1 && y == len(image)-1 {
		return true
	}

	if y == len(image)-1 {
		tileForNeighbours = image[x][0]
	}

	nextX := x
	nextY := y
	if y == len(image)-1 {
		nextY = 0
		nextX++
	} else {
		nextY++
	}

	for _, neighbour := range tileForNeighbours.potentialAdjacentTiles {
		for _, orientedNeighbour := range neighbour.orientations() {
			if !tileUsed(image, orientedNeighbour.id, x, y) {
				if nextX == 0 || image[nextX-1][nextY].edges()[2] == orientedNeighbour.edges()[0] {
					if nextY == 0 || image[nextX][nextY-1].edges()[1] == orientedNeighbour.edges()[3] {
						image[nextX][nextY] = &orientedNeighbour
						if assembleTilesFromTile(image, nextX, nextY, &orientedNeighbour) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func tileUsed(image image, tileID, x, y int) bool {
	for row := 0; row <= x; row++ {
		for column := 0; column <= y; column++ {
			if tileID == image[x][y].id {
				return true
			}
		}
	}
	return false
}

type image [][]*imageTile

func (i image) printIDs() {
	for _, row := range i {
		for _, tile := range row {
			if tile == nil {
				fmt.Print("nil ")
			} else {
				fmt.Printf("%d ", tile.id)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (i image) formPicture() *imageTile {
	picture := make([]string, len(i)*(len(i[0][0].image)-2))
	for i, row := range i {
		for _, tile := range row {
			for k, line := range tile.image {
				if k != 0 && k != len(tile.image)-1 {
					picture[(i*(len(tile.image)-2))+k-1] += line[1 : len(line)-1]
				}
			}
		}
	}
	return &imageTile{image: picture}
}

func emptyImage(tiles []*imageTile) image {
	length := int(math.Sqrt(float64(len(tiles))))
	image := make([][]*imageTile, length)
	for i := 0; i < length; i++ {
		imageRow := make([]*imageTile, length)
		image[i] = imageRow
	}
	return image
}

func waterRoughness(picture *imageTile, seaMonster []string) int {
	seaMonsterCount := 0
	for _, orientation := range picture.orientations() {
		for i := 0; i < len(orientation.image)-len(seaMonster); i++ {
			for j := 0; j < len(orientation.image[0])-len(seaMonster[0]); j++ {
				if foundSeaMonster(orientation.image, i, j, seaMonster) {
					seaMonsterCount++
				}
			}
		}
		if seaMonsterCount > 0 {
			break
		}
	}
	return countElements(picture.image) - (seaMonsterCount * countElements(seaMonster))
}

func foundSeaMonster(image []string, x, y int, seaMonster []string) bool {
	for i, row := range seaMonster {
		for j, element := range row {
			if element == '#' {
				if image[x+i][y+j] != '#' {
					return false
				}
			}
		}
	}
	return true
}

func countElements(image []string) int {
	count := 0
	for _, row := range image {
		for _, element := range row {
			if element == '#' {
				count++
			}
		}
	}
	return count
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

	image := assembleTiles(edgeTiles, imageTiles)
	seaMonster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	fmt.Printf("part 2: %d\n", waterRoughness(image.formPicture(), seaMonster))
}
