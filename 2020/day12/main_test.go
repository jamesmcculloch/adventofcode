package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFerryNavigate(t *testing.T) {
	testCases := []struct {
		directions       []string
		expectedDistance int
	}{
		{
			directions: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expectedDistance: 25,
		},
	}

	for _, testCase := range testCases {
		ferry := newFerry()

		ferry.navigate(testCase.directions)

		assert.Equal(t, testCase.expectedDistance, ferry.manhattenDistanceTravelled())
	}
}

func TestFerryWithWaypointNavigate(t *testing.T) {
	testCases := []struct {
		directions       []string
		expectedDistance int
	}{
		{
			directions: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expectedDistance: 286,
		},
	}

	for _, testCase := range testCases {
		ferryWithWayPoint := newFerryWithWayPoint()

		ferryWithWayPoint.navigate(testCase.directions)

		assert.Equal(t, testCase.expectedDistance, ferryWithWayPoint.manhattenDistanceTravelled())
	}
}
