package utils

// FindMaxInt finds the maximum int in a unsorted slice of ints
func FindMaxInt(ints []int) int {
	var currentMax int

	for _, candidate := range ints {
		if candidate > currentMax {
			currentMax = candidate
		}
	}
	return currentMax
}
