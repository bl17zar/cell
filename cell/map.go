package cell

import (
	"fmt"
	"strings"
)

type SignType string

const (
	signEmpty          SignType = "·"
	signNode           SignType = "◼"
	signEdgeHorizontal SignType = "—"
	signEdgeVertical   SignType = "|"
)

type Map struct {
	xExtender string
	size      int
	values    [][]SignType
}

func NewMap(size, xMult int) *Map {
	values := make([][]SignType, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, make([]SignType, size, size))
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			values[i][j] = signEmpty
		}
	}

	xMultList := make([]string, 0, xMult)
	for i := 0; i < xMult; i++ {
		xMultList = append(xMultList, " ")
	}

	return &Map{
		xExtender: strings.Join(xMultList, ""),
		size:      size,
		values:    values,
	}
}

func (m *Map) Copy() *Map {
	res := NewMap(m.size, len(strings.Split(m.xExtender, "")))

	for i, row := range m.values {
		res.values[i] = row
	}

	return res
}

func (m *Map) Set(row int, column int, v SignType) {
	m.values[row-1][column-1] = v
}

func (m *Map) String() string {
	b := strings.Builder{}

	for _, row := range m.values {
		for _, el := range row {
			b.WriteString(fmt.Sprint(m.xExtender, el))
		}

		b.WriteString("\n")
	}

	return b.String()
}

func (m *Map) IsEmpty(row int, col int) bool {
	return row-1 >= 0 && row-1 < m.size && col-1 >= 0 && col-1 < m.size && m.values[row-1][col-1] == signEmpty
}
