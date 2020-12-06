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

// Filter a given slice of strings and returns a new slice with all the elements that passes the test.
func Filter(slice []string, test func(element string) bool) []string {
	if slice == nil {
		return nil
	}

	if len(slice) == 0 {
		return slice
	}

	filtered := []string{}
	for _, element := range slice {
		if test(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered

}

// Contains returns true if the element is contained in the slice and false otherwise
func Contains(slice []string, element string) bool {
	for _, sliceElement := range slice {
		if sliceElement == element {
			return true
		}
	}
	return false
}

// ContainsAll return true if all the elements are contained in the slice and false otherwise
func ContainsAll(slice []string, elements []string) bool {
	for _, element := range elements {
		if !Contains(slice, element) {
			return false
		}
	}
	return true
}

// FilterEmpty a given slice of strings and returns a new slice with all the elements that are not empty strings
func FilterEmpty(slice []string) []string {
	return Filter(slice, func(element string) bool {
		return len(element) > 0
	})
}

// Partition returns an array of array of strings partitioned by the test function.
func Partition(lines []string, test func(line string) bool) [][]string {
	startIndex := 0
	reduced := [][]string{}
	for currentIndex, line := range lines {
		if test(line) {
			reduced = append(reduced, lines[startIndex:currentIndex+1])
			startIndex = currentIndex + 1
		}
	}

	// append the last part of the array of arrays if needed.
	if startIndex != len(lines) {
		reduced = append(reduced, lines[startIndex:len(lines)])
	}

	return reduced
}
