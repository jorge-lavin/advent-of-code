package days

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_day2ByCount(t *testing.T) {
	t.Run("", func(t *testing.T) {
		actual := day2("day_2_input.txt", isValidPasswordByCount)
		expected := 622
		if actual != expected {
			t.Errorf("Expected:%v. Actual:%v", expected, actual)
		}
	})
}

func Test_day2ByPosition(t *testing.T) {
	t.Run("", func(t *testing.T) {
		actual := day2("day_2_input.txt", isValidPasswordByPosition)
		expected := 263
		if actual != expected {
			t.Errorf("Expected:%v. Actual:%v", expected, actual)
		}
	})
}

func Test_isValidPasswordByCount(t *testing.T) {
	tests := []struct {
		record   *passwordRecord
		expected bool
	}{{
		record:   &passwordRecord{min: 4, max: 8, repeated: "n", password: "dnjjrtclnzdnghnbnn"},
		expected: true,
	}, {
		record:   &passwordRecord{min: 5, max: 6, repeated: "r", password: "rrrrcqr"},
		expected: true,
	}, {
		record:   &passwordRecord{min: 1, max: 3, repeated: "b", password: "cdefg"},
		expected: false,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := isValidPasswordByCount(tt.record); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_isValidPasswordByPosition(t *testing.T) {
	tests := []struct {
		record   *passwordRecord
		expected bool
	}{{
		record:   &passwordRecord{min: 1, max: 3, repeated: "a", password: "abcde"},
		expected: true,
	}, {
		record:   &passwordRecord{min: 1, max: 3, repeated: "b", password: "cdefg"},
		expected: false,
	}, {
		record:   &passwordRecord{min: 2, max: 9, repeated: "c", password: "ccccccccc"},
		expected: false,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := isValidPasswordByPosition(tt.record); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_newPasswordRecord(t *testing.T) {

	tests := []struct {
		line     string
		expected *passwordRecord
	}{{
		line:     "4-8 n: dnjjrtclnzdnghnbnn",
		expected: &passwordRecord{min: 4, max: 8, repeated: "n", password: "dnjjrtclnzdnghnbnn"},
	}, {
		line:     "5-6 r: rrrrcqr",
		expected: &passwordRecord{min: 5, max: 6, repeated: "r", password: "rrrrcqr"},
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := newPasswordRecord(tt.line); !cmp.Equal(tt.expected, actual, cmp.AllowUnexported(passwordRecord{})) {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}
