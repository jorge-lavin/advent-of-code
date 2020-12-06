package days

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lavinj87/advent/internal/lib"
)

func Test_findSeat(t *testing.T) {
	cases := []struct {
		boardingPass string
		expected     seat
	}{{
		boardingPass: "FBFBBFFRLR",
		expected:     seat{row: 44, column: 5},
	}, {
		boardingPass: "BFFFBBFRRR",
		expected:     seat{row: 70, column: 7},
	}, {
		boardingPass: "FFFBBBFRRR",
		expected:     seat{row: 14, column: 7},
	}, {
		boardingPass: "BBFFBBFRLL",
		expected:     seat{row: 102, column: 4},
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := findSeatForBoardingPass(tt.boardingPass)
			expected := tt.expected
			if !cmp.Equal(expected, actual, cmp.AllowUnexported(seat{})) {
				t.Errorf("Expected:%v, Actual:%v", expected, actual)
			}
		})
	}
}

func Test_seatID(t *testing.T) {
	cases := []struct {
		seat     seat
		expected int
	}{{
		seat:     seat{row: 44, column: 5},
		expected: 357,
	}, {
		seat:     seat{row: 70, column: 7},
		expected: 567,
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := tt.seat.seatID()
			expected := tt.expected
			if expected != actual {
				t.Errorf("Expected:%v, Actual:%v", expected, actual)
			}
		})
	}
}

func Test_HighestSeatID(t *testing.T) {
	expected := 998
	actual := 0
	lib.Lines("day_5_input.txt", func(line string) {
		currentID := findSeatForBoardingPass(line).seatID()
		if currentID > actual {
			actual = currentID
		}
	})

	if expected != actual {
		t.Errorf("Expected:%v, Actual:%v", expected, actual)
	}
}

func Test_FindMySeat(t *testing.T) {
	expected := seat{row: 1, column: 2}
	otherSeats := []seat{}
	lib.Lines("day_5_input.txt", func(line string) {
		seat := findSeatForBoardingPass(line)
		otherSeats = append(otherSeats, seat)
	})
	actual := findMySeat(otherSeats)

	if !cmp.Equal(expected, actual, cmp.AllowUnexported(seat{})) {
		t.Errorf("Expected:%v, Actual:%v", expected, actual)
	}

}
