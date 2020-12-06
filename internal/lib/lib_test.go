package lib

import (
	"log"
	"strings"
	"testing"
)

func Test_linesPathFound(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fileLines := ""
		action := func(line string) {
			fileLines = fileLines + line
		}

		err := Lines("../../go.mod", action)

		if err != nil {
			t.Errorf("Expected no error but found:%s", err)
		}

		if len(fileLines) == 0 {
			t.Errorf("Expected some content at the file.")
		}
	})
}

func Test_linesPathNotFound(t *testing.T) {
	t.Run("", func(t *testing.T) {
		err := Lines("this-is-not-a-file", func(line string) {})

		if err == nil {
			t.Error("Expected error but none found", err)
		}
	})
}

func Test_StringToInt(t *testing.T) {
	tests := []struct {
		text     string
		expected int
	}{{
		text:     "2",
		expected: 2,
	}, {
		text:     "3",
		expected: 3,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := StringToInt(tt.text); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_Tuple2IntMultiply(t *testing.T) {
	expected := 6
	actual := Tuple2Int{2, 3}.Multiply()
	AssertEquals(t, expected, actual)
}

func Test_Tuple2IntSum(t *testing.T) {
	expected := 5
	actual := Tuple2Int{2, 3}.Sum()
	AssertEquals(t, expected, actual)
}

func Test_Tuple3IntMultiply(t *testing.T) {
	expected := 24
	actual := Tuple3Int{2, 3, 4}.Multiply()
	AssertEquals(t, expected, actual)
}

func Test_Tuple3IntSum(t *testing.T) {
	expected := 9
	actual := Tuple3Int{2, 3, 4}.Sum()
	AssertEquals(t, expected, actual)
}

func Test_Contains(t *testing.T) {
	cases := []struct {
		slice    []string
		element  string
		expected bool
	}{{
		slice:    []string{"aa", "bb", "cc"},
		element:  "aa",
		expected: true,
	}, {
		slice:    []string{"aa", "bb", "cc"},
		element:  "dd",
		expected: false,
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if actual := Contains(tt.slice, tt.element); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_ContainsAll(t *testing.T) {
	cases := []struct {
		slice    []string
		elements []string
		expected bool
	}{{
		slice:    []string{"aa", "bb", "cc"},
		elements: []string{"aa", "cc"},
		expected: true,
	}, {
		slice:    []string{"aa", "bb", "cc"},
		elements: []string{"aa", "dd"},
		expected: false,
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if actual := ContainsAll(tt.slice, tt.elements); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_Filter(t *testing.T) {
	cases := []struct {
		slice    []string
		expected []string
	}{{
		slice:    nil,
		expected: nil,
	}, {
		slice:    []string{},
		expected: []string{},
	}, {
		slice:    []string{"aa", "bb", "cc", "ab", "ac"},
		expected: []string{"aa", "ab", "ac"},
	}}

	filterTestFunction := func(element string) bool {
		return strings.HasPrefix(element, "a")
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := Filter(tt.slice, filterTestFunction)
			if !ContainsAll(actual, tt.expected) {
				t.Errorf("Expected:%v, Actual:%v", tt.expected, actual)
			}

			if len(tt.expected) != len(actual) {
				t.Errorf("Lengths do not match: expected:%v, actual:%v", len(tt.expected), len(actual))
			}
		})
	}
}

func Test_FilterEmpty(t *testing.T) {
	cases := []struct {
		slice    []string
		expected []string
	}{{
		slice:    nil,
		expected: nil,
	}, {
		slice:    []string{},
		expected: []string{},
	}, {
		slice:    []string{"aa", "bb", "cc", "", "ab", "ac", ""},
		expected: []string{"aa", "bb", "cc", "ab", "ac"},
	}}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := FilterEmpty(tt.slice)
			if !ContainsAll(actual, tt.expected) {
				t.Errorf("Expected:%v, Actual:%v", tt.expected, actual)
			}

			if len(tt.expected) != len(actual) {
				t.Errorf("Lengths do not match: expected:%v, actual:%v", len(tt.expected), len(actual))
			}
		})
	}
}

func Test_Partition(t *testing.T) {
	cases := []struct {
		given    []string
		expected [][]string
	}{{
		given:    []string{"", "a", "b", "", "c", "", "d"},
		expected: [][]string{{""}, {"a", "b", ""}, {"c", ""}, {"d"}},
	}}

	partitionBy := func(line string) bool { return line == "" }

	equals := func(expected [][]string, actual [][]string) bool {
		if len(expected) != len(actual) {
			return false
		}

		for i := 0; i < len(expected); i++ {
			for j := 0; j < len(expected[i]); j++ {
				currentExpected := expected[i][j]
				currentActual := actual[i][j]
				if currentExpected != currentActual {
					return false
				}
			}
		}
		return true
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			actual := Partition(tt.given, partitionBy)
			log.Printf("Expected:%v, Actual:%v", tt.expected, actual)
			if !equals(tt.expected, actual) {
				t.Errorf("Expected:%v, Actual:%v", tt.expected, actual)
			}
		})
	}
}
