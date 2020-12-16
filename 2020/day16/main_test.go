package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanningErrorRate(t *testing.T) {
	testCases := []struct {
		inputFile                 string
		expectedScanningErrorRate int
	}{
		{
			inputFile:                 "input.sample1",
			expectedScanningErrorRate: 71,
		},
	}

	for _, testCase := range testCases {
		rules, _, nearbyTickets, err := loadInput(testCase.inputFile)
		assert.NoError(t, err)

		_, scanningErrorRate := scanningErrorRate(rules, nearbyTickets)
		assert.Equal(t, testCase.expectedScanningErrorRate, scanningErrorRate)
	}
}

func TestFindRulePositions(t *testing.T) {
	testCases := []struct {
		ticketRules           rules
		validTickets          tickets
		expectedRulePositions map[string]int
	}{
		{
			ticketRules: rules{
				"class": []rule{
					{
						lowerLimit: 0,
						upperLimit: 1,
					},
					{
						lowerLimit: 4,
						upperLimit: 19,
					},
				},
				"row": []rule{
					{
						lowerLimit: 0,
						upperLimit: 5,
					},
					{
						lowerLimit: 8,
						upperLimit: 19,
					},
				},
				"seat": []rule{
					{
						lowerLimit: 0,
						upperLimit: 13,
					},
					{
						lowerLimit: 16,
						upperLimit: 19,
					},
				},
			},
			validTickets: tickets{
				ticket{3, 9, 18},
				ticket{15, 1, 5},
				ticket{5, 14, 9},
			},
			expectedRulePositions: map[string]int{
				"class": 1,
				"row":   0,
				"seat":  2,
			},
		},
	}

	for _, testCase := range testCases {
		rulePositions := findRulePositions(testCase.ticketRules, testCase.validTickets)

		for rule, position := range rulePositions {
			assert.Equal(t, testCase.expectedRulePositions[rule], position)
		}
	}
}
