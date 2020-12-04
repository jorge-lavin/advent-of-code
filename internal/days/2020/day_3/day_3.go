package days

import (
	"strings"
)

func day3() {}

type world struct {
	rows []string
}

func (m world) cursor(slope slope) *cursor {
	return &cursor{world: m, slope: slope}
}

func (m world) pointAt(position position) byte {
	if position.vertical >= uint(len(m.rows)) {
		return 0
	}

	row := m.rows[position.vertical]

	if position.horizontal >= uint(len(row)) {
		return 0
	}

	return byte(row[position.horizontal])
}

func newWorld(description string) *world {
	rows := strings.Split(description, "\n")
	return &world{rows: rows}
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
	horizontal uint
	vertical   uint
}

func isTree(point byte) bool {
	return point == '#'
}

func isOpen(point byte) bool {
	return point == '.'
}
