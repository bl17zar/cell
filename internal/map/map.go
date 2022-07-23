package _map

import (
	"fmt"
	"strings"
)

const (
	LEFT  direction = "LEFT"
	RIGHT direction = "RIGHT"
	UP    direction = "UP"
	DOWN  direction = "DOWN"
)

type direction string

type Map struct {
	Values [][]Symbol
}

func (m *Map) Hash() string {
	b := strings.Builder{}
	for _, row := range m.Values {
		for _, sym := range row {
			b.WriteString(sym.String())
		}
	}

	return b.String()
}

func New(size int) *Map {
	values := make([][]Symbol, size)
	for i := 0; i < size; i++ {
		values[i] = make([]Symbol, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			values[i][j] = SymbolEmpty
		}
	}

	return &Map{
		Values: values,
	}
}

func (m *Map) Copy() *Map {
	res := New(len(m.Values))

	for i, row := range m.Values {
		for j := range row {
			res.Values[i][j] = row[j]
		}

	}

	return res
}

func (m *Map) AddNode(row int, column int) {
	m.set(row, column, SymbolNode)
}

func (m *Map) set(row int, column int, v Symbol) {
	m.Values[row][column] = v
}

func (m *Map) String() string {
	b := strings.Builder{}

	for _, row := range m.Values {
		for _, el := range row {
			b.WriteString(fmt.Sprint(el))
		}

		b.WriteString("\n")
	}

	return b.String()
}

func (m *Map) IsInsideBorders(row, col int) bool {
	return row-1 >= 0 && row < len(m.Values) && col >= 0 && col < len(m.Values)
}

func (m *Map) IsObstacle(row int, col int) bool {
	return m.IsInsideBorders(row, col) && m.Values[row][col] == SymbolObstacle || m.IsInsideBorders(row, col) && m.Values[row][col] == SymbolNode
}

func (m *Map) CheckPathForNewNode(row, col int, dir direction) bool {
	switch dir {
	case LEFT:
		return !m.IsObstacle(row, col-1) && m.IsInsideBorders(row, col-1) && !m.IsObstacle(row, col-2) && m.IsInsideBorders(row, col-2)
	case RIGHT:
		return !m.IsObstacle(row, col+1) && m.IsInsideBorders(row, col+1) && !m.IsObstacle(row, col+2) && m.IsInsideBorders(row, col+2)
	case UP:
		return !m.IsObstacle(row-1, col) && m.IsInsideBorders(row-1, col) && !m.IsObstacle(row-2, col) && m.IsInsideBorders(row-2, col)
	case DOWN:
		return !m.IsObstacle(row+1, col) && m.IsInsideBorders(row+1, col) && !m.IsObstacle(row+2, col) && m.IsInsideBorders(row+2, col)
	}

	return false
}

func (m *Map) AddEdge(n1Row, n1Col, n2Row, n2Col int) {
	if n1Row == n2Row {
		var col int
		if (n2Col - n1Col) > 0 {
			col = n1Col + 1
		} else {
			col = n2Col + 1
		}

		m.set(n1Row, col, SymbolEdgeHorizontal)
	} else {
		var row int
		if (n2Row - n1Row) > 0 {
			row = n1Row + 1
		} else {
			row = n2Row + 1
		}

		m.set(row, n1Col, SymbolEdgeVertical)
	}
}

func (m *Map) AddCycleNode(row int, col int) {
	m.set(row, col, SymbolCycleNode)
}

func (m *Map) SetEmpty(col int, row int) {
	m.set(row, col, SymbolEmpty)
}

func (m *Map) SetEmptyEdge(n1Row, n1Col, n2Row, n2Col int) {
	if n1Row == n2Row {
		var col int
		if (n2Col - n1Col) > 0 {
			col = n1Col + 1
		} else {
			col = n2Col + 1
		}

		m.set(n1Row, col, SymbolEmpty)
	} else {
		var row int
		if (n2Row - n1Row) > 0 {
			row = n1Row + 1
		} else {
			row = n2Row + 1
		}

		m.set(row, n1Col, SymbolEmpty)
	}
}

func (m *Map) AddCycleEdge(n1Row, n1Col, n2Row, n2Col int) {
	if n1Row == n2Row {
		var col int
		if (n2Col - n1Col) > 0 {
			col = n1Col + 1
		} else {
			col = n2Col + 1
		}

		m.set(n1Row, col, SymbolCycleEdgeHorizontal)
	} else {
		var row int
		if (n2Row - n1Row) > 0 {
			row = n1Row + 1
		} else {
			row = n2Row + 1
		}

		m.set(row, n1Col, SymbolCycleEdgeVertical)
	}
}

func (m *Map) AddObstacle(row, col int) {
	m.set(row, col, SymbolObstacle)
}
