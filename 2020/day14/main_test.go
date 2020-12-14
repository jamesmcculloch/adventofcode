package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadProgramAndSumMemory(t *testing.T) {
	testCases := []struct {
		intructions []string
		expectedSum int64
		isVerison2  bool
	}{
		{
			intructions: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			expectedSum: 165,
			isVerison2:  false,
		},
		{
			intructions: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			expectedSum: 208,
			isVerison2:  true,
		},
	}

	for _, testCase := range testCases {
		computer := new(testCase.isVerison2)

		computer.loadProgram(testCase.intructions)

		assert.Equal(t, int64(testCase.expectedSum), int64(computer.sumMemory()))
	}
}

func TestApplyMask(t *testing.T) {
	testCases := []struct {
		mask                string
		value               int64
		expectedMaskedValue int64
		isVerison2          bool
	}{
		{
			mask:                "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			value:               11,
			expectedMaskedValue: 73,
			isVerison2:          false,
		},
		{
			mask:                "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			value:               101,
			expectedMaskedValue: 101,
			isVerison2:          false,
		},
		{
			mask:                "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			value:               0,
			expectedMaskedValue: 64,
			isVerison2:          false,
		},
	}

	for _, testCase := range testCases {
		computer := new(testCase.isVerison2)

		computer.mask = testCase.mask

		maskedValue := computer.applyMaskToValue(testCase.value)

		assert.Equal(t, testCase.expectedMaskedValue, maskedValue)
	}
}
