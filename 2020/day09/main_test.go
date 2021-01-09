package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestFirstInvalidNumber(t *testing.T) {
	xmas, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	_, invalidNumber := firstInvalidNumber(xmas, 5)
	assert.Equal(t, 127, invalidNumber)
}

func TestEncryptionWeakness(t *testing.T) {
	xmas, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}

	_, invalidNumber := firstInvalidNumber(xmas, 5)
	encryptionWeakness := encryptionWeakness(xmas, invalidNumber)

	assert.Equal(t, 62, encryptionWeakness)
}
