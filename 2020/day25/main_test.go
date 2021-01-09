package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoopSize(t *testing.T) {
	testCases := []struct {
		key              int
		expectedLoopSize int
	}{
		{
			5764801,
			8,
		},
		{
			17807724,
			11,
		},
	}

	for _, testCase := range testCases {
		size := loopSize(testCase.key)

		assert.Equal(t, testCase.expectedLoopSize, size)
	}
}

func TestEncryptionKey(t *testing.T) {
	testCases := []struct {
		loopSize      int
		subjectNumber int
		expectedKey   int
	}{
		{
			8,
			17807724,
			14897079,
		},
		{
			11,
			5764801,
			14897079,
		},
	}

	for _, testCase := range testCases {
		key := encryptionKeyValue(testCase.loopSize, testCase.subjectNumber)

		assert.Equal(t, testCase.expectedKey, key)
	}
}

func TestPart1(t *testing.T) {
	keys, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	assert.Equal(t, 14897079, encryptionKey(keys[0], keys[1]))
}
