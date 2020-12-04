package days

import (
	"strings"

	"github.com/lavinj87/advent/internal/lib"
)

func day3() {}

type world struct {
	rows []string
}

func newWorld(description string) *world {
	rows := strings.Split(description, "\n")
	return &world{rows: rows}
}

func (m world) cursor(slope slope) *cursor {
	return &cursor{world: m, slope: slope}
}

func (m world) pointAt(position position) byte {
	if position.vertical >= uint(len(m.rows)) {
		return 0
	}

	row := m.rows[position.vertical]

	return byte(row[position.horizontal%uint(len(row))])
}

func (m world) countTrees(slope slope) int {
	counter := 0
	cursor := m.cursor(slope)
	pointAt := m.pointAt(cursor.current)

	for pointAt != 0 {
		if isTree(pointAt) {
			counter++
		}
		cursor.moveNext()
		pointAt = m.pointAt(cursor.current)
	}
	return counter
}

func descriptionFromFile(inputPath string) string {
	lines := []string{}

	lib.Lines(inputPath, func(line string) {
		lines = append(lines, line)
	})

	return strings.Join(lines, "\n")
}

type cursor struct {
	world   world
	slope   slope
	current position
}

func (m *cursor) moveNext() {
	m.current.horizontal += m.slope.horizontal
	m.current.vertical += m.slope.vertical
}

type slope struct {
	horizontal uint
	vertical   uint
}

type position struct {
	vertical   uint
	horizontal uint
}

func isTree(point byte) bool {
	return point == '#'
}

func isOpen(point byte) bool {
	return point == '.'
}
