package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindTwoNumbersThatSumToTarget(t *testing.T) {
	expenseReport := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	target := 2020

	first, second, found := findTwoNumbersThatSumToTarget(expenseReport, target)

	assert.True(t, found)
	assert.Equal(t, 1721, first)
	assert.Equal(t, 299, second)
}

func TestPart1(t *testing.T) {
	expenseReport, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}
	target := 2020

	testResult := part1(expenseReport, target)

	assert.Equal(t, 514579, testResult)
}

func TestFindThreeNumbersThatSumToTarget(t *testing.T) {
	expenseReport := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	target := 2020

	first, second, third, found := findThreeNumbersThatSumToTarget(expenseReport, target)

	assert.True(t, found)
	assert.Equal(t, 979, first)
	assert.Equal(t, 366, second)
	assert.Equal(t, 675, third)
}

func TestPart2(t *testing.T) {
	expenseReport, err := utils.LoadNumbersFromInput("input.sample")
	if err != nil {
		assert.Fail(t, "failed to load input sample", err)
	}
	target := 2020

	testResult := part2(expenseReport, target)

	assert.Equal(t, 241861950, testResult)
}
