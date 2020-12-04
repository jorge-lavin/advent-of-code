package lib

import (
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
