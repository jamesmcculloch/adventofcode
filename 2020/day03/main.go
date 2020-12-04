package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getMapPattern(filepath string) (*geologicalMap, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return &geologicalMap{}, err
	}
	pattern := string(bytes)
	lines := strings.Split(pattern, "\r\n")

	return &geologicalMap{
		pattern: lines,
		width:   len(lines[0]),
		height:  len(lines),
	}, nil
}

type geologicalMap struct {
	pattern []string
	width   int
	height  int
}

func (m geologicalMap) isTree(x, y int) bool {
	x = (x % m.width)
	return string(m.pattern[y][x]) == "#"
}

func (m geologicalMap) treeCountForRun(xOffset int, yOffset int) int {
	currentX, currentY, treeCount := 0, 0, 0

	for {
		if currentY >= m.height-1 {
			break
		}

		currentX += xOffset
		currentY += yOffset

		if m.isTree(currentX, currentY) {
			treeCount++
		}
	}
	return treeCount
}

func main() {
	puzzleMap, err := getMapPattern("input")
	if err != nil {
		panic(err)
	}

	numberOfTreeEncountered := puzzleMap.treeCountForRun(3, 1)
	fmt.Printf("part 1: %d\r\n", numberOfTreeEncountered)

	part2Inputs := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	part2Result := 1
	for _, input := range part2Inputs {
		numberOfTreeEncountered := puzzleMap.treeCountForRun(input[0], input[1])
		fmt.Printf("dx: %d, dy: %d, tree count: %d\r\n", input[0], input[1], numberOfTreeEncountered)
		part2Result *= numberOfTreeEncountered
	}
	fmt.Printf("part 2: %d\r\n", part2Result)
}
