package days

import (
	"github.com/lavinj87/advent/internal/lib"
)

func day1Tuple3(inputPath string) int {
	var numbers []int
	lib.Lines(inputPath, func(line string) {
		num := lib.StringToInt(line)
		numbers = append(numbers, num)
	})

	tuple := findTuple3ThatSums2020(numbers)
	return tuple.Multiply()
}

func day1Tuple2(inputPath string) int {
	var numbers []int
	lib.Lines(inputPath, func(line string) {
		num := lib.StringToInt(line)
		numbers = append(numbers, num)
	})

	tuple := findTuple2ThatSums2020(numbers)
	return tuple.Multiply()
}

func findTuple2ThatSums2020(numbers []int) lib.Tuple2Int {

	for _, outer := range numbers {
		for _, inner := range numbers {
			if outer+inner == 2020 {
				return lib.Tuple2Int{One: outer, Two: inner}
			}
		}
	}
	return lib.Tuple2Int{One: 1, Two: 1}
}

func findTuple3ThatSums2020(numbers []int) lib.Tuple3Int {
	for _, first := range numbers {
		for _, second := range numbers {
			for _, third := range numbers {
				if first+second+third == 2020 {
					return lib.Tuple3Int{One: first, Two: second, Three: third}
				}
			}
		}
	}
	return lib.Tuple3Int{One: 1, Two: 1, Three: 1}
}
