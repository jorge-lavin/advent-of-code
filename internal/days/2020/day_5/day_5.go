package days

import (
	"math"
)

type seat struct {
	row    int
	column int
}

func findSeatForBoardingPass(pass string) seat {
	row := 0
	column := 0

	min, max := 0, 127
	for i := 0; i < 6; i++ {
		// front
		if pass[i] == 'F' {
			min, max = lowerHalf(min, max)
		}
		// back
		if pass[i] == 'B' {
			min, max = upperHalf(min, max)
		}
	}

	if pass[6] == 'F' {
		row = min
	}

	if pass[6] == 'B' {
		row = max
	}

	min, max = 0, 7
	for i := 7; i < 9; i++ {
		// front
		if pass[i] == 'L' {
			min, max = lowerHalf(min, max)
		}
		// back
		if pass[i] == 'R' {
			min, max = upperHalf(min, max)
		}
	}

	if pass[9] == 'L' {
		column = min
	}

	if pass[9] == 'R' {
		column = max
	}

	return seat{row: row, column: column}
}

func lowerHalf(min int, max int) (int, int) {
	halfDifference := halfDifference(min, max)
	return min, max - halfDifference
}

func upperHalf(min int, max int) (int, int) {
	halfDifference := halfDifference(min, max)
	return min + halfDifference, max
}

func halfDifference(min int, max int) int {
	return int(math.Ceil(float64(max-min) / 2))
}

func (m seat) seatID() int {
	return m.row*8 + m.column
}

func findMySeat(others []seat) seat {
	maxID := 127*8 + 7
	minID := 0

	seen := map[int]bool{}

	for i := minID; i <= maxID; i++ {
		seen[i] = false
	}

	for _, seat := range others {
		currentID := seat.seatID()
		seen[currentID] = true
	}
	return seat{}
}
