package _map

import (
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

type coords struct {
	row int
	col int
}

func (c coords) Row() int {
	return c.row
}

func (c coords) Col() int {
	return c.col
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
		res.Values[i] = row
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
			b.WriteString(string(el))
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

// func (m *Map) WithoutCycles(nodes []*Node) *Map {
// 	diff := []*Display{}
//
// 	for _, n := range nodes {
// 		if !m.IsCycleNode(n.Row(), n.Col()) {
// 			panic(errors.New("not cycle node"))
// 		}
//
// 		diff = append(diff, &Display{
// 			Row:  n.Row,
// 			Col:  n.Col,
// 			Sign: SymbolEmpty,
// 		})
//
// 		if m.IsInsideBorders(n.Row, n.Col+1) {
// 			diff = append(diff, &Display{
// 				Row:  n.Row,
// 				Col:  n.Col + 1,
// 				Sign: SymbolEmpty,
// 			})
// 		}
//
// 		if m.IsInsideBorders(n.Row+1, n.Col) {
// 			diff = append(diff, &Display{
// 				Row:  n.Row + 1,
// 				Col:  n.Col,
// 				Sign: SymbolEmpty,
// 			})
// 		}
//
// 		if m.IsInsideBorders(n.Row, n.Col-1) {
// 			diff = append(diff, &Display{
// 				Row:  n.Row,
// 				Col:  n.Col - 1,
// 				Sign: SymbolEmpty,
// 			})
// 		}
//
// 		if m.IsInsideBorders(n.Row-1, n.Col) {
// 			diff = append(diff, &Display{
// 				Row:  n.Row - 1,
// 				Col:  n.Col,
// 				Sign: SymbolEmpty,
// 			})
// 		}
// 	}
//
// 	for _, d := range diff {
// 		m.Set(d.Row, d.Col, d.Sign)
// 	}
//
// 	return m.Copy()
// }
//
//

//
// func (m *Map) IsEmpty(row int, col int) bool {
// 	return m.IsInsideBorders(row, col) && m.Values[row-1][col-1] == SymbolEmpty
// }
//

//
// func (m *Map) IsCycleNode(row int, col int) bool {
// 	return m.IsInsideBorders(row, col) && m.Values[row-1][col-1] == SymbolCycleNode
// }
//
// func (m *Map) IsCycleEdge(row int, col int) bool {
// 	return m.IsInsideBorders(row, col) && (m.Values[row-1][col-1] == SymbolCycleEdgeHorizontal || m.Values[row-1][col-1] == SymbolCycleEdgeVertical)
// }
//

//
// 	return display
// }
//
// func GetCycleEdgeDisplay(n1, n2 *Node) *Display {
// 	var display *Display
//
// 	if n1.Row == n2Row
// 		var col int
// 		if (n2.Col - n1.Col) > 0 {
// 			col = n1.Col + 1
// 		} else {
// 			col = n2.Col + 1
// 		}
//
// 		display = &Display{
// 			Row:  n1.Row,
// 			Col:  col,
// 			Sign: SymbolCycleEdgeHorizontal,
// 		}
// 	} else {
// 		var row int
// 		if (n2.Row - n1.Row) > 0 {
// 			row = n1.Row + 1
// 		} else {
// 			row = n2.Row + 1
// 		}
//
// 		display = &Display{
// 			Row:  row,
// 			Col:  n1.Col,
// 			Sign: SymbolCycleEdgeVertical,
// 		}
// 	}
//
// 	return display
// }
//
// func GetNodeDisplay(n *Node) *Display {
// 	return &Display{
// 		Row:  n.Row,
// 		Col:  n.Col,
// 		Sign: SymbolNode,
// 	}
// }
//
// func GetCycleNodeDisplay(n *Node) *Display {
// 	return &Display{
// 		Row:  n.Row,
// 		Col:  n.Col,
// 		Sign: SymbolCycleNode,
// 	}
// }
