package lib

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

// Tuple2Int is a pair of ints
type Tuple2Int struct {
	One int
	Two int
}

// Sum returns the sum of the two elements of the tuple
func (m Tuple2Int) Sum() int {
	return m.One + m.Two
}

// Multiply returns the product of the two elements of the tuple
func (m Tuple2Int) Multiply() int {
	return m.One * m.Two
}

// Tuple3Int is a trio of ints
type Tuple3Int struct {
	One   int
	Two   int
	Three int
}

// Sum returns the sum of the three elements of the tuple
func (m Tuple3Int) Sum() int {
	return m.One + m.Two + m.Three
}

// Multiply returns the product of the two elements of the tuple
func (m Tuple3Int) Multiply() int {
	return m.One * m.Two * m.Three
}

// Lines reads the file and applies the given action to each line, one at a time.
func Lines(path string, action func(string)) error {
	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		action(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// StringToInt convert a String to a decimal int64
func StringToInt(number string) int {
	num, _ := strconv.ParseInt(number, 10, 64)
	return int(num)
}

// AssertEquals compares expected with actual and invoques ErrorF if they are not equal (according to ==)
func AssertEquals(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected:%v. Actual:%v", expected, actual)
	}
}

func AssertEqualsWithTest(t *testing.T, expected, actual interface{}, test func(expected, actual interface{}) bool) {
	if !test(expected, actual) {
		t.Errorf("Expected:%v. Actual:%v", expected, actual)
	}
}
