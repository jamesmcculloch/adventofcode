package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	passports, err := loadPassports("input.sample1")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	assert.Equal(t, 2, numberOfValidPassports1(passports))
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		inputFile                  string
		expectedValidPassportCount int
	}{
		{
			inputFile:                  "input.sample2",
			expectedValidPassportCount: 0,
		},
		{
			inputFile:                  "input.sample3",
			expectedValidPassportCount: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.inputFile, func(t *testing.T) {
			passports, err := loadPassports(testCase.inputFile)
			if err != nil {
				assert.Fail(t, "failed to load input sample", err)
			}

			assert.Equal(t, testCase.expectedValidPassportCount, numberOfValidPassports2(passports))
		})
	}
}
