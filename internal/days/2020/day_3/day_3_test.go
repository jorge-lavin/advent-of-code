package days

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_isTree(t *testing.T) {
	tests := []struct {
		point    byte
		expected bool
	}{{
		point:    '#',
		expected: true,
	}, {
		point:    '.',
		expected: false,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := isTree(tt.point); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_isOpen(t *testing.T) {
	tests := []struct {
		point    byte
		expected bool
	}{{
		point:    '#',
		expected: false,
	}, {
		point:    '.',
		expected: true,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := isOpen(tt.point); actual != tt.expected {
				t.Errorf("Expected:%v. Actual:%v", tt.expected, actual)
			}
		})
	}
}

func Test_newWorld(t *testing.T) {
	t.Run("", func(t *testing.T) {
		worldDescription := "..##.......\n" + "#...#...#.."

		actual := newWorld(worldDescription)
		expected := &world{rows: []string{"..##.......", "#...#...#.."}}

		if (!cmp.Equal(actual, expected, cmp.AllowUnexported(world{}))) {
			t.Errorf("Expected:%v. Actual:%v", expected, actual)
		}

	})
}

func Test_moveNext(t *testing.T) {
	t.Run("", func(t *testing.T) {
		worldDescription := "..##.......\n" + "#...#...#.."

		world := newWorld(worldDescription)
		slope := slope{horizontal: 1, vertical: 2}
		cursor := world.cursor(slope)
		cursor.moveNext()

		actual := cursor.current
		expected := position{horizontal: 1, vertical: 2}

		if !cmp.Equal(expected, actual, cmp.AllowUnexported(position{})) {
			t.Errorf("Expected:%v, actual:%v", expected, actual)
		}
	})
}

func Test_pointAt(t *testing.T) {
	worldDescription := "..##.......\n" + "#...#...#.."

	world := newWorld(worldDescription)

	tests := []struct {
		position position
		expected byte
	}{{
		position: position{0, 0},
		expected: '.',
	}, {
		position: position{0, 2},
		expected: '#',
	}, {
		position: position{2, 0},
		expected: 0,
	}, {
		position: position{0, 14},
		expected: '#',
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if actual := world.pointAt(tt.position); string(actual) != string(tt.expected) {
				t.Errorf("Expected:%c, actual:%c", tt.expected, actual)
			}
		})
	}
}

func Test_countTrees(t *testing.T) {
	worldDescription := strings.Join([]string{
		"#.##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}, "\n")

	world := newWorld(worldDescription)

	slope := slope{3, 1}
	expected := 8
	actual := world.countTrees(slope)

	if actual != expected {
		t.Errorf("Expected %d trees, got %d", expected, actual)
	}
}

func Test_day3(t *testing.T) {
	worldDescription := descriptionFromFile("day_3_input.txt")
	world := newWorld(worldDescription)
	slope := slope{3, 1}
	expected := 244
	actual := world.countTrees(slope)

	if actual != expected {
		t.Errorf("Expected %d trees, got %d", expected, actual)
	}
}

func Test_day3Multiply(t *testing.T) {
	worldDescription := descriptionFromFile("day_3_input.txt")
	world := newWorld(worldDescription)
	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	multiplier := 1
	for _, slope := range slopes {
		multiplier = multiplier * world.countTrees(slope)
	}

	expected := 9406609920
	actual := multiplier

	if actual != expected {
		t.Errorf("Expected %d trees, got %d", expected, actual)
	}
}
