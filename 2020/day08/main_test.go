package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseInstruction(t *testing.T) {
	testCases := []struct {
		instruction       string
		expectedOperation string
		expectedArgument  int
	}{
		{
			instruction:       "nop +0",
			expectedOperation: "nop",
			expectedArgument:  0,
		},
		{
			instruction:       "acc -99",
			expectedOperation: "acc",
			expectedArgument:  -99,
		},
		{
			instruction:       "jmp +4",
			expectedOperation: "jmp",
			expectedArgument:  4,
		},
	}

	g := &gameConsole{}

	for _, testCase := range testCases {
		operation, argument := g.parseInstruction(testCase.instruction)

		assert.Equal(t, testCase.expectedOperation, operation)
		assert.Equal(t, testCase.expectedArgument, argument)
	}
}

func TestPart1(t *testing.T) {
	instructions, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	gc := new(instructions)

	gc.run()

	assert.Equal(t, 5, gc.accumulator)
}

func TestPart2(t *testing.T) {
	instructions, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	gc := new(instructions)
	instructionsRun, _ := gc.run()
	gc.reset()
	_, accumulatorOnceFixed := fixCorruptedInstruction(gc, instructionsRun)

	assert.Equal(t, 8, accumulatorOnceFixed)
}
