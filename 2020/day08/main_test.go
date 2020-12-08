package main

import (
	"testing"

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
