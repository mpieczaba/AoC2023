package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type position struct {
	line   int
	column int
}

func newPosition(line, column int) position {
	return position{line, column}
}

type maze struct {
	tiles           [][]byte
	tilesMask       [][]bool
	startPosition   position
	currentPosition position
}

func (m *maze) checkPosition(p position) bool {
	return 0 <= p.column && p.column < len(m.tiles[0]) && 0 <= p.line && p.line < len(m.tiles) && !m.mapAtPosition(p)
}

func (m *maze) tileAtPosition(p position) byte {
	return m.tiles[p.line][p.column]
}

func (m *maze) mapAtPosition(p position) bool {
	return m.tilesMask[p.line][p.column]
}

func (m *maze) checkNorth() (position, bool) {
	north := newPosition(m.currentPosition.line-1, m.currentPosition.column)

	return north, m.checkPosition(north) && bytes.Contains([]byte{'|', '7', 'F'}, []byte{m.tileAtPosition(north)})
}

func (m *maze) checkEast() (position, bool) {
	east := newPosition(m.currentPosition.line, m.currentPosition.column+1)

	return east, m.checkPosition(east) && bytes.Contains([]byte{'-', '7', 'J'}, []byte{m.tileAtPosition(east)})
}

func (m *maze) checkSouth() (position, bool) {
	south := newPosition(m.currentPosition.line+1, m.currentPosition.column)

	return south, m.checkPosition(south) && bytes.Contains([]byte{'|', 'L', 'J'}, []byte{m.tileAtPosition(south)})
}

func (m *maze) checkWest() (position, bool) {
	west := newPosition(m.currentPosition.line, m.currentPosition.column-1)

	return west, m.checkPosition(west) && bytes.Contains([]byte{'-', 'L', 'F'}, []byte{m.tileAtPosition(west)})
}

func (m *maze) nextPosition() position {
	newPosition := m.startPosition

	switch m.tileAtPosition(m.currentPosition) {
	case 'S':
		// Can go NESW
		if p, ok := m.checkNorth(); ok {
			newPosition = p

		} else if p, ok = m.checkEast(); ok {
			newPosition = p

		} else if p, ok := m.checkSouth(); ok {
			newPosition = p

		} else if p, ok := m.checkWest(); ok {
			newPosition = p

		}
	case '|':
		// Can go NS
		if p, ok := m.checkNorth(); ok {
			newPosition = p
		} else if p, ok = m.checkSouth(); ok {
			newPosition = p
		}
	case '-':
		// Can fo EW
		if p, ok := m.checkEast(); ok {
			newPosition = p
		} else if p, ok = m.checkWest(); ok {
			newPosition = p
		}
	case 'L':
		// Can go NE
		if p, ok := m.checkNorth(); ok {
			newPosition = p
		} else if p, ok = m.checkEast(); ok {
			newPosition = p
		}

	case 'J':
		// Can go NW
		if p, ok := m.checkNorth(); ok {
			newPosition = p
		} else if p, ok = m.checkWest(); ok {
			newPosition = p
		}

	case '7':
		// Can go SW
		if p, ok := m.checkSouth(); ok {
			newPosition = p
		} else if p, ok = m.checkWest(); ok {
			newPosition = p
		}
	case 'F':
		// Can so ES
		if p, ok := m.checkEast(); ok {
			newPosition = p
		} else if p, ok = m.checkSouth(); ok {
			newPosition = p
		}
	}

	m.tilesMask[m.currentPosition.line][m.currentPosition.column] = true

	return newPosition
}

func main() {
	file, _ := os.Open("day10/input.txt")
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Buffer(make([]byte, 141*141), 141*141)

	var m maze

	for s.Scan() {
		m.tiles = append(m.tiles, s.Bytes())

		if i := bytes.IndexRune(s.Bytes(), 'S'); i != -1 {
			m.startPosition = newPosition(len(m.tiles)-1, i)
			m.currentPosition = m.startPosition
		}
	}

	m.tilesMask = make([][]bool, len(m.tiles))

	for i := 0; i < len(m.tiles); i++ {
		m.tilesMask[i] = make([]bool, len(m.tiles[0]))
	}

	m.currentPosition = m.nextPosition()

	resultPartOne := 1

	for m.currentPosition != m.startPosition {
		m.currentPosition = m.nextPosition()

		resultPartOne++
	}

	resultPartOne /= 2

	var resultPartTwo int

	for i := 0; i < len(m.tilesMask); i++ {
		isIn := false

		for j := 0; j < len(m.tilesMask[i]); j++ {
			if m.mapAtPosition(newPosition(i, j)) {
				tile := m.tileAtPosition(newPosition(i, j))

				if tile == 'F' || tile == '7' || tile == '|' {

					isIn = !isIn
				}
			} else if isIn {
				resultPartTwo++
			}
		}
	}

	fmt.Printf("Part one: %d\n", resultPartOne)
	fmt.Printf("Part two: %d\n", resultPartTwo)
}
