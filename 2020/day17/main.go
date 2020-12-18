package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type pocketDimension struct {
	grid            fourDimensionalGrid
	cycles          int
	fourDimensional bool
	debug           bool
}

func (p *pocketDimension) active() int {
	return p.grid.active()
}

func (p *pocketDimension) bootProcess() {
	for i := 1; i <= p.cycles; i++ {
		gridToUpdate := copyGrid(p.grid)
		p.cycle(gridToUpdate)
		p.grid = gridToUpdate
		if p.debug {
			p.print()
		}
	}
}

func (p *pocketDimension) print() {
	p.grid.print()
}

func (p *pocketDimension) cycle(gridToUpdate fourDimensionalGrid) {
	if p.fourDimensional {
		for w := 1; w < len(p.grid)-1; w++ {
			for z := 1; z < len(p.grid[w])-1; z++ {
				for y := 1; y < len(p.grid[w][z])-1; y++ {
					for x := 1; x < len(p.grid[w][z][y])-1; x++ {
						p.update(w, z, y, x, gridToUpdate)
					}
				}
			}
		}
	} else {
		for z := 1; z < len(p.grid[0])-1; z++ {
			for y := 1; y < len(p.grid[0][0])-1; y++ {
				for x := 1; x < len(p.grid[0][0][0])-1; x++ {
					p.update(0, z, y, x, gridToUpdate)
				}
			}
		}
	}
}

func (p *pocketDimension) update(w, z, y, x int, gridToUpdate fourDimensionalGrid) {
	neighbours := p.getNeighbours(w, z, y, x)
	activeNeighBours := oneDimensionalGrid(neighbours).active()

	if p.getCube(w, z, y, x).isActive() {
		if activeNeighBours != 2 && activeNeighBours != 3 {
			gridToUpdate.setCube(w, z, y, x, inactive)
		}
	} else {
		if activeNeighBours == 3 {
			gridToUpdate.setCube(w, z, y, x, active)
		}
	}
}

func (p *pocketDimension) getCube(w, z, y, x int) cube {
	return p.grid[w][z][y][x]
}

func (p *pocketDimension) getNeighbours(wIndex, zIndex, yIndex, xIndex int) []cube {
	neighBours := []cube{}

	if p.fourDimensional {
		for w := wIndex - 1; w <= wIndex+1; w++ {
			for z := zIndex - 1; z <= zIndex+1; z++ {
				for y := yIndex - 1; y <= yIndex+1; y++ {
					for x := xIndex - 1; x <= xIndex+1; x++ {
						if w == wIndex && z == zIndex && y == yIndex && x == xIndex {
							continue
						}
						neighBours = append(neighBours, p.getCube(w, z, y, x))
					}
				}
			}
		}
	} else {
		for z := zIndex - 1; z <= zIndex+1; z++ {
			for y := yIndex - 1; y <= yIndex+1; y++ {
				for x := xIndex - 1; x <= xIndex+1; x++ {
					if z == zIndex && y == yIndex && x == xIndex {
						continue
					}
					neighBours = append(neighBours, p.getCube(0, z, y, x))
				}
			}
		}

	}
	return neighBours
}

func copyGrid(grid fourDimensionalGrid) fourDimensionalGrid {
	newGrid := make(fourDimensionalGrid, len(grid))
	for w, threeDGrid := range grid {
		newThreeDGrid := make(threeDimensionalGrid, len(grid[w]))
		for z, level := range threeDGrid {
			newLevel := make(twoDimensionalGrid, len(threeDGrid[z]))
			for y, row := range level {
				newRow := make(oneDimensionalGrid, len(level[y]))
				for x, cube := range row {
					newRow[x] = cube
				}
				newLevel[y] = newRow
			}
			newThreeDGrid[z] = newLevel
		}
		newGrid[w] = newThreeDGrid
	}

	return newGrid
}

type fourDimensionalGrid []threeDimensionalGrid

func (g fourDimensionalGrid) active() int {
	activeCount := 0
	for _, threeDGrid := range g {
		activeCount += threeDGrid.active()
	}
	return activeCount
}

func (g fourDimensionalGrid) print() {
	for w, threeDGrid := range g {
		threeDGrid.print(w)
	}
}

func (g fourDimensionalGrid) setCube(w, z, y, x int, cube cube) {
	g[w][z][y][x] = cube
}

type threeDimensionalGrid []twoDimensionalGrid

func (g threeDimensionalGrid) active() int {
	activeCount := 0
	for _, twoDGrid := range g {
		activeCount += twoDGrid.active()
	}
	return activeCount
}

func (g threeDimensionalGrid) print(fourthDimension int) {
	for z, level := range g {
		fmt.Printf("z=%d, w=%d\n", z, fourthDimension)
		level.print()
	}
}

func (g threeDimensionalGrid) setCube(z, y, x int, cube cube) {
	g[z][y][x] = cube
}

type twoDimensionalGrid []oneDimensionalGrid

func (g twoDimensionalGrid) active() int {
	activeCount := 0
	for _, oneDGrid := range g {
		activeCount += oneDGrid.active()
	}
	return activeCount
}

func (g twoDimensionalGrid) print() {
	for _, row := range g {
		row.print()
	}
}

type oneDimensionalGrid []cube

func (g oneDimensionalGrid) active() int {
	activeCount := 0
	for _, cube := range g {
		if cube.isActive() {
			activeCount++
		}
	}
	return activeCount
}

func (g oneDimensionalGrid) print() {
	for _, cube := range g {
		fmt.Printf("%s", string(cube))
	}
	fmt.Printf("\n")
}

type cube rune

func (c cube) isActive() bool {
	if c == active {
		return true
	}
	return false
}

const active cube = '#'
const inactive cube = '.'

func new(initialState []string, cycles int, fourDimensional bool) *pocketDimension {
	p := &pocketDimension{
		grid:            newInactiveGrid(len(initialState), len(initialState[0]), cycles, fourDimensional),
		cycles:          cycles,
		fourDimensional: fourDimensional,
	}
	offset := cycles + 1
	z := offset
	w := 0
	if fourDimensional {
		w = offset
	}
	for y, row := range initialState {
		for x, state := range row {
			offsetX := x + offset
			offsetY := y + offset
			p.grid[w][z][offsetY][offsetX] = cube(state)
		}
	}
	return p
}

func newInactiveGrid(rows, columns int, cycles int, fourDimensional bool) fourDimensionalGrid {
	maxW := 1
	if fourDimensional {
		maxW = ((cycles + 1) * 2) + 1
	}
	maxZ := ((cycles + 1) * 2) + 1
	maxY := ((cycles + 1) * 2) + rows
	maxX := ((cycles + 1) * 2) + columns

	grid := make(fourDimensionalGrid, maxW)
	for w := 0; w <= maxW-1; w++ {
		threeDGrid := make(threeDimensionalGrid, maxZ)
		for z := 0; z <= maxZ-1; z++ {
			twoDGrid := make(twoDimensionalGrid, maxY)
			for y := 0; y <= maxY-1; y++ {
				oneDGrid := make(oneDimensionalGrid, maxX)
				for x := 0; x <= maxX-1; x++ {
					oneDGrid[x] = '.'
				}
				twoDGrid[y] = oneDGrid
			}
			threeDGrid[z] = twoDGrid
		}
		grid[w] = threeDGrid
	}

	return grid
}

func main() {
	initialState, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	p := new(initialState, 6, false)

	p.bootProcess()

	fmt.Printf("part 1: %d\n", p.active())

	fourD := new(initialState, 6, true)

	fourD.bootProcess()

	fmt.Printf("part 2: %d\n", fourD.active())
}
