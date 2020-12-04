package days

import (
	"testing"

	"github.com/lavinj87/advent/internal/lib"
)

func Test_day1Tuple2(t *testing.T) {
	t.Run("", func(t *testing.T) {
		actual := day1Tuple2("day_1_input.txt")
		expected := 719796
		if actual != expected {
			t.Errorf("Expected:%v. Actual:%v", expected, actual)
		}
	})
}

func Test_day1Tuple3(t *testing.T) {
	t.Run("", func(t *testing.T) {
		actual := day1Tuple3("day_1_input.txt")
		expected := 144554112
		if actual != expected {
			t.Errorf("Expected:%v. Actual:%v", expected, actual)
		}
	})
}

func Test_findTuple2ThatSums2020(t *testing.T) {

	tests := []struct {
		given    []int
		expected lib.Tuple2Int
	}{{
		given:    []int{1010, 1, 2, 1010},
		expected: lib.Tuple2Int{One: 1010, Two: 1010},
	}, {
		given:    []int{1, 2, 3},
		expected: lib.Tuple2Int{One: 1, Two: 1},
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := findTuple2ThatSums2020(tt.given)
			if tt.expected.Sum() != actual.Sum() {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_findTuple3ThatSums2020(t *testing.T) {

	tests := []struct {
		given    []int
		expected lib.Tuple3Int
	}{{
		given:    []int{673, 673, 674},
		expected: lib.Tuple3Int{One: 673, Two: 673, Three: 674},
	}, {
		given:    []int{1, 2, 3},
		expected: lib.Tuple3Int{One: 1, Two: 1, Three: 1},
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := findTuple3ThatSums2020(tt.given)
			if tt.expected.Sum() != actual.Sum() {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}
