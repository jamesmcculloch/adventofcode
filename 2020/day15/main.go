package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type memoryGame struct {
	startingNumbers []int
	memory          map[int][]int
}

func newMemoryGame(startingConfig string) (*memoryGame, error) {
	values := strings.Split(startingConfig, ",")

	startingNumbers := make([]int, len(values))
	for index, value := range values {
		number, err := strconv.Atoi(value)
		if err != nil {
			return &memoryGame{}, err
		}
		startingNumbers[index] = number
	}

	return &memoryGame{
		startingNumbers: startingNumbers,
		memory:          make(map[int][]int),
	}, nil
}

func (mg *memoryGame) reset() {
	mg.memory = make(map[int][]int)
}

func (mg *memoryGame) play(turns int) int {
	var lastSpokenNumber int

	for index, startingNumber := range mg.startingNumbers {
		lastSpokenNumber = startingNumber
		mg.rememberTurn(index+1, lastSpokenNumber)
	}

	for turn := len(mg.startingNumbers) + 1; turn <= turns; turn++ {
		if !!!mg.seenNumber(lastSpokenNumber) {
			lastSpokenNumber = 0
			mg.rememberTurn(turn, lastSpokenNumber)
			continue
		}

		if mg.seenNumberOnce(lastSpokenNumber) {
			lastSpokenNumber = 0
			mg.rememberTurn(turn, lastSpokenNumber)
			continue
		}

		lastSpoken, previous := mg.recallTurnsForNumber(lastSpokenNumber)
		new := lastSpoken - previous
		lastSpokenNumber = new
		mg.rememberTurn(turn, lastSpokenNumber)
	}

	return lastSpokenNumber
}

func (mg *memoryGame) seenNumber(number int) bool {
	if _, ok := mg.memory[number]; !ok {
		return false
	}
	return true
}

func (mg *memoryGame) seenNumberOnce(number int) bool {
	if _, ok := mg.memory[number]; !ok {
		return false
	}
	if len(mg.memory[number]) == 1 {
		return true
	}
	return false
}

func (mg *memoryGame) rememberTurn(turn int, spoken int) {
	if _, ok := mg.memory[spoken]; !ok {
		mg.memory[spoken] = []int{turn}
	}
	if len(mg.memory[spoken]) == 1 {
		mg.memory[spoken] = append(mg.memory[spoken], turn)
	}
	if len(mg.memory[spoken]) == 2 {
		lastSpoken := mg.memory[spoken][len(mg.memory[spoken])-1]
		mg.memory[spoken][len(mg.memory[spoken])-2] = lastSpoken
		mg.memory[spoken][len(mg.memory[spoken])-1] = turn
	}
}

func (mg *memoryGame) recallTurnsForNumber(number int) (int, int) {
	lastSpoken := mg.memory[number][len(mg.memory[number])-1]
	previous := mg.memory[number][len(mg.memory[number])-2]
	return lastSpoken, previous
}

func main() {
	gameConfig, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	game, err := newMemoryGame(gameConfig[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", game.play(2020))
	game.reset()
	fmt.Printf("part 2: %d\n", game.play(30000000))
}
