package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		expression         string
		expectedResult     int
		additionPreference bool
	}{
		{
			expression:         "1 + 2 * 3 + 4 * 5 + 6",
			expectedResult:     71,
			additionPreference: false,
		},
		{
			expression:         "2 * 3 + (4 * 5)",
			expectedResult:     26,
			additionPreference: false,
		},
		{
			expression:         "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expectedResult:     437,
			additionPreference: false,
		},
		{
			expression:         "(8 * 3 + 9 + 3 * 4 * 3) + 5",
			expectedResult:     437,
			additionPreference: false,
		},
		{
			expression:         "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expectedResult:     12240,
			additionPreference: false,
		},
		{
			expression:         "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expectedResult:     13632,
			additionPreference: false,
		},
		{
			expression:         "1 + (2 * 3) + (4 * (5 + 6))",
			expectedResult:     51,
			additionPreference: false,
		},
		{
			expression:         "(2 * 3)",
			expectedResult:     6,
			additionPreference: false,
		},
		{
			expression:         "3 * ((2 * 3))",
			expectedResult:     18,
			additionPreference: false,
		},
		{
			expression:         "((2 * 3) + (2 * 5))",
			expectedResult:     16,
			additionPreference: false,
		},
		{
			expression:         "2 * (2 * 3) * 5",
			expectedResult:     60,
			additionPreference: false,
		},
		{
			expression:         "4 * ((4 * 4 + 9) + 8) * 2",
			expectedResult:     264,
			additionPreference: false,
		},
		{
			expression:         "3 * ((7 * 8 * 9 + 7) + (2 * 4 + 4 + 9) * 9 + 8)",
			expectedResult:     14388,
			additionPreference: false,
		},
		{
			expression:         "1 + 2 * 3 + 4 * 5 + 6",
			expectedResult:     231,
			additionPreference: true,
		},
		{
			expression:         "1 + (2 * 3) + (4 * (5 + 6))",
			expectedResult:     51,
			additionPreference: true,
		},
		{
			expression:         "2 * 3 + (4 * 5)",
			expectedResult:     46,
			additionPreference: true,
		},
		{
			expression:         "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expectedResult:     1445,
			additionPreference: true,
		},
		{
			expression:         "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expectedResult:     669060,
			additionPreference: true,
		},
		{
			expression:         "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expectedResult:     23340,
			additionPreference: true,
		},
		{
			expression:         "1 + (2 * 3) + (4 * (5 + 6))",
			expectedResult:     51,
			additionPreference: true,
		},
		{
			expression:         "(2 * 3)",
			expectedResult:     6,
			additionPreference: true,
		},
		{
			expression:         "3 * ((2 * 3))",
			expectedResult:     18,
			additionPreference: true,
		},
		{
			expression:         "((2 * 3) + (2 * 5))",
			expectedResult:     16,
			additionPreference: true,
		},
		{
			expression:         "2 * (2 * 3) * 5",
			expectedResult:     60,
			additionPreference: true,
		},
		{
			expression:         "4 * ((4 * 4 + 9) + 8) * 2",
			expectedResult:     480,
			additionPreference: true,
		},
		{
			expression:         "3 * ((7 * 8 * 9 + 7) + (2 * 4 + 4 + 9) * 9 + 8)",
			expectedResult:     47430,
			additionPreference: true,
		},
		{
			expression:         "6 * 3 + ((2 * 9 + 4) + (8 * 7 * 8 * 9 * 7 + 7) + 3 + 8 + (2 * 5 + 4 + 2 * 2 * 3)) + 2",
			expectedResult:     339732,
			additionPreference: true,
		},
	}

	for id, testCase := range testCases {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			result := evaluate(testCase.expression, testCase.additionPreference)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}
