package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBagAndCount(t *testing.T) {
	testCases := []struct {
		ruleSection   string
		expectedBag   string
		expectedCount int
	}{
		{
			ruleSection:   " 2 muted yellow bags.",
			expectedBag:   "muted yellow",
			expectedCount: 2,
		},
		{
			ruleSection:   "2 muted yellow bags",
			expectedBag:   "muted yellow",
			expectedCount: 2,
		},
		{
			ruleSection:   "1 shiny gold bag.",
			expectedBag:   "shiny gold",
			expectedCount: 1,
		},
		{
			ruleSection:   "no other bags.",
			expectedBag:   "",
			expectedCount: 0,
		},
	}

	for _, testCase := range testCases {
		bag, count := getBagAndCount(testCase.ruleSection)

		assert.Equal(t, testCase.expectedBag, bag)
		assert.Equal(t, testCase.expectedCount, count)
	}
}

func TestGetOuterBag(t *testing.T) {
	testCases := []struct {
		ruleSection   string
		expectedBag   string
		expectedCount int
	}{
		{
			ruleSection: "dotted black bags ",
			expectedBag: "dotted black",
		},
	}

	for _, testCase := range testCases {
		bag := getOuterBag(testCase.ruleSection)

		assert.Equal(t, testCase.expectedBag, bag)
	}
}
