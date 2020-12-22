package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	passwords, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	candidates, err := getPasswordsToValidate(passwords)
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	assert.Equal(t, 2, validLetterCountPasswordCount(candidates))
}

func TestPart2(t *testing.T) {
	passwords, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	candidates, err := getPasswordsToValidate(passwords)
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	assert.Equal(t, 1, validPositionPasswordCount(candidates))
}
