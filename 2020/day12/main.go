package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type ferry struct {
	xDistanceTravelled int
	yDistanceTravelled int
	orientations       []orientation
	currentOrientation int
}

type orientation struct {
	y int
	x int
}

func newFerry() *ferry {
	orientations := []orientation{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	return &ferry{
		orientations:       orientations,
		currentOrientation: 1,
	}
}

func (f *ferry) navigate(directions []string) {
	for _, direction := range directions {
		f.travel(direction)
	}
}

func (f *ferry) travel(direction string) {
	var action string
	var value int
	fmt.Sscanf(direction, "%1s%d", &action, &value)

	switch action {
	case "N":
		f.yDistanceTravelled += value
	case "S":
		f.yDistanceTravelled -= value
	case "E":
		f.xDistanceTravelled += value
	case "W":
		f.xDistanceTravelled -= value
	case "L":
		f.changeOrientation(-1, value)
	case "R":
		f.changeOrientation(1, value)
	case "F":
		f.yDistanceTravelled += f.orientations[f.currentOrientation].y * value
		f.xDistanceTravelled += f.orientations[f.currentOrientation].x * value
	}
}
func (f *ferry) changeOrientation(rotation int, magnitude int) {
	currentRotation := 0
	for {
		if currentRotation == magnitude {
			return
		}
		f.currentOrientation = (len(f.orientations) + f.currentOrientation + rotation) % len(f.orientations)
		currentRotation += 90
	}
}

func (f *ferry) manhattenDistanceTravelled() int {
	return int(math.Abs(float64(f.xDistanceTravelled))) + int(math.Abs(float64(f.yDistanceTravelled)))
}

type ferryWithWayPoint struct {
	xDistanceTravelled int
	yDistanceTravelled int
	wayPoint           *wayPoint
}

type wayPoint struct {
	xOffset int
	yOffset int
}

func newFerryWithWayPoint() *ferryWithWayPoint {
	return &ferryWithWayPoint{
		wayPoint: &wayPoint{
			xOffset: 10,
			yOffset: 1,
		},
	}
}

func (f *ferryWithWayPoint) navigate(directions []string) {
	for _, direction := range directions {
		f.travel(direction)
	}
}

func (f *ferryWithWayPoint) travel(direction string) {
	var action string
	var value int
	fmt.Sscanf(direction, "%1s%d", &action, &value)

	switch action {
	case "N":
		f.wayPoint.yOffset += value
	case "S":
		f.wayPoint.yOffset -= value
	case "E":
		f.wayPoint.xOffset += value
	case "W":
		f.wayPoint.xOffset -= value
	case "L":
		f.moveWayPoint(false, value)
	case "R":
		f.moveWayPoint(true, value)
	case "F":
		f.xDistanceTravelled += f.wayPoint.xOffset * value
		f.yDistanceTravelled += f.wayPoint.yOffset * value
	}
}

func (f *ferryWithWayPoint) moveWayPoint(clockWise bool, magnitude int) {
	reflections := []int{-1, 1}
	if clockWise {
		reflections = []int{1, -1}
	}

	currentRotation := 0
	for {
		if currentRotation == magnitude {
			return
		}
		f.wayPoint.xOffset, f.wayPoint.yOffset = f.wayPoint.yOffset, f.wayPoint.xOffset
		f.wayPoint.xOffset *= reflections[0]
		f.wayPoint.yOffset *= reflections[1]
		currentRotation += 90
	}
}

func (f *ferryWithWayPoint) manhattenDistanceTravelled() int {
	return int(math.Abs(float64(f.xDistanceTravelled))) + int(math.Abs(float64(f.yDistanceTravelled)))
}

func main() {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	directions := strings.Split(strings.TrimSpace(fileString), "\n")

	ferry := newFerry()
	ferry.navigate(directions)

	fmt.Printf("part 1: %d\n", ferry.manhattenDistanceTravelled())

	ferryWithWayPoint := newFerryWithWayPoint()
	ferryWithWayPoint.navigate(directions)

	fmt.Printf("part 2: %d\n", ferryWithWayPoint.manhattenDistanceTravelled())
}
