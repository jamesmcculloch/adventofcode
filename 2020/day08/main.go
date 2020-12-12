package main

import (
	"fmt"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type gameConsole struct {
	accumulator        int
	instructionPointer int
	instructions       []*instruction
}

type instruction struct {
	operation string
	argument  int
}

func new(instructionDetails []string) *gameConsole {
	g := &gameConsole{}
	gameInstructions := make([]*instruction, len(instructionDetails))
	for index, instructionDetail := range instructionDetails {
		operation, argument := g.parseInstruction(instructionDetail)
		gameInstructions[index] = &instruction{
			operation: operation,
			argument:  argument,
		}
	}
	g.instructions = gameInstructions

	return g
}

func (g *gameConsole) run() ([]int, int) {
	instructionsRun := []int{}
	executedInstructions := make([]bool, len(g.instructions))
	for g.instructionPointer < len(g.instructions) {
		if executedInstructions[g.instructionPointer] {
			return instructionsRun, -1
		}
		executedInstructions[g.instructionPointer] = true
		instructionsRun = append(instructionsRun, g.instructionPointer)

		g.execute()
	}
	return instructionsRun, 0
}

func (g *gameConsole) execute() {
	currentInstruction := g.instructions[g.instructionPointer]

	switch currentInstruction.operation {
	case "acc":
		g.accumulator += currentInstruction.argument
		g.instructionPointer++
	case "jmp":
		g.instructionPointer += currentInstruction.argument
	case "nop":
		g.instructionPointer++
	default:
		panic(fmt.Sprintf("unkown instruction: %+v", currentInstruction))
	}
}

func (g *gameConsole) reset() {
	g.accumulator = 0
	g.instructionPointer = 0
}

func fixCorruptedInstruction(g *gameConsole, instructionsRun []int) (int, int) {
	for _, instructionToModify := range instructionsRun {
		g.modifyInstruction(instructionToModify)
		_, returnCode := g.run()

		if returnCode == 0 {
			return instructionToModify, g.accumulator
		}
		g.modifyInstruction(instructionToModify)
		g.reset()
	}
	return 0, 0
}

func (g *gameConsole) modifyInstruction(instructionPointer int) {
	instructionTomodify := g.instructions[instructionPointer]
	if instructionTomodify.operation == "jmp" {
		instructionTomodify.operation = "nop"
		return
	}
	if instructionTomodify.operation == "nop" {
		instructionTomodify.operation = "jmp"
		return
	}
}

func (g *gameConsole) parseInstruction(instruction string) (string, int) {
	var operation string
	var argument int
	fmt.Sscanf(instruction, "%s %d", &operation, &argument)
	return operation, argument
}

func (g *gameConsole) printRun(instructions []int) {
	fmt.Printf("Printing run with %d steps\n", len(instructions))
	for _, instructionPointer := range instructions {
		currentInstruction := g.instructions[instructionPointer]
		fmt.Printf("Executing %d, op %s, arg %d, acc %d\n",
			instructionPointer,
			currentInstruction.operation,
			currentInstruction.argument,
			g.accumulator)
	}
}

func main() {
	instructions, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	gc := new(instructions)

	instructionsRun, _ := gc.run()
	fmt.Printf("Part 1: %d\n", gc.accumulator)
	gc.reset()

	_, accumulatorOnceFixed := fixCorruptedInstruction(gc, instructionsRun)
	fmt.Printf("Part 2: %d\n", accumulatorOnceFixed)
}
