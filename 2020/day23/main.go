package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type crabCups struct {
	currentCup   *cup
	minCupValue  int
	maxCupValue  int
	pickedUpCups map[int]bool
	cups         map[int]*cup
}

type cup struct {
	value int
	next  *cup
}

func new(config string, millionTotal bool) *crabCups {
	newGame := &crabCups{
		minCupValue:  utils.MaxInt,
		maxCupValue:  utils.MinInt,
		cups:         make(map[int]*cup),
		pickedUpCups: make(map[int]bool),
	}

	var lastCup *cup
	for _, cupConfig := range config {
		cupValue, err := strconv.Atoi(string(cupConfig))
		if err != nil {
			panic(err)
		}
		newCup := &cup{
			value: cupValue,
		}
		newGame.cups[newCup.value] = newCup

		if newGame.currentCup == nil {
			newGame.currentCup = newCup
			lastCup = newCup
		} else {
			lastCup.next = newCup
			lastCup = newCup
		}

		if cupValue > newGame.maxCupValue {
			newGame.maxCupValue = cupValue
		}
		if cupValue < newGame.minCupValue {
			newGame.minCupValue = cupValue
		}
	}

	if millionTotal {
		for i := 10; i <= 1000000; i++ {
			cupValue := i
			newCup := &cup{
				value: cupValue,
			}
			newGame.cups[newCup.value] = newCup

			if newGame.currentCup == nil {
				newGame.currentCup = newCup
				lastCup = newCup
			} else {
				lastCup.next = newCup
				lastCup = newCup
			}

			if cupValue > newGame.maxCupValue {
				newGame.maxCupValue = cupValue
			}
			if cupValue < newGame.minCupValue {
				newGame.minCupValue = cupValue
			}
		}
	}

	lastCup.next = newGame.currentCup

	return newGame
}

func (c *crabCups) play(numberOfMoves int) {
	for i := 1; i <= numberOfMoves; i++ {
		pickedUp := c.pickUpCups(3)

		destinationCup := c.findDestinationCup()

		c.placeCups(destinationCup, pickedUp)

		c.newCurrentCup()
	}
}

func (c *crabCups) string() string {
	if c.currentCup == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(strconv.Itoa(c.currentCup.value))
	sb.WriteString(")")
	currentCup := c.currentCup.next
	for currentCup != c.currentCup {
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(currentCup.value))
		currentCup = currentCup.next
	}
	return sb.String()
}

func (c *crabCups) pickUpCups(cupCount int) *cup {
	startCup := c.currentCup.next
	currentCup := startCup
	c.pickedUpCups[currentCup.value] = true
	for count := 1; count < cupCount; count++ {
		currentCup = currentCup.next
		c.pickedUpCups[currentCup.value] = true
	}
	endCup := currentCup.next
	currentCup.next = nil
	c.currentCup.next = endCup
	return startCup
}

func (c *crabCups) findDestinationCup() *cup {
	var currentTarget int
	if c.currentCup.value == c.minCupValue {
		currentTarget = c.maxCupValue
	} else {
		currentTarget = c.currentCup.value - 1
	}

	for {
		if _, pickedUp := c.pickedUpCups[currentTarget]; !pickedUp {
			break
		}
		if currentTarget == c.minCupValue {
			currentTarget = c.maxCupValue
		} else {
			currentTarget--
		}
	}
	return c.cups[currentTarget]
}

func (c *crabCups) placeCups(destinationCup *cup, additionalCups *cup) {
	afterDestination := destinationCup.next
	cupToAddTo := destinationCup
	for additionalCups != nil {
		cupToAddTo.next = additionalCups
		cupToAddTo = additionalCups
		additionalCups = additionalCups.next
	}
	cupToAddTo.next = afterDestination
}

func (c *crabCups) newCurrentCup() {
	c.currentCup = c.currentCup.next
	c.pickedUpCups = make(map[int]bool)
}

func (c *crabCups) resultingOrder() string {
	var cupWithValueOne *cup
	var sb strings.Builder
	currentCup := c.currentCup
	for {
		if currentCup.value == 1 {
			cupWithValueOne = currentCup
			break
		}
		currentCup = currentCup.next
	}

	currentCup = cupWithValueOne.next
	for currentCup != cupWithValueOne {
		sb.WriteString(strconv.Itoa(currentCup.value))
		currentCup = currentCup.next
	}
	return sb.String()
}

func (c *crabCups) part2() int {
	var cupWithValueOne *cup
	currentCup := c.currentCup
	for {
		if currentCup.value == 1 {
			cupWithValueOne = currentCup
			break
		}
		currentCup = currentCup.next
	}

	return cupWithValueOne.next.value * cupWithValueOne.next.next.value
}

func main() {
	crabCupsConfig, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	game := new(crabCupsConfig[0], false)

	game.play(100)

	fmt.Printf("part 1: %s\n", game.resultingOrder())

	game = new(crabCupsConfig[0], true)

	game.play(10000000)

	fmt.Printf("part 2: %d\n", game.part2())
}
