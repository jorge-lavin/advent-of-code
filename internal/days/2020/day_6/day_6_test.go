package days

import (
	"testing"
)

func Test_CountAnswers(t *testing.T) {
	cases := []struct {
		answers  []string
		expected int
	}{{
		answers:  []string{"abc"},
		expected: 3,
	}, {
		answers:  []string{"a", "b", "c"},
		expected: 3,
	}, {
		answers:  []string{"ab", "ac"},
		expected: 3,
	}, {
		answers:  []string{"a", "a", "a", "a"},
		expected: 1,
	}, {
		answers:  []string{"b"},
		expected: 1,
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := countAnswers(tt.answers)
			expected := tt.expected
			if actual != expected {
				t.Errorf("Expected:%v, Actual:%v", expected, actual)
			}
		})
	}
}

func Test_CountAnswersFromFile(t *testing.T) {
	expected := 6249
	actual := countAnswersFromFile("day_6_input.txt")

	if actual != expected {
		t.Errorf("Expected:%v, Actual:%v", expected, actual)
	}
}

func Test_CountEveryoneAnswers(t *testing.T) {
	cases := []struct {
		answers  []string
		expected int
	}{{
		answers:  []string{"abc"},
		expected: 3,
	}, {
		answers:  []string{"a", "b", "c"},
		expected: 0,
	}, {
		answers:  []string{"ab", "ac"},
		expected: 1,
	}, {
		answers:  []string{"a", "a", "a", "a"},
		expected: 1,
	}, {
		answers:  []string{"b"},
		expected: 1,
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := countEveryoneAnswers(tt.answers)
			expected := tt.expected
			if actual != expected {
				t.Errorf("Expected:%v, Actual:%v", expected, actual)
			}
		})
	}
}

func Test_CountEveryoneAnswersFromFile(t *testing.T) {
	expected := 6249
	actual := countEveryoneAnswersFromFile("day_6_input.txt")

	if actual != expected {
		t.Errorf("Expected:%v, Actual:%v", expected, actual)
	}
}
